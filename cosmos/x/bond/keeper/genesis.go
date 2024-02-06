package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/berachain/polaris/cosmos/x/bond/types"
)

// InitGenesis initializes the module state from a genesis state.
func (k *Keeper) InitGenesis(ctx sdk.Context, data *types.GenesisState) error {
	k.SetParams(ctx, data.Params)

	// Save bonds in store.
	for _, bond := range data.Bonds {
		if err := k.Bonds.Set(ctx, bond.Id, *bond); err != nil {
			return err
		}
	}

	return nil
}

// ExportGenesis exports the module state to a genesis state.
func (k *Keeper) ExportGenesis(ctx sdk.Context) (*types.GenesisState, error) {
	params := k.GetParams(ctx)

	bonds, err := k.ListBonds(ctx)
	if err != nil {
		return nil, err
	}

	return &types.GenesisState{Params: params, Bonds: bonds}, nil
}
