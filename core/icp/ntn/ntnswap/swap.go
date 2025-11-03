package ntnswap

import (
	"fmt"
	"math/big"

	"github.com/aviate-labs/agent-go/candid/idl"
	"github.com/niccolofant/ic-arb/core/icp"
)

func (n *ntnswap) Swap(params icp.DexSwapParams) (*big.Int, error) {
	fromTokenPrincipalRaw := params.FromToken.CanisterID().Raw()
	toTokenPrincipalRaw := params.ToToken.CanisterID().Raw()
	botPrincipalRaw := n.api.Agent().Sender().Raw()

	// SWAP IS INSTANT BUT WE DONT FUCKING KNOW THE AMOUNT OUT, ATM WE DO SOME DOGSHIT
	// STUFF BY QUOTING BEFORE THE SWAP
	quoteResult, err := n.Quote(icp.DexQuoteParams{
		FromToken: params.FromToken,
		ToToken:   params.ToToken,
		AmountIn:  params.AmountIn,
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

	swapResult, err := n.api.DexSwap(SwapRequest{
		Account: Account{
			Owner: botPrincipalRaw,
		},
		Amount:       idl.NewBigNat(params.AmountIn),
		MinAmountOut: idl.NewBigNat(quoteResult),
		LedgerFrom: SupportedLedger{
			Ic: &fromTokenPrincipalRaw,
		},
		LedgerTo: SupportedLedger{
			Ic: &toTokenPrincipalRaw,
		},
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
			"failed to swap %s %s tokens to %s tokens: %s",
			params.AmountIn,
			params.FromToken,
			params.ToToken,
			*swapResult.Err,
		)
	}

	return quoteResult, nil
}
