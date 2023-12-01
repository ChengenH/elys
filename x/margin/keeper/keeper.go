package keeper

import (
	"fmt"
	"math/big"

	"math"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	ammtypes "github.com/elys-network/elys/x/amm/types"
	assetprofiletypes "github.com/elys-network/elys/x/assetprofile/types"
	"github.com/elys-network/elys/x/margin/types"
	pkeeper "github.com/elys-network/elys/x/parameter/keeper"
	ptypes "github.com/elys-network/elys/x/parameter/types"
)

type (
	Keeper struct {
		types.AuthorizationChecker
		types.PositionChecker
		types.PoolChecker
		types.OpenChecker
		types.OpenLongChecker
		types.OpenShortChecker
		types.CloseLongChecker
		types.CloseShortChecker

		cdc             codec.BinaryCodec
		storeKey        storetypes.StoreKey
		memKey          storetypes.StoreKey
		authority       string
		parameterKeeper *pkeeper.Keeper
		amm             types.AmmKeeper
		bankKeeper      types.BankKeeper
		oracleKeeper    ammtypes.OracleKeeper
		apKeeper        types.AssetProfileKeeper

		hooks types.MarginHooks
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	authority string,
	amm types.AmmKeeper,
	bk types.BankKeeper,
	oracleKeeper ammtypes.OracleKeeper,
	apKeeper types.AssetProfileKeeper,
	parameterKeeper *pkeeper.Keeper,
) *Keeper {
	// ensure that authority is a valid AccAddress
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic("authority is not a valid acc address")
	}

	keeper := &Keeper{
		cdc:             cdc,
		storeKey:        storeKey,
		memKey:          memKey,
		authority:       authority,
		amm:             amm,
		bankKeeper:      bk,
		oracleKeeper:    oracleKeeper,
		apKeeper:        apKeeper,
		parameterKeeper: parameterKeeper,
	}

	keeper.AuthorizationChecker = keeper
	keeper.PositionChecker = keeper
	keeper.PoolChecker = keeper
	keeper.OpenChecker = keeper
	keeper.OpenLongChecker = keeper
	keeper.OpenShortChecker = keeper
	keeper.CloseLongChecker = keeper
	keeper.CloseShortChecker = keeper

	return keeper
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) GetMTPCount(ctx sdk.Context) uint64 {
	var count uint64
	countBz := ctx.KVStore(k.storeKey).Get(types.MTPCountPrefix)
	if countBz == nil {
		count = 0
	} else {
		count = types.GetUint64FromBytes(countBz)
	}
	return count
}

func (k Keeper) GetOpenMTPCount(ctx sdk.Context) uint64 {
	var count uint64
	countBz := ctx.KVStore(k.storeKey).Get(types.OpenMTPCountPrefix)
	if countBz == nil {
		count = 0
	} else {
		count = types.GetUint64FromBytes(countBz)
	}
	return count
}

func (k Keeper) CheckIfWhitelisted(ctx sdk.Context, address string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetWhitelistKey(address))
}

// Swap estimation using amm CalcInAmtGivenOut function
func (k Keeper) EstimateSwapGivenOut(ctx sdk.Context, tokenOutAmount sdk.Coin, tokenInDenom string, ammPool ammtypes.Pool) (sdk.Int, error) {
	marginEnabled := k.IsPoolEnabled(ctx, ammPool.PoolId)
	if !marginEnabled {
		return sdk.ZeroInt(), sdkerrors.Wrap(types.ErrMarginDisabled, "Margin disabled pool")
	}

	tokensOut := sdk.Coins{tokenOutAmount}
	// Estimate swap
	snapshot := k.amm.GetPoolSnapshotOrSet(ctx, ammPool)
	swapResult, err := k.amm.CalcInAmtGivenOut(ctx, ammPool.PoolId, k.oracleKeeper, &snapshot, tokensOut, tokenInDenom, sdk.ZeroDec())

	if err != nil {
		return sdk.ZeroInt(), err
	}

	if swapResult.IsZero() {
		return sdk.ZeroInt(), types.ErrAmountTooLow
	}
	return swapResult.Amount, nil
}

