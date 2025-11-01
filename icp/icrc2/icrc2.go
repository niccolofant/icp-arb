package icrc2

import (
	"fmt"
	"math/big"

	"github.com/niccolofant/ic-arb/icp"
)

type ICRC2 interface {
	icp.Token
	Approve(spender icp.Principal, amount *big.Int) error
}

type icrc2 struct {
	api        API
	canisterID icp.Principal
	metadata   icp.Metadata
}

func NewWithMetadata(
	agent *icp.Agent,
	canisterID icp.Principal,
	metadata icp.Metadata,
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

func (i *icrc2) Metadata() icp.Metadata {
	return i.metadata
}
