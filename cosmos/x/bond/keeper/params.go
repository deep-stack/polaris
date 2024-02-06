package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/berachain/polaris/cosmos/x/bond/types"
)

// GetMaxBondAmount max bond amount
func (k Keeper) GetMaxBondAmount(ctx sdk.Context) (res sdk.Coin) {
	// TODO: Implement
	// k.paramSubspace.Get(ctx, types.ParamStoreKeyMaxBondAmount, &res)
	return sdk.NewCoin(sdk.DefaultBondDenom, types.DefaultMaxBondAmountTokens)
}

// GetParams - Get all parameter as types.Params.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	getMaxBondAmount := k.GetMaxBondAmount(ctx)
	return types.Params{MaxBondAmount: getMaxBondAmount}
}

// SetParams - set the params.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	// TODO: Implement
	// k.paramSubspace.SetParamSet(ctx, &params)
}