func (k Keeper) Borrow(ctx sdk.Context, collateralAsset string, custodyAsset string, collateralAmount sdk.Int, custodyAmount sdk.Int, mtp *types.MTP, ammPool *ammtypes.Pool, pool *types.Pool, eta sdk.Dec, baseCurrency string) error {
	mtpAddress, err := sdk.AccAddressFromBech32(mtp.Address)
	if err != nil {
		return err
	}
	collateralCoin := sdk.NewCoin(collateralAsset, collateralAmount)

	if !k.bankKeeper.HasBalance(ctx, mtpAddress, collateralCoin) {
		return types.ErrBalanceNotAvailable
	}

	collateralAmountDec := sdk.NewDecFromBigInt(collateralAmount.BigInt())
	liabilitiesDec := collateralAmountDec.Mul(eta)

	// If collateral asset is not base currency, should calculate liability in base currency with the given out.
	// Liability has to be in base currency
	if collateralAsset != baseCurrency {
		// ATOM amount
		etaAmt := liabilitiesDec.TruncateInt()
		etaAmtToken := sdk.NewCoin(collateralAsset, etaAmt)
		// Calculate base currency amount given atom out amount and we use it liabilty amount in base currency
		liabilityAmt, err := k.OpenLongChecker.EstimateSwapGivenOut(ctx, etaAmtToken, baseCurrency, *ammPool)
		if err != nil {
			return err
		}

		liabilitiesDec = sdk.NewDecFromInt(liabilityAmt)
	}

	collateralIndex, custodyIndex := types.GetMTPAssetIndex(mtp, collateralAsset, custodyAsset)
	if collateralIndex < 0 || custodyIndex < 0 {
		return sdkerrors.Wrap(types.ErrBalanceNotAvailable, "MTP collateral or custody invalid!")
	}

	mtp.Collaterals[collateralIndex].Amount = mtp.Collaterals[collateralIndex].Amount.Add(collateralAmount)
	mtp.Liabilities = mtp.Liabilities.Add(sdk.NewIntFromBigInt(liabilitiesDec.TruncateInt().BigInt()))
	mtp.Custodies[custodyIndex].Amount = mtp.Custodies[custodyIndex].Amount.Add(custodyAmount)

	// calculate mtp take profit custody, delta y_tp_c = delta x_l / take profit price (take profit custody = liabilities / take profit price)
	mtp.TakeProfitCustodies = types.CalcMTPTakeProfitCustodies(mtp)

	// calculate mtp take profit liablities, delta x_tp_l = delta y_tp_c * current price (take profit liabilities = take profit custody * current price)
	mtp.TakeProfitLiabilities, err = k.CalcMTPTakeProfitLiability(ctx, mtp, baseCurrency)
	if err != nil {
		return err
	}

	mtp.Leverages = append(mtp.Leverages, eta.Add(sdk.OneDec()))

	// print mtp.CustodyAmount
	ctx.Logger().Info(fmt.Sprintf("mtp.CustodyAmount: %s", mtp.Custodies[custodyIndex].Amount.String()))

	h, err := k.UpdateMTPHealth(ctx, *mtp, *ammPool, baseCurrency) // set mtp in func or return h?
	if err != nil {
		return err
	}
	mtp.MtpHealth = h

	ammPoolAddr, err := sdk.AccAddressFromBech32(ammPool.Address)
	if err != nil {
		return err
	}

	collateralCoins := sdk.NewCoins(collateralCoin)
	err = k.bankKeeper.SendCoins(ctx, mtpAddress, ammPoolAddr, collateralCoins)

	if err != nil {
		return err
	}

	err = pool.UpdateBalance(ctx, collateralAsset, collateralAmount, true, mtp.Position)
	if err != nil {
		return err
	}

	// All liability has to be in base currency
	err = pool.UpdateLiabilities(ctx, baseCurrency, mtp.Liabilities, true, mtp.Position)
	if err != nil {
		return err
	}

	// All take profit liability has to be in base currency
	err = pool.UpdateTakeProfitLiabilities(ctx, baseCurrency, mtp.TakeProfitLiabilities, true, mtp.Position)
	if err != nil {
		return err
	}

	// All take profit custody has to be in base currency
	err = pool.UpdateTakeProfitCustody(ctx, baseCurrency, mtp.TakeProfitCustodies[custodyIndex].Amount, true, mtp.Position)
	if err != nil {
		return err
	}

	k.SetPool(ctx, *pool)

	return k.SetMTP(ctx, mtp)
}

