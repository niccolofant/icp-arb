package icrc1

import (
	"github.com/niccolofant/ic-arb/icp"
)

type ICRC1 struct {
	agent      *icp.Agent
	canisterID icp.Principal
	metadata   Metadata
}

func NewWithMetadata(
	agent *icp.Agent,
	canisterID icp.Principal,
	metadata Metadata,
) (*ICRC1, error) {
	return &ICRC1{
		agent:      agent,
		canisterID: canisterID,
		metadata:   metadata,
	}, nil
}
