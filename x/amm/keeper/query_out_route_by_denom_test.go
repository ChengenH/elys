package keeper_test

import (
	"cosmossdk.io/math"
	"github.com/elys-network/elys/x/amm/types"
	ptypes "github.com/elys-network/elys/x/parameter/types"
)

func (suite *AmmKeeperTestSuite) TestOutRouteByDenom() {
	testCases := []struct {
		name                 string
		prerequisiteFunction func()
		postValidateFunction func()
	}{
		{
			"nil request",
			func() {
				suite.ResetSuite()
			},
			func() {
				_, err := suite.app.AmmKeeper.OutRouteByDenom(suite.ctx, nil)
				suite.Require().Error(err)
			},
		},
		{
			"base currency does not exist in asset profile",
			func() {
				suite.ResetSuite()
				suite.app.AssetprofileKeeper.RemoveEntry(suite.ctx, ptypes.BaseCurrency)
			},
			func() {
				_, err := suite.app.AmmKeeper.OutRouteByDenom(suite.ctx, &types.QueryOutRouteByDenomRequest{})
				suite.Require().Error(err)
			},
		},
		{
			"denom in and denom out are the same",
			func() {
				suite.ResetSuite()
			},
			func() {
				_, err := suite.app.AmmKeeper.OutRouteByDenom(suite.ctx, &types.QueryOutRouteByDenomRequest{
					DenomIn:  ptypes.ATOM,
					DenomOut: ptypes.ATOM,
				})
				suite.Require().Error(err)
			},
		},
		{
			"valid request",
			func() {
				suite.ResetSuite()
				suite.SetupCoinPrices()
			},
			func() {
				addr := suite.AddAccounts(1, nil)[0]

				amount := math.NewInt(100000000000)
				_ = suite.CreateNewAmmPool(addr, true, math.LegacyZeroDec(), math.LegacyZeroDec(), ptypes.ATOM, amount.MulRaw(10), amount.MulRaw(10))

				_, err := suite.app.AmmKeeper.OutRouteByDenom(suite.ctx, &types.QueryOutRouteByDenomRequest{
					DenomIn:  ptypes.BaseCurrency,
					DenomOut: ptypes.ATOM,
				})
				suite.Require().NoError(err)
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			tc.prerequisiteFunction()
			tc.postValidateFunction()
		})
	}
}