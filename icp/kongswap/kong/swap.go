package kong

import (
	"fmt"
	"math/big"

	"github.com/aviate-labs/agent-go/candid/idl"
	"github.com/niccolofant/ic-arb/icp"
	"github.com/niccolofant/ic-arb/icp/icrc1"
)

func (k *kong) Swap(from, to icp.Token, amountIn, amountOutMin *big.Int) (*big.Int, error) {
	if from.Metadata().Standard == icp.TokenStandardICRC1 {
		return k.swapIcrc1(from, to, amountIn, amountOutMin)
	}

	return k.swapIcrc2(from, to, amountIn, amountOutMin)
}

func (k *kong) swapIcrc1(from, to icp.Token, amountIn, amountOutMin *big.Int) (*big.Int, error) {
	fromIcrc1, isIcrc1 := from.(icrc1.ICRC1)
	if !isIcrc1 {
		return nil, fmt.Errorf("from token does not implement icrc1 interface")
	}

	transferResult, err := fromIcrc1.Transfer(amountIn, k.CanisterID(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to transfer %s %s tokens to kong: %w", amountIn, fromIcrc1, err)
	}

	blockIdx := idl.NewBigNat(transferResult.BlockIdx)
	receiveAmount := idl.NewBigNat(amountOutMin)
	fromFees := from.Metadata().Fee
	maxSlippage := 100.0

	swapResult, err := k.api.Swap(SwapArgs{
		PayToken:      fmt.Sprintf("IC.%s", from.CanisterID()),
		PayAmount:     idl.NewBigNat(new(big.Int).Sub(amountIn, fromFees)),
		ReceiveToken:  fmt.Sprintf("IC.%s", to.CanisterID()),
		ReceiveAmount: &receiveAmount,
		MaxSlippage:   &maxSlippage,
		PayTxId: &TxId{
			BlockIndex: &blockIdx,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to swap %s %s tokens to %s tokens: %w", amountIn, from, to, err)
	}

	if swapResult.Err != nil {
		return nil, fmt.Errorf("failed to swap %s %s tokens to %s tokens: %s", amountIn, from, to, *swapResult.Err)
	}

	return swapResult.Ok.ReceiveAmount.BigInt(), nil
}

func (k *kong) swapIcrc2(from, to icp.Token, amountIn, amountOutMin *big.Int) (*big.Int, error) {
	receiveAmount := idl.NewBigNat(amountOutMin)
	fromFees := from.Metadata().Fee
	maxSlippage := 100.0

	swapResult, err := k.api.Swap(SwapArgs{
		PayToken:      fmt.Sprintf("IC.%s", from.CanisterID()),
		PayAmount:     idl.NewBigNat(new(big.Int).Sub(amountIn, fromFees)),
		ReceiveToken:  fmt.Sprintf("IC.%s", to.CanisterID()),
		ReceiveAmount: &receiveAmount,
		MaxSlippage:   &maxSlippage,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to swap %s %s tokens to %s tokens: %w", amountIn, from, to, err)
	}

	if swapResult.Err != nil {
		return nil, fmt.Errorf("failed to swap %s %s tokens to %s tokens: %s", amountIn, from, to, *swapResult.Err)
	}

	return swapResult.Ok.ReceiveAmount.BigInt(), nil
}
