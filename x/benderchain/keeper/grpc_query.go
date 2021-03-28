package keeper

import (
	"github.com/faddat/bender-chain/x/benderchain/types"
)

var _ types.QueryServer = Keeper{}
