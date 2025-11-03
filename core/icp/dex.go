package icp

import (
	"math/big"
)

type DexType string

const (
	DexTypeIcpSwap  DexType = "icpswap"
	DexTypeIcpExV2  DexType = "icpex-v2"
	DexTypeIcpExV1  DexType = "icpex-v1"
	DexTypeIcLight  DexType = "iclight"
	DexTypeNtn      DexType = "ntn"
	DexTypeKongSwap DexType = "kongswap"
	DexTypeSonic    DexType = "sonic"
)

func (dt DexType) String() string {
	return string(dt)
}

type Dex interface {
	Canister
	Type() DexType
}

type DexNotAggregated interface {
	Token1() Token
	Token0() Token
	SupportToken(token Token) bool
}

type DexSwap interface {
	Swap(params DexSwapParams) (*big.Int, error)
}

type DexOneStepSwap interface {
	OneStepSwap(params DexSwapParams, opts *DexOneStepSwapOpts) (*big.Int, error)
}

type DexSwapParams struct {
	FromToken    Token
	ToToken      Token
	AmountIn     *big.Int
	AmountOutMin *big.Int
}

type DexQuoteParams struct {
	FromToken Token
	ToToken   Token
	AmountIn  *big.Int
}

type DexOneStepSwapOpts struct {
	SkipDeposit  bool
	SkipWithdraw bool
}

type DexQuote interface {
	Quote(params DexQuoteParams) (*big.Int, error)
}

type DexOneStepQuote interface {
	OneStepQuote(params DexQuoteParams, opts *DexOneStepSwapOpts) (*big.Int, error)
}

// type DexDeposit interface {
// 	Deposit(token Token, amount *big.Int) (*big.Int, error)
// }

// type DexWithdraw interface {
// 	Withdraw(token Token, amount *big.Int) (*big.Int, error)
// }
