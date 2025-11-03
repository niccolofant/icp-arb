package ntnswap

import (
	"fmt"
	"math/big"

	"github.com/niccolofant/ic-arb/core/icp"
	"github.com/niccolofant/ic-arb/core/icp/icrc1"
)

func (n *ntnswap) Deposit(token icp.Token, amount *big.Int) (*big.Int, error) {
	subAccount := n.BotSubaccount()

	fromIcrc1, isIcrc1 := token.(icrc1.ICRC1)
	if !isIcrc1 {
		return nil, fmt.Errorf("token does not implement icrc1 interface")
	}

	transferResult, err := fromIcrc1.Transfer(
		amount,
		n.CanisterID(),
		&subAccount,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to deposit: %w", err)
	}

	txToFind := Tx{
		Token:          token,
		FromSubaccount: &n.botSubaccount,
		To:             n.CanisterID(),
		Amount:         transferResult.Amount,
	}

	tx, err := n.waitForTxCompletion(txToFind)
	if err != nil {
		return nil, fmt.Errorf("failed to deposit: %w", err)
	}

	return new(big.Int).Sub(tx.Amount, token.Metadata().Fee), nil
}