func (k Keeper) UpdatePoolHealth(ctx sdk.Context, pool *types.Pool) error {
	pool.Health = k.CalculatePoolHealth(ctx, pool)
	k.SetPool(ctx, *pool)

	return nil
}

func (k Keeper) CalculatePoolHealthByPosition(ctx sdk.Context, pool *types.Pool, ammPool ammtypes.Pool, position types.Position) sdk.Dec {
	poolAssets := pool.GetPoolAssets(position)
	H := sdk.NewDec(1)
	for _, asset := range *poolAssets {
		ammBalance, err := types.GetAmmPoolBalance(ammPool, asset.AssetDenom)
		if err != nil {
			return sdk.ZeroDec()
		}

		balance := sdk.NewDecFromInt(asset.AssetBalance.Add(ammBalance))

		// X_L = X_P_L - X_TP_L (pool liabilities = pool synthetic liabilities - pool take profit liabilities)
		liabilities := sdk.NewDecFromInt(asset.Liabilities.Sub(asset.TakeProfitLiabilities))

		if balance.Add(liabilities).IsZero() {
			return sdk.ZeroDec()
		}

		mul := balance.Quo(balance.Add(liabilities))
		H = H.Mul(mul)
	}
	return H
}

func (k Keeper) CalculatePoolHealth(ctx sdk.Context, pool *types.Pool) sdk.Dec {
	ammPool, found := k.amm.GetPool(ctx, pool.AmmPoolId)
	if !found {
		return sdk.ZeroDec()
	}

	H := k.CalculatePoolHealthByPosition(ctx, pool, ammPool, types.Position_LONG)
	H = H.Mul(k.CalculatePoolHealthByPosition(ctx, pool, ammPool, types.Position_SHORT))

	return H
}

func (k Keeper) TakeInCustody(ctx sdk.Context, mtp types.MTP, pool *types.Pool) error {
	for i := range mtp.Custodies {
		err := pool.UpdateBalance(ctx, mtp.Custodies[i].Denom, mtp.Custodies[i].Amount, false, mtp.Position)
		if err != nil {
			return nil
		}
		err = pool.UpdateCustody(ctx, mtp.Custodies[i].Denom, mtp.Custodies[i].Amount, true, mtp.Position)
		if err != nil {
			return nil
		}
	}

	k.SetPool(ctx, *pool)

	return nil
}

