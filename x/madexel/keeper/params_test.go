package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/MadeXeL/madexel/testutil/keeper"
	"github.com/MadeXeL/madexel/x/madexel/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MadexelKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
