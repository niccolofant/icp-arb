package kong

import (
	"fmt"
	"math/big"

	"github.com/aviate-labs/agent-go/candid/idl"
	"github.com/niccolofant/ic-arb/core/icp"
	"github.com/niccolofant/ic-arb/core/icp/icrc1"
)

func (k *kong) OneStepSwap(
	params icp.DexSwapParams,
	opts *icp.DexOneStepSwapOpts, // unused
) (*big.Int, error) {
	if params.FromToken.Metadata().Standard == icp.TokenStandardICRC1 {
		return k.swapIcrc1(params)
	}

	return k.swapIcrc2(params)
}

func (k *kong) swapIcrc1(params icp.DexSwapParams) (*big.Int, error) {
	fromIcrc1, isIcrc1 := params.FromToken.(icrc1.ICRC1)
	if !isIcrc1 {
		return nil, fmt.Errorf("from token does not implement icrc1 interface")
	}

	transferResult, err := fromIcrc1.Transfer(params.AmountIn, k.CanisterID(), nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to transfer %s %s tokens to kong: %w", params.AmountIn, fromIcrc1, err)
	}

	blockIdx := idl.NewBigNat(transferResult.BlockIdx)
	receiveAmount := idl.NewBigNat(params.AmountOutMin)
	fromFees := params.FromToken.Metadata().Fee
	maxSlippage := 100.0

	swapResult, err := k.api.Swap(SwapArgs{
		PayToken:      fmt.Sprintf("IC.%s", params.FromToken.CanisterID()),
		PayAmount:     idl.NewBigNat(new(big.Int).Sub(params.AmountIn, fromFees)),
		ReceiveToken:  fmt.Sprintf("IC.%s", params.ToToken.CanisterID()),
		ReceiveAmount: &receiveAmount,
		MaxSlippage:   &maxSlippage,
		PayTxId: &TxId{
			BlockIndex: &blockIdx,
		},
	})
	if err != nil {
		return nil, fmt.Errorf(
			"failed to swap %s %s tokens to %s tokens: %w",
			params.AmountIn,
			params.FromToken,
			params.ToToken,
			err,
		)
	}

	if swapResult.Err != nil {
		return nil, fmt.Errorf(
			"failed to swap %s %s tokens to %s tokens: %w",
			params.AmountIn,
			params.FromToken,
			params.ToToken,
			*swapResult.Err,
		)
	}

	return swapResult.Ok.ReceiveAmount.BigInt(), nil
}

func (k *kong) swapIcrc2(params icp.DexSwapParams) (*big.Int, error) {
	receiveAmount := idl.NewBigNat(params.AmountOutMin)
	fromFees := params.FromToken.Metadata().Fee
	maxSlippage := 100.0

	swapResult, err := k.api.Swap(SwapArgs{
		PayToken:      fmt.Sprintf("IC.%s", params.FromToken.CanisterID()),
		PayAmount:     idl.NewBigNat(new(big.Int).Sub(params.AmountIn, fromFees)),
		ReceiveToken:  fmt.Sprintf("IC.%s", params.ToToken.CanisterID()),
		ReceiveAmount: &receiveAmount,
		MaxSlippage:   &maxSlippage,
	})
	if err != nil {
		return nil, fmt.Errorf(
			"failed to swap %s %s tokens to %s tokens: %w",
			params.AmountIn,
			params.FromToken,
			params.ToToken,
			err,
		)
	}

	if swapResult.Err != nil {
		return nil, fmt.Errorf(
			"failed to swap %s %s tokens to %s tokens: %w",
			params.AmountIn,
			params.FromToken,
			params.ToToken,
			*swapResult.Err,
		)
	}

	return swapResult.Ok.ReceiveAmount.BigInt(), nil
}
