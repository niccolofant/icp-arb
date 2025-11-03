package icrc2

import (
	"fmt"
	"math/big"

	"github.com/niccolofant/ic-arb/core/icp"
	"github.com/niccolofant/ic-arb/core/icp/icrc1"
)

var _ ICRC2 = (*icrc2)(nil)

type ICRC2 interface {
	icrc1.ICRC1
	Approve(spender icp.Principal, amount *big.Int) error
}

type icrc2 struct {
	icrc1.ICRC1
	api API
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

	icrc1, err := icrc1.NewWithMetadata(
		agent,
		canisterID,
		metadata,
	)
	if err != nil {
		return nil, err
	}

	return &icrc2{
		ICRC1: icrc1,
		api:   api,
	}, nil
}

func (i *icrc2) CanisterID() icp.Principal {
	return i.ICRC1.CanisterID()
}

func (i *icrc2) Equal(other icp.Canister) bool {
	return i.CanisterID().Equal(other.CanisterID())
}

func (i *icrc2) Metadata() icp.TokenMetadata {
	return i.ICRC1.Metadata()
}

func (i *icrc2) String() string {
	return i.Metadata().Symbol
}
