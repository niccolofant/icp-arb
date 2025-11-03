package pair

import (
	"fmt"
	"math/big"

	"github.com/aviate-labs/agent-go/candid/idl"
	"github.com/niccolofant/ic-arb/core/icp"
	"github.com/niccolofant/ic-arb/core/icp/icrc1"
)

func (p *pair) OneStepSwap(
	params icp.DexSwapParams,
	opts *icp.DexOneStepSwapOpts, // unused
) (*big.Int, error) {
	if params.FromToken.Metadata().Standard == icp.TokenStandardICRC1 {
		return p.depositAndSwap(params)
	}

	return p.depositFromAndSwap(params)
}

func (p *pair) depositAndSwap(params icp.DexSwapParams) (*big.Int, error) {
	fromIcrc1, isIcrc1 := params.FromToken.(icrc1.ICRC1)
	if !isIcrc1 {
		return nil, fmt.Errorf("from token does not implement icrc1 interface")
	}

	zeroForOne := true
	if fromIcrc1.Equal(p.Token1()) {
		zeroForOne = false
	}

	fromFees := fromIcrc1.Metadata().Fee
	toFees := params.ToToken.Metadata().Fee
	subaccount := p.api.Agent().Sender().Blob()

	transferResult, err := fromIcrc1.Transfer(
		params.AmountIn,
		p.CanisterID(),
		&subaccount,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to transfer %s %s tokens to subaccount %s: %w",
			params.AmountIn,
			fromIcrc1,
			subaccount,
			err,
		)
	}

	amountToSwap := new(big.Int).Sub(transferResult.Amount, fromFees)

	swapResult, err := p.api.DepositAndSwap(DepositAndSwapArgs{
		AmountIn:         amountToSwap.String(),
		AmountOutMinimum: params.AmountOutMin.String(),
		TokenInFee:       idl.NewBigNat(fromIcrc1.Metadata().Fee),
		TokenOutFee:      idl.NewBigNat(toFees),
		ZeroForOne:       zeroForOne,
	})
	if err != nil {
		return nil, fmt.Errorf(
			"failed to swap %s %s tokens to %s tokens: %w",
			params.AmountIn,
			fromIcrc1,
			params.ToToken,
			err,
		)
	}

	if swapResult.Err != nil {
		return nil, fmt.Errorf(
			"failed to swap %s %s tokens to %s tokens: %w",
			params.AmountIn,
			fromIcrc1,
			params.ToToken,
			swapResult.Err.Decode(),
		)

	}

	return new(big.Int).Sub(swapResult.Ok.BigInt(), toFees), nil
}

func (p *pair) depositFromAndSwap(params icp.DexSwapParams) (*big.Int, error) {
	zeroForOne := true
	if params.FromToken.Equal(p.Token1()) {
		zeroForOne = false
	}

	fromFees := params.FromToken.Metadata().Fee
	toFees := params.ToToken.Metadata().Fee

	swapResult, err := p.api.DepositFromAndSwap(DepositAndSwapArgs{
		AmountIn:         new(big.Int).Sub(params.AmountIn, fromFees).String(),
		AmountOutMinimum: params.AmountOutMin.String(),
		TokenInFee:       idl.NewBigNat(fromFees),
		TokenOutFee:      idl.NewBigNat(toFees),
		ZeroForOne:       zeroForOne,
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
			swapResult.Err.Decode(),
		)

	}

	return new(big.Int).Sub(swapResult.Ok.BigInt(), toFees), nil
}
