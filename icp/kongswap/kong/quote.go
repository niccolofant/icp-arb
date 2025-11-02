package kong

import (
	"fmt"
	"math/big"

	"github.com/aviate-labs/agent-go/candid/idl"
	"github.com/niccolofant/ic-arb/icp"
)

func (k *kong) Quote(from, to icp.Token, amountIn *big.Int) (*big.Int, error) {
	fromMetadata := from.Metadata()
	fromFees := fromMetadata.Fee

	amountInWithFees := new(big.Int).Sub(amountIn, fromFees)
	// if fromMetadata.Standard == icp.TokenStandardICRC1 {
	// 	amountInWithFees = new(big.Int).Sub(amountInWithFees, fromFees)
	// }

	quoteResult, err := k.api.SwapAmounts(
		fmt.Sprintf("IC.%s", from.CanisterID()),
		idl.NewBigNat(amountInWithFees),
		fmt.Sprintf("IC.%s", to.CanisterID()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to quote swap %s %s tokens to %s tokens: %w", amountIn, from, to, err)
	}

	if quoteResult.Err != nil {
		return nil, fmt.Errorf("failed to quote swap %s %s tokens to %s tokens: %s", amountIn, from, to, *quoteResult.Err)
	}

	return quoteResult.Ok.ReceiveAmount.BigInt(), nil
}
