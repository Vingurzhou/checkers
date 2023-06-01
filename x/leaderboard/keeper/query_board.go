package keeper

import (
	"context"

	"github.com/Vingurzhou/checkers/x/leaderboard/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Board(goCtx context.Context, req *types.QueryGetBoardRequest) (*types.QueryGetBoardResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetBoard(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetBoardResponse{Board: val}, nil
}
