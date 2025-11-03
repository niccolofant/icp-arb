package ntnswap

import (
	"fmt"
	"time"

	"github.com/niccolofant/ic-arb/core/icp"
)

var _ Ntnswap = (*ntnswap)(nil)

type Ntnswap interface {
	icp.Dex
	icp.DexQuote
	icp.DexOneStepQuote
	icp.DexSwap
	icp.DexOneStepSwap
	FindPendingTx(tx Tx) (*Tx, error)
}

type ntnswap struct {
	api           API
	canisterID    icp.Principal
	botSubaccount []byte
	pollingCfg    struct {
		delay       time.Duration
		maxAttempts int
	}
}

func NewWithMetadata(agent *icp.Agent) (*ntnswap, error) {
	canisterID := icp.MustDecodePrincipal("togwv-zqaaa-aaaal-qr7aa-cai")

	subaccount := []byte{
		227, 123, 208, 208, 254, 48, 150, 133, 189,
		99, 31, 102, 158, 226, 147, 129, 172, 33, 159,
		202, 97, 122, 10, 93, 147, 210, 142, 28, 253, 82, 202, 108,
	}

	api, err := NewAPI(canisterID, agent)
	if err != nil {
		return nil, fmt.Errorf("failed to create api client for %s: %w", canisterID, err)
	}

	return &ntnswap{
		api:           api,
		canisterID:    canisterID,
		botSubaccount: subaccount,
		pollingCfg: struct {
			delay       time.Duration
			maxAttempts int
		}{
			delay:       1 * time.Second,
			maxAttempts: 60,
		},
	}, nil
}

func (n *ntnswap) CanisterID() icp.Principal {
	return n.canisterID
}

func (n *ntnswap) BotSubaccount() []byte {
	return n.botSubaccount
}

func (n *ntnswap) Equal(other icp.Canister) bool {
	return n.CanisterID().Equal(other.CanisterID())
}

func (n *ntnswap) Type() icp.DexType {
	return icp.DexTypeNtn
}

func (n *ntnswap) waitForTxCompletion(tx Tx) (*Tx, error) {
	var pendingTx *Tx

	for i := 0; i < n.pollingCfg.maxAttempts; i++ {
		time.Sleep(n.pollingCfg.delay)

		tx, err := n.FindPendingTx(tx)
		if err != nil {
			continue
		}

		// tx is pending
		if tx != nil {
			pendingTx = tx
			continue
		}

		// tx was pending, now it's gone aka has been processed
		if pendingTx != nil {
			return pendingTx, nil
		}
	}

	return nil, fmt.Errorf("timeout waiting for completion of tx")
}
