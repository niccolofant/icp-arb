package icrc2

import (
	"fmt"
	"math/big"

	"github.com/niccolofant/ic-arb/icp"
)

var _ ICRC2 = (*icrc2)(nil)

type ICRC2 interface {
	icp.Token
	Approve(spender icp.Principal, amount *big.Int) error
}

type icrc2 struct {
	api        API
	canisterID icp.Principal
	metadata   icp.TokenMetadata
}

func NewWithMetadata(
	agent *icp.Agent,
	canisterID icp.Principal,
	metadata icp.TokenMetadata,
) (*icrc2, error) {
	api, err := NewAPI(canisterID, agent)
	if err != nil {
		return nil, fmt.Errorf("failed to create api client for %s: %w", canisterID, err)
	}

	return &icrc2{
		api:        api,
		canisterID: canisterID,
		metadata:   metadata,
	}, nil
}

func (i *icrc2) CanisterID() icp.Principal {
	return i.canisterID
}

func (i *icrc2) Equal(other icp.Canister) bool {
	return i.CanisterID().Equal(other.CanisterID())
}

func (i *icrc2) Metadata() icp.TokenMetadata {
	return i.metadata
}
