package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/Vingurzhou/checkers/testutil/keeper"
	"github.com/Vingurzhou/checkers/x/leaderboard/keeper"
	"github.com/Vingurzhou/checkers/x/leaderboard/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.LeaderboardKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
