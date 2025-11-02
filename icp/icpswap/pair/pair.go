package pair

import (
	"fmt"

	"github.com/niccolofant/ic-arb/icp"
)

var _ Pair = (*pair)(nil)

type Pair interface {
	icp.Dex
	icp.DexQuote
	icp.DexSwap
	icp.DexNotAggregated
}

type pair struct {
	api        API
	canisterID icp.Principal
	token0     icp.Token
	token1     icp.Token
}

func NewWithMetadata(
	agent *icp.Agent,
	canisterID icp.Principal,
	token0, token1 icp.Token,
) (*pair, error) {
	api, err := NewAPI(canisterID, agent)
	if err != nil {
		return nil, fmt.Errorf("failed to create api client for %s: %w", canisterID, err)
	}

	return &pair{
		api:        api,
		canisterID: canisterID,
		token0:     token0,
		token1:     token1,
	}, nil
}

func (p *pair) CanisterID() icp.Principal {
	return p.canisterID
}

func (p *pair) Equal(other icp.Canister) bool {
	return p.CanisterID().Equal(other.CanisterID())
}

func (p *pair) Token0() icp.Token {
	return p.token0
}

func (p *pair) Token1() icp.Token {
	return p.token1
}

func (p *pair) Type() icp.DexType {
	return icp.DexTypeIcpSwap
}
