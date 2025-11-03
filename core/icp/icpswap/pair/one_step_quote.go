package pair

import (
	"fmt"
	"math/big"

	icarb "github.com/niccolofant/ic-arb"
	"github.com/niccolofant/ic-arb/core/icp"
)

func (p *pair) OneStepQuote(
	params icp.DexQuoteParams,
	opts *icp.DexOneStepSwapOpts, // unused
) (*big.Int, error) {
	zeroForOne := true
	if params.FromToken.Equal(p.Token1()) {
		zeroForOne = false
	}

	fromMetadata := params.FromToken.Metadata()
	fromFees := fromMetadata.Fee
	toFees := params.ToToken.Metadata().Fee

	amountInWithFees := new(big.Int).Sub(params.AmountIn, fromFees)
	if fromMetadata.Standard == icp.TokenStandardICRC1 {
		amountInWithFees.Sub(amountInWithFees, fromFees)
	}

	quoteResult, err := p.api.Quote(SwapArgs{
		AmountIn:         amountInWithFees.String(),
		AmountOutMinimum: icarb.Zero.String(),
		ZeroForOne:       zeroForOne,
	})
	if err != nil {
		return nil, fmt.Errorf(
			"failed to quote swap %s %s tokens to %s tokens: %w",
			params.AmountIn,
			params.FromToken,
			params.ToToken,
			err,
		)
	}

	if quoteResult.Err != nil {
		return nil, fmt.Errorf(
			"failed to quote swap %s %s tokens to %s tokens: %w",
			params.AmountIn,
			params.FromToken,
			params.ToToken,
			quoteResult.Err.Decode(),
		)
	}

	return new(big.Int).Sub(quoteResult.Ok.BigInt(), toFees), nil
}
