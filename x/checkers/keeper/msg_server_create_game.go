package keeper

import (
	"context"
	"strconv"

	"github.com/Vingurzhou/checkers/x/checkers/rules"
	"github.com/Vingurzhou/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateGame(goCtx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	systemInfo, found := k.Keeper.GetSystemInfo(ctx)
	if !found {
		panic("SystemInfo not found")
	}
	newIndex := strconv.FormatUint(systemInfo.NextId, 10)

	newGame := rules.New()
	storedGame := types.StoredGame{
		Index:       newIndex,
		Board:       newGame.String(),
		Turn:        rules.PieceStrings[newGame.Turn],
		Black:       msg.Black,
		Red:         msg.Red,
		Winner:      rules.PieceStrings[rules.NO_PLAYER],
		Deadline:    types.FormatDeadline(types.GetNextDeadline(ctx)),
		MoveCount:   uint64(0),
		BeforeIndex: types.NoFifoIndex,
		AfterIndex:  types.NoFifoIndex,
		Wager:       msg.Wager,
		Denom:       msg.Denom,
	}

	err := storedGame.Validate()
	if err != nil {
		return nil, err
	}

	k.Keeper.SendToFifoTail(ctx, &storedGame, &systemInfo)
	k.Keeper.SetStoredGame(ctx, storedGame)
	systemInfo.NextId++
	k.Keeper.SetSystemInfo(ctx, systemInfo)

	ctx.GasMeter().ConsumeGas(types.CreateGameGas, "Create game")

	ctx.EventManager().EmitTypedEvent(
		&types.MsgCreateGame{
			Creator: msg.Creator,
			Black:   msg.Black,
			Red:     msg.Red,
			Wager:   msg.Wager,
			Denom:   msg.Denom,
		},
	)

	return &types.MsgCreateGameResponse{
		GameIndex: newIndex,
	}, nil
}
