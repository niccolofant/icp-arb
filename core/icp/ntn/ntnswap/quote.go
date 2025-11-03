package ntnswap

import (
	"fmt"
	"math/big"

	"github.com/aviate-labs/agent-go/candid/idl"
	"github.com/niccolofant/ic-arb/core/icp"
)

func (n *ntnswap) Quote(params icp.DexQuoteParams) (*big.Int, error) {
	fromTokenPrincipalRaw := params.FromToken.CanisterID().Raw()
	toTokenPrincipalRaw := params.ToToken.CanisterID().Raw()

	quoteResult, err := n.api.DexQuote(QuoteRequest{
		Amount: idl.NewBigNat(params.AmountIn),
		LedgerFrom: SupportedLedger{
			Ic: &fromTokenPrincipalRaw,
		},
		LedgerTo: SupportedLedger{
			Ic: &toTokenPrincipalRaw,
		},
	})

	if err != nil {
		return nil, fmt.Errorf(
			"failed to quote %s %s tokens to %s tokens: %w",
			params.AmountIn,
			params.FromToken,
			params.ToToken,
			err,
		)
	}

	if quoteResult.Err != nil {
		return nil, fmt.Errorf(
			"failed to quote %s %s tokens to %s tokens: %s",
			params.AmountIn,
			params.FromToken,
			params.ToToken,
			*quoteResult.Err,
		)
	}

	return new(big.Int).Sub(quoteResult.Ok.AmountOut.BigInt(), params.ToToken.Metadata().Fee), nil
}
