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
) (*big.Int, error) {
	amountToTransfer := new(big.Int).Sub(amount, i.Metadata().Fee)

	transferResult, err := i.api.Icrc1Transfer(TransferArgs{
		Amount: idl.NewBigNat(amountToTransfer),
		To: Account{
			Owner:      to.Raw(),
			Subaccount: subaccount,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to transfer %s tokens to %s: %w", amountToTransfer, to, err)
	}

	if transferResult.Err != nil {
		return nil, fmt.Errorf("failed to transfer %s tokens to %s: %w", amountToTransfer, to, transferResult.Err.Decode())
	}

	return amountToTransfer, nil
}
