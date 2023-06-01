package keeper

import (
	"github.com/Vingurzhou/checkers/x/leaderboard/types"
)

var _ types.QueryServer = Keeper{}
