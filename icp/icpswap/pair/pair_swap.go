package icpswap_pair

import (
	"fmt"
	"math/big"

	"github.com/aviate-labs/agent-go/candid/idl"
	"github.com/niccolofant/ic-arb/icp"
	"github.com/niccolofant/ic-arb/icp/icrc1"
)

func (p *pair) Swap(from, to icp.Token, amountIn, amountOutMin *big.Int) (*big.Int, error) {
	if from.Metadata().Standard == icp.TokenStandardICRC1 {
		return p.depositAndSwap(from, to, amountIn, amountOutMin)
	}

	return p.depositFromAndSwap(from, to, amountIn, amountOutMin)
}

func (p *pair) depositAndSwap(from, to icp.Token, amountIn, amountOutMin *big.Int) (*big.Int, error) {
	fromIcrc1, isIcrc1 := from.(icrc1.ICRC1)
	if !isIcrc1 {
		return nil, fmt.Errorf("from token does not implement icrc1 interface")
	}

	zeroForOne := true
	if fromIcrc1.Equal(p.Token1()) {
		zeroForOne = false
	}

	fromFees := fromIcrc1.Metadata().Fee
	toFees := to.Metadata().Fee
	subaccount := p.api.Agent().Sender().Blob()

	transferredAmount, err := fromIcrc1.Transfer(
		amountIn,
		p.CanisterID(),
		&subaccount,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to transfer %s %s tokens to subaccount %s: %w", amountIn, fromIcrc1, subaccount, err)
	}

	amountToSwap := new(big.Int).Sub(transferredAmount, fromFees)

	swapResult, err := p.api.DepositAndSwap(DepositAndSwapArgs{
		AmountIn:         amountToSwap.String(),
		AmountOutMinimum: amountOutMin.String(),
		TokenInFee:       idl.NewBigNat(fromIcrc1.Metadata().Fee),
		TokenOutFee:      idl.NewBigNat(toFees),
		ZeroForOne:       zeroForOne,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to swap %s %s tokens to %s tokens: %w", amountIn, fromIcrc1, to, err)
	}

	if swapResult.Err != nil {
		return nil, fmt.Errorf("failed to swap %s %s tokens to %s tokens: %w", amountIn, fromIcrc1, to, swapResult.Err.Decode())

	}

	return new(big.Int).Sub(swapResult.Ok.BigInt(), toFees), nil
}

func (p *pair) depositFromAndSwap(from, to icp.Token, amountIn, amountOutMin *big.Int) (*big.Int, error) {
	zeroForOne := true
	if from.Equal(p.Token1()) {
		zeroForOne = false
	}

	fromFees := from.Metadata().Fee
	toFees := to.Metadata().Fee

	swapResult, err := p.api.DepositFromAndSwap(DepositAndSwapArgs{
		AmountIn:         new(big.Int).Sub(amountIn, fromFees).String(),
		AmountOutMinimum: amountOutMin.String(),
		TokenInFee:       idl.NewBigNat(from.Metadata().Fee),
		TokenOutFee:      idl.NewBigNat(toFees),
		ZeroForOne:       zeroForOne,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to swap %s %s tokens to %s tokens: %w", amountIn, from, to, err)
	}

	if swapResult.Err != nil {
		return nil, fmt.Errorf("failed to swap %s %s tokens to %s tokens: %w", amountIn, from, to, swapResult.Err.Decode())

	}

	return new(big.Int).Sub(swapResult.Ok.BigInt(), toFees), nil
}
