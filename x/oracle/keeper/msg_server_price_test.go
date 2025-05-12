package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "realfin/testutil/keeper"
	"realfin/x/oracle/keeper"
	"realfin/x/oracle/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestPriceMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.OracleKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreatePrice{Creator: creator,
			Symbol: strconv.Itoa(i),
		}
		_, err := srv.CreatePrice(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetPrice(ctx,
			expected.Symbol,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestPriceMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdatePrice
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdatePrice{Creator: creator,
				Symbol: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdatePrice{Creator: "B",
				Symbol: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdatePrice{Creator: creator,
				Symbol: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.OracleKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreatePrice{Creator: creator,
				Symbol: strconv.Itoa(0),
			}
			_, err := srv.CreatePrice(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdatePrice(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetPrice(ctx,
					expected.Symbol,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestPriceMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeletePrice
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeletePrice{Creator: creator,
				Symbol: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeletePrice{Creator: "B",
				Symbol: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeletePrice{Creator: creator,
				Symbol: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.OracleKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreatePrice(ctx, &types.MsgCreatePrice{Creator: creator,
				Symbol: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeletePrice(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetPrice(ctx,
					tc.request.Symbol,
				)
				require.False(t, found)
			}
		})
	}
}
