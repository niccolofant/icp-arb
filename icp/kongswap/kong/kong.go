package kong

import (
	"fmt"

	"github.com/niccolofant/ic-arb/icp"
)

var _ Kong = (*kong)(nil)

type Kong interface {
	icp.Dex
	icp.DexQuote
	icp.DexSwap
}

type kong struct {
	api        API
	canisterID icp.Principal
}

func NewWithMetadata(agent *icp.Agent) (*kong, error) {
	canisterID := icp.MustDecodePrincipal("2ipq2-uqaaa-aaaar-qailq-cai")

	api, err := NewAPI(canisterID, agent)
	if err != nil {
		return nil, fmt.Errorf("failed to create api client for %s: %w", canisterID, err)
	}

	return &kong{
		api:        api,
		canisterID: canisterID,
	}, nil
}

func (k *kong) CanisterID() icp.Principal {
	return k.canisterID
}

func (k *kong) Equal(other icp.Canister) bool {
	return k.CanisterID().Equal(other.CanisterID())
}

func (k *kong) Type() icp.DexType {
	return icp.DexTypeKongSwap
}
