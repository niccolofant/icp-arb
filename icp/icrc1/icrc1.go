package icrc1

import (
	"fmt"
	"math/big"

	"github.com/niccolofant/ic-arb/icp"
)

var _ ICRC1 = (*icrc1)(nil)

type ICRC1 interface {
	icp.Token
	Transfer(amount *big.Int, to icp.Principal, subaccount *[]byte) (*big.Int, error)
}

type icrc1 struct {
	api        API
	canisterID icp.Principal
	metadata   icp.TokenMetadata
}

func NewWithMetadata(
	agent *icp.Agent,
	canisterID icp.Principal,
	metadata icp.TokenMetadata,
) (*icrc1, error) {
	api, err := NewAPI(canisterID, agent)
	if err != nil {
		return nil, fmt.Errorf("failed to create api client for %s: %w", canisterID, err)
	}

	return &icrc1{
		api:        api,
		canisterID: canisterID,
		metadata:   metadata,
	}, nil
}

func (i *icrc1) CanisterID() icp.Principal {
	return i.canisterID
}

func (i *icrc1) Equal(other icp.Canister) bool {
	return i.CanisterID().Equal(other.CanisterID())
}

func (i *icrc1) Metadata() icp.TokenMetadata {
	return i.metadata
}

func (i *icrc1) String() string {
	return i.metadata.Symbol
}
