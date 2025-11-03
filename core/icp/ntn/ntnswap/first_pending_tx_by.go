package ntnswap

import (
	"bytes"
	"fmt"

	"github.com/niccolofant/ic-arb/core/icp"
)

func (n *ntnswap) FindPendingTx(tx Tx) (*Tx, error) {
	tokensPendingTxs, err := n.api.GetPendingTransactions()
	if err != nil {
		return nil, fmt.Errorf("failed to get pending transactions: %w", err)
	}

	var pendingTxsForToken []TransactionShared

	for _, tokenTxGroup := range *tokensPendingTxs {
		if tokenTxGroup.Id.Equal(tx.Token.CanisterID().Raw()) {
			pendingTxsForToken = tokenTxGroup.Transactions
			break
		}
	}

	if len(pendingTxsForToken) == 0 {
		return nil, nil
	}

	for _, pendingTx := range pendingTxsForToken {
		match := true

		// amount
		if tx.Amount != nil && pendingTx.Amount.BigInt().Cmp(tx.Amount) != 0 {
			match = false
		}

		// recipient
		if !tx.To.Raw().Equal(pendingTx.To.Icrc.Owner) {
			match = false
		}

		// subaccount
		if tx.FromSubaccount != nil && pendingTx.FromSubaccount != nil {
			if !bytes.Equal(*tx.FromSubaccount, *pendingTx.FromSubaccount) {
				match = false
			}
		} else if tx.FromSubaccount != nil || pendingTx.FromSubaccount != nil {
			// One is nil, other is not
			match = false
		}

		if match {
			return &Tx{
				Token:          tx.Token,
				Memo:           pendingTx.Memo,
				FromSubaccount: pendingTx.FromSubaccount,
				To:             icp.NewPrincipal(pendingTx.To.Icrc.Owner),
				Amount:         pendingTx.Amount.BigInt(),
			}, nil
		}
	}

	return nil, nil
}
