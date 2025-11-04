package icrc1

import (
	"fmt"

	"github.com/niccolofant/ic-arb/core/icp"
)

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

func (i *icrc1) Metadata() icp.TokenMetadata {
	return i.metadata
}

func (i *icrc1) setMetadata() error {
	metadataResult, err := i.api.Icrc1Metadata()
	if err != nil {
		return err
	}

	if metadataResult == nil {
		return fmt.Errorf("invalid metadata")
	}

	supportedStandardsResult, err := i.api.Icrc1SupportedStandards()
	if err != nil {
		return err
	}

	if supportedStandardsResult == nil {
		return fmt.Errorf("invalid supported standards")
	}

	metadata := icp.TokenMetadata{
		Standard: icp.TokenStandardICRC1,
	}

	for _, meta := range *metadataResult {
		switch meta.Field0 {
		case "icrc1:fee":
			metadata.Fee = meta.Field1.Nat.BigInt()
		case "icrc1:name":
			metadata.Name = *meta.Field1.Text
		case "icrc1:symbol":
			metadata.Symbol = *meta.Field1.Text
		case "icrc1:decimals":
			metadata.Decimals = int(meta.Field1.Nat.BigInt().Int64())
		}
	}

	for _, standard := range *supportedStandardsResult {
		if standard.Name == icp.TokenStandardICRC2.String() {
			metadata.Standard = icp.TokenStandardICRC2
			break
		}
	}

	i.metadata = metadata
	return nil
}
