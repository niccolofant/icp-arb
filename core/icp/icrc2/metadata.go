package icrc2

import (
	"fmt"

	"github.com/niccolofant/ic-arb/core/icp"
	"github.com/niccolofant/ic-arb/core/icp/icrc1"
)

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

func (i *icrc2) Metadata() icp.TokenMetadata {
	return i.ICRC1.Metadata()
}
