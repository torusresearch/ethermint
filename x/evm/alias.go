package evm

import (
	"github.com/torusresearch/ethermint/x/evm/keeper"
	"github.com/torusresearch/ethermint/x/evm/types"
)

// nolint
const (
	ModuleName = types.ModuleName
	StoreKey   = types.StoreKey
	RouterKey  = types.RouterKey
)

// nolint
var (
	NewKeeper = keeper.NewKeeper
	TxDecoder = types.TxDecoder
)

//nolint
type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState
)
