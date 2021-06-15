package keeper

import (
	"github.com/hardikpatel047/sensibleschain/x/sensibleschain/types"
)

var _ types.QueryServer = Keeper{}
