package bond

import "github.com/berachain/polaris/cosmos/x/bond/types"

// DefaultGenesisState sets default evm genesis state with empty accounts and default params and
// chain config values.
func DefaultGenesisState() *types.GenesisState {
	return &types.GenesisState{
		Params: types.DefaultParams(),
		Bonds:  []*types.Bond{},
	}
}

func NewGenesisState(params types.Params, bonds []*types.Bond) *types.GenesisState {
	return &types.GenesisState{
		Params: params,
		Bonds:  bonds,
	}
}

// ValidateGenesis - validating the genesis data
func ValidateGenesis(data types.GenesisState) error {
	err := data.Params.Validate()
	if err != nil {
		return err
	}

	return nil
}
