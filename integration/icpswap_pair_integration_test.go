package integration

import (
	"math/big"
	"net/http"
	"testing"

	"github.com/niccolofant/ic-arb/icp"
	icpswap_pair "github.com/niccolofant/ic-arb/icp/icpswap/pair"
	"github.com/niccolofant/ic-arb/icp/icrc2"
	"github.com/stretchr/testify/assert"
)

var pairCanisterId = icp.MustDecodePrincipal("xmiu5-jqaaa-aaaag-qbz7q-cai")

func TestNewWithMetadata_Integration(t *testing.T) {
	id, err := icp.LoadIntentity("../identity.pem")
	assert.NoError(t, err, "failed to load identity")

	agent, err := icp.NewAgent(id, http.DefaultClient)
	assert.NoError(t, err, "failed to create agent")

	icpToken, err := icrc2.NewWithMetadata(
		agent,
		icp.LedgerPrincipal,
		icp.TokenMetadata{
			Name:   "ICP",
			Symbol: "ICP",
			Fee:    big.NewInt(0.0001 * icp.E8S),
		},
	)
	assert.NoError(t, err, "failed to create icp token")

	ckbtcToken, err := icrc2.NewWithMetadata(
		agent,
		icp.MustDecodePrincipal("mxzaz-hqaaa-aaaar-qaada-cai"),
		icp.TokenMetadata{
			Name:   "ckBTC",
			Symbol: "ckBTC",
			Fee:    big.NewInt(10),
		},
	)
	assert.NoError(t, err, "failed to create ckbtc token")

	p, err := icpswap_pair.NewWithMetadata(
		agent,
		pairCanisterId,
		ckbtcToken,
		icpToken,
	)
	assert.NoError(t, err, "failed to create icp/ckbtc pair")

	assert.Equal(t, pairCanisterId, p.CanisterID())
	assert.Equal(t, ckbtcToken, p.Token0())
	assert.Equal(t, icpToken, p.Token1())
	assert.Equal(t, icp.DexTypeIcpSwap, p.Type())
}
