package ntnswap

import (
	"math/big"

	"github.com/niccolofant/ic-arb/core/icp"
)

func (n *ntnswap) OneStepSwap(
	params icp.DexSwapParams,
	opts *icp.DexOneStepSwapOpts,
) (*big.Int, error) {
	if opts == nil {
		opts = &icp.DexOneStepSwapOpts{}
	}

	amountToSwap := new(big.Int).Set(params.AmountIn)

	if !opts.SkipDeposit {
		depositedAmount, err := n.Deposit(params.FromToken, params.AmountIn)
		if err != nil {
			return nil, err
		}
		amountToSwap.Set(depositedAmount)
	}

	swappedAmount, err := n.Swap(params)
	if err != nil {
		return nil, err
	}

	if opts.SkipWithdraw {
		return swappedAmount, nil
	}

	withdrawnAmount, err := n.Withdraw(params.ToToken, swappedAmount)
	if err != nil {
		return nil, err
	}

	return withdrawnAmount, nil
}
