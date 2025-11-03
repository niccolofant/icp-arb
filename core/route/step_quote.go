package route

import (
	"fmt"
	"math/big"

	icarb "github.com/niccolofant/ic-arb"
	"github.com/niccolofant/ic-arb/core/icp"
)

// todo: use one step swap
func (s *Step) Quote(amountIn *big.Int) (*big.Int, error) {
	quoter, ok := s.Dex().(icp.DexQuote)
	if !ok {
		return icarb.Zero, fmt.Errorf("step dex does not implement DexQuote interface")
	}

	return quoter.Quote(icp.DexQuoteParams{
		FromToken: s.FromToken(),
		ToToken:   s.ToToken(),
		AmountIn:  amountIn,
	})
}
