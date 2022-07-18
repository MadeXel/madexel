package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/MadeXeL/madexel/x/madexel/types"
    "github.com/MadeXeL/madexel/x/madexel/keeper"
    keepertest "github.com/MadeXeL/madexel/testutil/keeper"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.MadexelKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
