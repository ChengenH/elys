package cli_test

import (
	assetprofilemoduletypes "github.com/elys-network/elys/x/assetprofile/types"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/elys-network/elys/testutil/network"
	"github.com/elys-network/elys/testutil/nullify"
	"github.com/elys-network/elys/x/tradeshield/types"
)

func networkWithPendingPerpetualOrderObjects(t *testing.T, n int) (*network.Network, []types.PerpetualOrder) {
	t.Helper()
	cfg := network.DefaultConfig(t.TempDir())
	state := types.GenesisState{}
	for i := 0; i < n; i++ {
		pendingPerpetualOrder := types.PerpetualOrder{
			OrderId: uint64(i),
		}
		nullify.Fill(&pendingPerpetualOrder)
		state.PendingPerpetualOrderList = append(state.PendingPerpetualOrderList, pendingPerpetualOrder)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf

	assetProfileGenesisState := assetprofilemoduletypes.DefaultGenesis()
	usdcEntry := assetprofilemoduletypes.Entry{
		BaseDenom:                "uusdc",
		Decimals:                 6,
		Denom:                    "uusdc",
		Path:                     "",
		IbcChannelId:             "",
		IbcCounterpartyChannelId: "",
		DisplayName:              "",
		DisplaySymbol:            "",
		Network:                  "",
		Address:                  "",
		ExternalSymbol:           "",
		TransferLimit:            "",
		Permissions:              nil,
		UnitDenom:                "",
		IbcCounterpartyDenom:     "",
		IbcCounterpartyChainId:   "",
		Authority:                "",
		CommitEnabled:            true,
		WithdrawEnabled:          true,
	}
	assetProfileGenesisState.EntryList = []assetprofilemoduletypes.Entry{usdcEntry}
	buf, err = cfg.Codec.MarshalJSON(assetProfileGenesisState)
	require.NoError(t, err)
	cfg.GenesisState[assetprofilemoduletypes.ModuleName] = buf
	return network.New(t, cfg), state.PendingPerpetualOrderList
}

// TODO: Add tests for the CLI queries in query task
// func TestShowPendingPerpetualOrder(t *testing.T) {
// 	net, objs := networkWithPendingPerpetualOrderObjects(t, 2)

// 	ctx := net.Validators[0].ClientCtx
// 	common := []string{
// 		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
// 	}
// 	tests := []struct {
// 		desc string
// 		id   string
// 		args []string
// 		err  error
// 		obj  types.PerpetualOrder
// 	}{
// 		{
// 			desc: "found",
// 			id:   fmt.Sprintf("%d", objs[0].OrderId),
// 			args: common,
// 			obj:  objs[0],
// 		},
// 		{
// 			desc: "not found",
// 			id:   "not_found",
// 			args: common,
// 			err:  status.Error(codes.NotFound, "not found"),
// 		},
// 	}
// 	for _, tc := range tests {
// 		t.Run(tc.desc, func(t *testing.T) {
// 			args := []string{tc.id}
// 			args = append(args, tc.args...)
// 			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowPendingPerpetualOrder(), args)
// 			if tc.err != nil {
// 				stat, ok := status.FromError(tc.err)
// 				require.True(t, ok)
// 				require.ErrorIs(t, stat.Err(), tc.err)
// 			} else {
// 				require.NoError(t, err)
// 				var resp types.QueryGetPendingPerpetualOrderResponse
// 				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
// 				require.NotNil(t, resp.PendingPerpetualOrder)
// 				require.Equal(t,
// 					nullify.Fill(&tc.obj),
// 					nullify.Fill(&resp.PendingPerpetualOrder),
// 				)
// 			}
// 		})
// 	}
// }

// func TestListPendingPerpetualOrder(t *testing.T) {
// 	net, objs := networkWithPendingPerpetualOrderObjects(t, 5)

// 	ctx := net.Validators[0].ClientCtx
// 	request := func(next []byte, offset, limit uint64, total bool) []string {
// 		args := []string{
// 			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
// 		}
// 		if next == nil {
// 			args = append(args, fmt.Sprintf("--%s=%d", flags.FlagOffset, offset))
// 		} else {
// 			args = append(args, fmt.Sprintf("--%s=%s", flags.FlagPageKey, next))
// 		}
// 		args = append(args, fmt.Sprintf("--%s=%d", flags.FlagLimit, limit))
// 		if total {
// 			args = append(args, fmt.Sprintf("--%s", flags.FlagCountTotal))
// 		}
// 		return args
// 	}
// 	t.Run("ByOffset", func(t *testing.T) {
// 		step := 2
// 		for i := 0; i < len(objs); i += step {
// 			args := request(nil, uint64(i), uint64(step), false)
// 			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListPendingPerpetualOrder(), args)
// 			require.NoError(t, err)
// 			var resp types.QueryAllPendingPerpetualOrderResponse
// 			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
// 			require.LessOrEqual(t, len(resp.PendingPerpetualOrder), step)
// 			require.Subset(t,
// 				nullify.Fill(objs),
// 				nullify.Fill(resp.PendingPerpetualOrder),
// 			)
// 		}
// 	})
// 	t.Run("ByKey", func(t *testing.T) {
// 		step := 2
// 		var next []byte
// 		for i := 0; i < len(objs); i += step {
// 			args := request(next, 0, uint64(step), false)
// 			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListPendingPerpetualOrder(), args)
// 			require.NoError(t, err)
// 			var resp types.QueryAllPendingPerpetualOrderResponse
// 			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
// 			require.LessOrEqual(t, len(resp.PendingPerpetualOrder), step)
// 			require.Subset(t,
// 				nullify.Fill(objs),
// 				nullify.Fill(resp.PendingPerpetualOrder),
// 			)
// 			next = resp.Pagination.NextKey
// 		}
// 	})
// 	t.Run("Total", func(t *testing.T) {
// 		args := request(nil, 0, uint64(len(objs)), true)
// 		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListPendingPerpetualOrder(), args)
// 		require.NoError(t, err)
// 		var resp types.QueryAllPendingPerpetualOrderResponse
// 		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
// 		require.NoError(t, err)
// 		require.Equal(t, len(objs), int(resp.Pagination.Total))
// 		require.ElementsMatch(t,
// 			nullify.Fill(objs),
// 			nullify.Fill(resp.PendingPerpetualOrder),
// 		)
// 	})
// }
