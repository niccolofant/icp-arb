package route

import (
	"errors"
	"math/big"

	icarb "github.com/niccolofant/ic-arb"
	"github.com/niccolofant/ic-arb/core/icp"
)

// todo: use one step swap
func (s *Step) Swap(amountIn, amountOutMin *big.Int) (*big.Int, error) {
	swapper, ok := s.dex.(icp.DexSwap)
	if !ok {
		return icarb.Zero, errors.New("step dex does not implement DexSwap interface")
	}

	return swapper.Swap(icp.DexSwapParams{
		FromToken:    s.FromToken(),
		ToToken:      s.ToToken(),
		AmountIn:     amountIn,
		AmountOutMin: amountOutMin,
	})
}
