package icrc1

import (
	"fmt"
	"math/big"

	"github.com/aviate-labs/agent-go/candid/idl"
	"github.com/niccolofant/ic-arb/core/icp"
)

func (i *icrc1) Transfer(
	amount *big.Int,
	to icp.Principal,
	subaccount *[]byte,
	memo *[]byte,
) (TransferResponse, error) {
	amountToTransfer := new(big.Int).Sub(amount, i.Metadata().Fee)

	transferResult, err := i.api.Icrc1Transfer(TransferArgs{
		Amount: idl.NewBigNat(amountToTransfer),
		Memo:   memo,
		To: Account{
			Owner:      to.Raw(),
			Subaccount: subaccount,
		},
	})
	if err != nil {
		return TransferResponse{}, fmt.Errorf("failed to transfer %s tokens to %s: %w", amountToTransfer, to, err)
	}

	if transferResult.Err != nil {
		return TransferResponse{}, fmt.Errorf("failed to transfer %s tokens to %s: %w", amountToTransfer, to, transferResult.Err.Decode())
	}

	return TransferResponse{
		Amount:   amountToTransfer,
		BlockIdx: transferResult.Ok.BigInt(),
	}, nil
}
