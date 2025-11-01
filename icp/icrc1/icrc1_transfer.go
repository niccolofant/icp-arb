package icrc1

import (
	"fmt"
	"math/big"

	"github.com/aviate-labs/agent-go/candid/idl"
	"github.com/niccolofant/ic-arb/icp"
)

func (i *icrc1) Transfer(
	amount *big.Int,
	to icp.Principal,
	subaccount *[]byte,
) error {
	transferResult, err := i.api.Icrc1Transfer(TransferArgs{
		Amount: idl.NewBigNat(amount),
		To: Account{
			Owner:      to.Raw(),
			Subaccount: subaccount,
		},
	})
	if err != nil {
		return fmt.Errorf("failed to transfer %s tokens to %s: %w", amount, to, err)
	}

	if transferResult.Err != nil {
		return fmt.Errorf("failed to transfer %s tokens to %s: %w", amount, to, transferResult.Err.Decode())
	}

	return nil
}
