package icrc2

import (
	"fmt"
	"math/big"

	"github.com/aviate-labs/agent-go/candid/idl"
	"github.com/niccolofant/ic-arb/core/icp"
)

func (i *icrc2) Approve(
	spender icp.Principal,
	amount *big.Int,
) error {
	approveResult, err := i.api.Icrc2Approve(ApproveArgs{
		Spender: Account{
			Owner: spender.Raw(),
		},
		Amount: idl.NewBigNat(amount),
	})
	if err != nil {
		return fmt.Errorf("failed to approve %s tokens for %s: %w", amount, spender, err)
	}

	if approveResult.Err != nil {
		return fmt.Errorf("failed to approve %s tokens for %s: %w", amount, spender, approveResult.Err.Decode())
	}

	return nil
}