func (k Keeper) IncrementalBorrowInterestPayment(ctx sdk.Context, collateralAsset string, custodyAsset string, borrowInterestPayment sdk.Int, mtp *types.MTP, pool *types.Pool, ammPool ammtypes.Pool, baseCurrency string) (sdk.Int, error) {
	collateralIndex, custodyIndex := types.GetMTPAssetIndex(mtp, collateralAsset, custodyAsset)

	// if mtp has unpaid borrow interest, add to payment
	// convert it into base currency
	if mtp.BorrowInterestUnpaidCollaterals[collateralIndex].Amount.GT(sdk.ZeroInt()) {
		if mtp.Collaterals[collateralIndex].Denom == baseCurrency {
			borrowInterestPayment = borrowInterestPayment.Add(mtp.BorrowInterestUnpaidCollaterals[collateralIndex].Amount)
		} else {
			unpaidCollateralIn := sdk.NewCoin(mtp.Collaterals[collateralIndex].Denom, mtp.BorrowInterestUnpaidCollaterals[collateralIndex].Amount)
			C, err := k.EstimateSwapGivenOut(ctx, unpaidCollateralIn, baseCurrency, ammPool)
			if err != nil {
				return sdk.ZeroInt(), err
			}

			borrowInterestPayment = borrowInterestPayment.Add(C)
		}
	}

	borrowInterestPaymentTokenIn := sdk.NewCoin(baseCurrency, borrowInterestPayment)
	// swap borrow interest payment to custody asset for payment
	borrowInterestPaymentCustody, err := k.EstimateSwap(ctx, borrowInterestPaymentTokenIn, mtp.Custodies[custodyIndex].Denom, ammPool)
	if err != nil {
		return sdk.ZeroInt(), err
	}

	// If collateralAsset is not in base currency, convert it to original asset format
	if collateralAsset != baseCurrency {
		// swap custody amount to collateral for updating borrow interest unpaid
		amtTokenIn := sdk.NewCoin(baseCurrency, borrowInterestPayment)
		borrowInterestPayment, err = k.EstimateSwap(ctx, amtTokenIn, collateralAsset, ammPool) // may need spot price here to not deduct fee
		if err != nil {
			return sdk.ZeroInt(), err
		}
	}

	// if paying unpaid borrow interest reset to 0
	mtp.BorrowInterestUnpaidCollaterals[collateralIndex].Amount = sdk.ZeroInt()

	// edge case, not enough custody to cover payment
	if borrowInterestPaymentCustody.GT(mtp.Custodies[custodyIndex].Amount) {
		// swap custody amount to collateral for updating borrow interest unpaid
		custodyAmtTokenIn := sdk.NewCoin(mtp.Custodies[custodyIndex].Denom, mtp.Custodies[custodyIndex].Amount)
		custodyAmountCollateral, err := k.EstimateSwap(ctx, custodyAmtTokenIn, collateralAsset, ammPool) // may need spot price here to not deduct fee
		if err != nil {
			return sdk.ZeroInt(), err
		}
		mtp.BorrowInterestUnpaidCollaterals[collateralIndex].Amount = borrowInterestPayment.Sub(custodyAmountCollateral)

		borrowInterestPayment = custodyAmountCollateral
		borrowInterestPaymentCustody = mtp.Custodies[custodyIndex].Amount
	}

	// add payment to total paid - collateral
	mtp.BorrowInterestPaidCollaterals[collateralIndex].Amount = mtp.BorrowInterestPaidCollaterals[collateralIndex].Amount.Add(borrowInterestPayment)

	// add payment to total paid - custody
	mtp.BorrowInterestPaidCustodies[custodyIndex].Amount = mtp.BorrowInterestPaidCustodies[custodyIndex].Amount.Add(borrowInterestPaymentCustody)

	// deduct borrow interest payment from custody amount
	mtp.Custodies[custodyIndex].Amount = mtp.Custodies[custodyIndex].Amount.Sub(borrowInterestPaymentCustody)

	takePercentage := k.GetIncrementalBorrowInterestPaymentFundPercentage(ctx)
	fundAddr := k.GetIncrementalBorrowInterestPaymentFundAddress(ctx)
	takeAmount, err := k.TakeFundPayment(ctx, borrowInterestPaymentCustody, mtp.Custodies[custodyIndex].Denom, takePercentage, fundAddr, &ammPool)
	if err != nil {
		return sdk.ZeroInt(), err
	}
	actualBorrowInterestPaymentCustody := borrowInterestPaymentCustody.Sub(takeAmount)

	if !takeAmount.IsZero() {
		k.EmitFundPayment(ctx, mtp, takeAmount, mtp.Custodies[custodyIndex].Denom, types.EventIncrementalPayFund)
	}

	err = pool.UpdateCustody(ctx, mtp.Custodies[custodyIndex].Denom, borrowInterestPaymentCustody, false, mtp.Position)
	if err != nil {
		return sdk.ZeroInt(), err
	}

	err = pool.UpdateBalance(ctx, mtp.Custodies[custodyIndex].Denom, actualBorrowInterestPaymentCustody, true, mtp.Position)
	if err != nil {
		return sdk.ZeroInt(), err
	}

	err = k.SetMTP(ctx, mtp)
	if err != nil {
		return sdk.ZeroInt(), err
	}

	k.SetPool(ctx, *pool)

	return actualBorrowInterestPaymentCustody, nil
}

func (k Keeper) BorrowInterestRateComputationByPosition(ctx sdk.Context, pool types.Pool, ammPool ammtypes.Pool, position types.Position) (sdk.Dec, error) {
	poolAssets := pool.GetPoolAssets(position)
	targetBorrowInterestRate := sdk.OneDec()
	for _, asset := range *poolAssets {
		ammBalance, err := types.GetAmmPoolBalance(ammPool, asset.AssetDenom)
		if err != nil {
			return sdk.ZeroDec(), err
		}

		balance := sdk.NewDecFromInt(asset.AssetBalance.Add(ammBalance))
		liabilities := sdk.NewDecFromInt(asset.Liabilities)

		// Ensure balance is not zero to avoid division by zero
		if balance.IsZero() {
			return sdk.ZeroDec(), nil
		}
		if balance.Add(liabilities).IsZero() {
			return sdk.ZeroDec(), nil
		}

		mul := balance.Add(liabilities).Quo(balance)
		targetBorrowInterestRate = targetBorrowInterestRate.Mul(mul)
	}
	return targetBorrowInterestRate, nil
}

