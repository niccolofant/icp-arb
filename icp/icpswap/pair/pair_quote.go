package icpswap_pair

import (
	"fmt"
	"math/big"

	icarb "github.com/niccolofant/ic-arb"
	"github.com/niccolofant/ic-arb/icp"
)

func (p *pair) Quote(from, to icp.Token, amountIn *big.Int) (*big.Int, error) {
	zeroForOne := true
	if from.Equal(p.Token1()) {
		zeroForOne = false
	}

	fromMetadata := from.Metadata()
	fromFees := fromMetadata.Fee
	toFees := to.Metadata().Fee

	amountInWithFees := new(big.Int).Sub(amountIn, fromFees)
	if fromMetadata.Standard == icp.TokenStandardICRC1 {
		amountInWithFees = new(big.Int).Sub(amountInWithFees, fromFees)
	}

	quoteResult, err := p.api.Quote(SwapArgs{
		AmountIn:         amountInWithFees.String(),
		AmountOutMinimum: icarb.Zero.String(),
		ZeroForOne:       zeroForOne,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to quote swap %s %s tokens to %s tokens: %w", amountIn, from, to, err)
	}

	if quoteResult.Err != nil {
		return nil, fmt.Errorf("failed to quote swap %s %s tokens to %s tokens: %w", amountIn, from, to, quoteResult.Err.Decode())
	}

	return new(big.Int).Sub(quoteResult.Ok.BigInt(), toFees), nil
}
