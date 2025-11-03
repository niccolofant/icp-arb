package kong

import (
	"fmt"
	"math/big"

	"github.com/aviate-labs/agent-go/candid/idl"
	"github.com/niccolofant/ic-arb/core/icp"
)

func (k *kong) OneStepQuote(
	params icp.DexQuoteParams,
	opts *icp.DexOneStepSwapOpts, // unused
) (*big.Int, error) {
	amountInWithFees := new(big.Int).Sub(params.AmountIn, params.FromToken.Metadata().Fee)

	quoteResult, err := k.api.SwapAmounts(
		fmt.Sprintf("IC.%s", params.FromToken.CanisterID()),
		idl.NewBigNat(amountInWithFees),
		fmt.Sprintf("IC.%s", params.ToToken.CanisterID()),
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to quote swap %s %s tokens to %s tokens: %w",
			params.AmountIn,
			params.FromToken,
			params.ToToken,
			err)
	}

	if quoteResult.Err != nil {
		return nil, fmt.Errorf(
			"failed to quote swap %s %s tokens to %s tokens: %s",
			params.AmountIn,
			params.FromToken,
			params.ToToken,
			*quoteResult.Err,
		)
	}

	return quoteResult.Ok.ReceiveAmount.BigInt(), nil
}