func (k Keeper) BorrowInterestRateComputation(ctx sdk.Context, pool types.Pool, ammPool ammtypes.Pool) (sdk.Dec, error) {
	ammPool, found := k.amm.GetPool(ctx, pool.AmmPoolId)
	if !found {
		return sdk.ZeroDec(), sdkerrors.Wrap(types.ErrBalanceNotAvailable, "Balance not available")
	}

	borrowInterestRateMax := k.GetBorrowInterestRateMax(ctx)
	borrowInterestRateMin := k.GetBorrowInterestRateMin(ctx)
	borrowInterestRateIncrease := k.GetBorrowInterestRateIncrease(ctx)
	borrowInterestRateDecrease := k.GetBorrowInterestRateDecrease(ctx)
	healthGainFactor := k.GetHealthGainFactor(ctx)

	prevBorrowInterestRate := pool.BorrowInterestRate

	targetBorrowInterestRate := healthGainFactor
	targetBorrowInterestRateLong, err := k.BorrowInterestRateComputationByPosition(ctx, pool, ammPool, types.Position_LONG)
	if err != nil {
		return sdk.ZeroDec(), err
	}
	targetBorrowInterestRateShort, err := k.BorrowInterestRateComputationByPosition(ctx, pool, ammPool, types.Position_SHORT)
	if err != nil {
		return sdk.ZeroDec(), err
	}
	targetBorrowInterestRate = targetBorrowInterestRate.Mul(targetBorrowInterestRateLong)
	targetBorrowInterestRate = targetBorrowInterestRate.Mul(targetBorrowInterestRateShort)

	borrowInterestRateChange := targetBorrowInterestRate.Sub(prevBorrowInterestRate)
	borrowInterestRate := prevBorrowInterestRate
	if borrowInterestRateChange.GTE(borrowInterestRateDecrease.Mul(sdk.NewDec(-1))) && borrowInterestRateChange.LTE(borrowInterestRateIncrease) {
		borrowInterestRate = targetBorrowInterestRate
	} else if borrowInterestRateChange.GT(borrowInterestRateIncrease) {
		borrowInterestRate = prevBorrowInterestRate.Add(borrowInterestRateIncrease)
	} else if borrowInterestRateChange.LT(borrowInterestRateDecrease.Mul(sdk.NewDec(-1))) {
		borrowInterestRate = prevBorrowInterestRate.Sub(borrowInterestRateDecrease)
	}

	newBorrowInterestRate := borrowInterestRate

	if borrowInterestRate.GT(borrowInterestRateMin) && borrowInterestRate.LT(borrowInterestRateMax) {
		newBorrowInterestRate = borrowInterestRate
	} else if borrowInterestRate.LTE(borrowInterestRateMin) {
		newBorrowInterestRate = borrowInterestRateMin
	} else if borrowInterestRate.GTE(borrowInterestRateMax) {
		newBorrowInterestRate = borrowInterestRateMax
	}

	return newBorrowInterestRate, nil
}

