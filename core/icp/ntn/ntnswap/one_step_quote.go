package ntnswap

import (
	"fmt"
	"math/big"

	"github.com/niccolofant/ic-arb/core/icp"
)

func (n *ntnswap) OneStepQuote(
	params icp.DexQuoteParams,
	opts *icp.DexOneStepSwapOpts,
) (*big.Int, error) {
	if opts == nil {
		opts = &icp.DexOneStepSwapOpts{}
	}

	amountToQuote := new(big.Int).Set(params.AmountIn)

	if !opts.SkipDeposit {
		fee := new(big.Int).Mul(big.NewInt(2), params.FromToken.Metadata().Fee)
		if amountToQuote.Cmp(fee) <= 0 {
			return nil, fmt.Errorf("amount too small: less than or equal to deposit fee")
		}
		amountToQuote.Sub(amountToQuote, fee)
	}

	quoteResult, err := n.Quote(params)
	if err != nil {
		return nil, err
	}

	withdrawAmount := new(big.Int).Sub(quoteResult, params.ToToken.Metadata().Fee)

	if !opts.SkipWithdraw {
		if withdrawAmount.Cmp(params.ToToken.Metadata().Fee) <= 0 {
			return nil, fmt.Errorf("withdraw amount too small: less than or equal to withdraw fee")
		}
		withdrawAmount.Sub(withdrawAmount, params.ToToken.Metadata().Fee)
	}

	return withdrawAmount, nil
}
