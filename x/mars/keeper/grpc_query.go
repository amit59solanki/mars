package keeper

import (
	"github.com/fadeev/mars/x/mars/types"
)

var _ types.QueryServer = Keeper{}