func (k Keeper) CheckMinLiabilities(ctx sdk.Context, collateralAmount sdk.Coin, eta sdk.Dec, pool types.Pool, ammPool ammtypes.Pool, custodyAsset string) error {
	var borrowInterestRational, liabilitiesRational, rate big.Rat
	minBorrowInterestRate := k.GetBorrowInterestRateMin(ctx)

	// Ensure minBorrowInterestRate is not zero to avoid division by zero
	if minBorrowInterestRate.IsZero() {
		return types.ErrAmountTooLow
	}

	collateralAmountDec := sdk.NewDecFromInt(collateralAmount.Amount)
	liabilitiesDec := collateralAmountDec.Mul(eta)
	liabilities := sdk.NewUint(liabilitiesDec.TruncateInt().Uint64())

	entry, found := k.apKeeper.GetEntry(ctx, ptypes.BaseCurrency)
	if !found {
		return sdkerrors.Wrapf(assetprofiletypes.ErrAssetProfileNotFound, "asset %s not found", ptypes.BaseCurrency)
	}
	baseCurrency := entry.Denom

	// liabilty has to be always in base currency
	if collateralAmount.Denom != baseCurrency {
		outAmt := liabilitiesDec.TruncateInt()
		outAmtToken := sdk.NewCoin(collateralAmount.Denom, outAmt)
		inAmt, err := k.OpenLongChecker.EstimateSwapGivenOut(ctx, outAmtToken, baseCurrency, ammPool)
		if err != nil {
			return types.ErrBorrowTooLow
		}
		liabilities = sdk.NewUint(inAmt.Uint64())
	}
	rate.SetFloat64(minBorrowInterestRate.MustFloat64())
	liabilitiesRational.SetInt(liabilities.BigInt())
	borrowInterestRational.Mul(&rate, &liabilitiesRational)

	borrowInterestNew := borrowInterestRational.Num().Quo(borrowInterestRational.Num(), borrowInterestRational.Denom())
	samplePayment := sdk.NewInt(borrowInterestNew.Int64())

	if samplePayment.IsZero() && !minBorrowInterestRate.IsZero() {
		return types.ErrBorrowTooLow
	}

	// If collateral is not base currency, custody amount is already checked in HasSufficientBalance function.
	// its liability balance checked in the above if statement, so return
	if collateralAmount.Denom != baseCurrency {
		return nil
	}

	samplePaymentTokenIn := sdk.NewCoin(collateralAmount.Denom, samplePayment)
	// swap borrow interest payment to custody asset
	_, err := k.EstimateSwap(ctx, samplePaymentTokenIn, custodyAsset, ammPool)
	if err != nil {
		return types.ErrBorrowTooLow
	}

	return nil
}

func (k Keeper) DestroyMTP(ctx sdk.Context, mtpAddress string, id uint64) error {
	key := types.GetMTPKey(mtpAddress, id)
	store := ctx.KVStore(k.storeKey)
	if !store.Has(key) {
		return types.ErrMTPDoesNotExist
	}
	store.Delete(key)
	// decrement open mtp count
	openCount := k.GetOpenMTPCount(ctx)
	openCount--

	// Set open MTP count
	k.SetOpenMTPCount(ctx, openCount)

	return nil
}

func (k Keeper) TakeFundPayment(ctx sdk.Context, returnAmount sdk.Int, returnAsset string, takePercentage sdk.Dec, fundAddr sdk.AccAddress, ammPool *ammtypes.Pool) (sdk.Int, error) {
	returnAmountDec := sdk.NewDecFromBigInt(returnAmount.BigInt())
	takeAmount := sdk.NewIntFromBigInt(takePercentage.Mul(returnAmountDec).TruncateInt().BigInt())

	if !takeAmount.IsZero() {
		takeCoins := sdk.NewCoins(sdk.NewCoin(returnAsset, sdk.NewIntFromBigInt(takeAmount.BigInt())))
		err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, ammPool.Address, fundAddr, takeCoins)
		if err != nil {
			return sdk.ZeroInt(), err
		}
	}
	return takeAmount, nil
}

func (k Keeper) SetMTP(ctx sdk.Context, mtp *types.MTP) error {
	store := ctx.KVStore(k.storeKey)
	count := k.GetMTPCount(ctx)
	openCount := k.GetOpenMTPCount(ctx)

	if mtp.Id == 0 {
		// increment global id count
		count++
		mtp.Id = count
		k.SetMTPCount(ctx, count)
		// increment open mtp count
		openCount++
		k.SetOpenMTPCount(ctx, openCount)
	}

	if err := mtp.Validate(); err != nil {
		return err
	}
	key := types.GetMTPKey(mtp.Address, mtp.Id)
	store.Set(key, k.cdc.MustMarshal(mtp))
	return nil
}

// Set Open MTP count
func (k Keeper) SetOpenMTPCount(ctx sdk.Context, count uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.OpenMTPCountPrefix, types.GetUint64Bytes(count))
}

// Set MTP count
func (k Keeper) SetMTPCount(ctx sdk.Context, count uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.MTPCountPrefix, types.GetUint64Bytes(count))
}

func (k Keeper) GetWhitelistAddressIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, types.WhitelistPrefix)
}

func (k Keeper) GetAllWhitelistedAddress(ctx sdk.Context) []string {
	var list []string
	iterator := k.GetWhitelistAddressIterator(ctx)
	defer func(iterator sdk.Iterator) {
		err := iterator.Close()
		if err != nil {
			panic(err)
		}
	}(iterator)

	for ; iterator.Valid(); iterator.Next() {
		list = append(list, (string)(iterator.Value()))
	}

	return list
}

