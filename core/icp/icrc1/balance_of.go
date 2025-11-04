package icrc1

import (
	"fmt"
	"math/big"

	"github.com/niccolofant/ic-arb/core/icp"
)

func (i *icrc1) BalanceOf(account icp.Principal) (*big.Int, error) {
	balanceOfResult, err := i.api.Icrc1BalanceOf(Account{
		Owner: account.Raw(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get balance of %s token for account %s: %w", i, account, err)
	}

	return balanceOfResult.BigInt(), nil
}
