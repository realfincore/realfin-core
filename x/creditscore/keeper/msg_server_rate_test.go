package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "realfin/testutil/keeper"
	"realfin/x/creditscore/keeper"
	"realfin/x/creditscore/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestRateMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.CreditscoreKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateRate{Creator: creator,
			Symbol: strconv.Itoa(i),
		}
		_, err := srv.CreateRate(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetRate(ctx,
			expected.Symbol,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestRateMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateRate
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateRate{Creator: creator,
				Symbol: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateRate{Creator: "B",
				Symbol: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateRate{Creator: creator,
				Symbol: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.CreditscoreKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateRate{Creator: creator,
				Symbol: strconv.Itoa(0),
			}
			_, err := srv.CreateRate(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateRate(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetRate(ctx,
					expected.Symbol,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestRateMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteRate
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteRate{Creator: creator,
				Symbol: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteRate{Creator: "B",
				Symbol: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteRate{Creator: creator,
				Symbol: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.CreditscoreKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateRate(ctx, &types.MsgCreateRate{Creator: creator,
				Symbol: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteRate(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetRate(ctx,
					tc.request.Symbol,
				)
				require.False(t, found)
			}
		})
	}
}