func (k Keeper) GetMTPIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, types.MTPPrefix)
}

func (k Keeper) GetAllMTPs(ctx sdk.Context) []types.MTP {
	var mtpList []types.MTP
	iterator := k.GetMTPIterator(ctx)
	defer func(iterator sdk.Iterator) {
		err := iterator.Close()
		if err != nil {
			panic(err)
		}
	}(iterator)

	for ; iterator.Valid(); iterator.Next() {
		var mtp types.MTP
		bytesValue := iterator.Value()
		k.cdc.MustUnmarshal(bytesValue, &mtp)
		mtpList = append(mtpList, mtp)
	}
	return mtpList
}

func (k Keeper) GetMTPs(ctx sdk.Context, pagination *query.PageRequest) ([]*types.MTP, *query.PageResponse, error) {
	var mtpList []*types.MTP
	store := ctx.KVStore(k.storeKey)
	mtpStore := prefix.NewStore(store, types.MTPPrefix)

	if pagination == nil {
		pagination = &query.PageRequest{
			Limit: math.MaxUint64 - 1,
		}
	}

	pageRes, err := query.Paginate(mtpStore, pagination, func(key []byte, value []byte) error {
		var mtp types.MTP
		k.cdc.MustUnmarshal(value, &mtp)
		mtpList = append(mtpList, &mtp)
		return nil
	})

	return mtpList, pageRes, err
}

func (k Keeper) GetMTPsForPool(ctx sdk.Context, ammPoolId uint64, pagination *query.PageRequest) ([]*types.MTP, *query.PageResponse, error) {
	var mtps []*types.MTP

	store := ctx.KVStore(k.storeKey)
	mtpStore := prefix.NewStore(store, types.MTPPrefix)

	if pagination == nil {
		pagination = &query.PageRequest{
			Limit: math.MaxUint64 - 1,
		}
	}

	pageRes, err := query.FilteredPaginate(mtpStore, pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		var mtp types.MTP
		k.cdc.MustUnmarshal(value, &mtp)
		if accumulate && mtp.AmmPoolId == ammPoolId {
			mtps = append(mtps, &mtp)
			return true, nil
		}

		return false, nil
	})

	return mtps, pageRes, err
}

func (k Keeper) GetMTPsForAddress(ctx sdk.Context, mtpAddress sdk.Address, pagination *query.PageRequest) ([]*types.MTP, *query.PageResponse, error) {
	var mtps []*types.MTP

	store := ctx.KVStore(k.storeKey)
	mtpStore := prefix.NewStore(store, types.GetMTPPrefixForAddress(mtpAddress.String()))

	if pagination == nil {
		pagination = &query.PageRequest{
			Limit: types.MaxPageLimit,
		}
	}

	if pagination.Limit > types.MaxPageLimit {
		return nil, nil, status.Error(codes.InvalidArgument, fmt.Sprintf("page size greater than max %d", types.MaxPageLimit))
	}

	pageRes, err := query.Paginate(mtpStore, pagination, func(key []byte, value []byte) error {
		var mtp types.MTP
		k.cdc.MustUnmarshal(value, &mtp)
		mtps = append(mtps, &mtp)
		return nil
	})
	if err != nil {
		return nil, nil, err
	}

	return mtps, pageRes, nil
}

func (k Keeper) GetWhitelistedAddress(ctx sdk.Context, pagination *query.PageRequest) ([]string, *query.PageResponse, error) {
	var list []string
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.WhitelistPrefix)

	if pagination == nil {
		pagination = &query.PageRequest{
			Limit: math.MaxUint64 - 1,
		}
	}

	pageRes, err := query.Paginate(prefixStore, pagination, func(key []byte, value []byte) error {
		list = append(list, string(value))
		return nil
	})

	return list, pageRes, err
}

func (k Keeper) WhitelistAddress(ctx sdk.Context, address string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetWhitelistKey(address), []byte(address))
}

func (k Keeper) DewhitelistAddress(ctx sdk.Context, address string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetWhitelistKey(address))
}

// Set the margin hooks.
func (k *Keeper) SetHooks(gh types.MarginHooks) *Keeper {
	if k.hooks != nil {
		panic("cannot set margin hooks twice")
	}

	k.hooks = gh

	return k
}
