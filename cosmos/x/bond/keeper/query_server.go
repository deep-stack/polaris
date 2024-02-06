package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/berachain/polaris/cosmos/x/bond/types"
)

// TODO: Add remaining query methods

type queryServer struct {
	k Keeper
}

var _ types.QueryServer = queryServer{}

// NewQueryServerImpl returns an implementation of the module QueryServer.
func NewQueryServerImpl(k Keeper) types.QueryServer {
	return queryServer{k}
}

func (qs queryServer) Bonds(c context.Context, _ *types.QueryGetBondsRequest) (*types.QueryGetBondsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	resp, err := qs.k.ListBonds(ctx)
	if err != nil {
		return nil, err
	}

	return &types.QueryGetBondsResponse{Bonds: resp}, nil
}
