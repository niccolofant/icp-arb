package icp

import "math/big"

type DexType string

const (
	DexTypeIcpSwap    DexType = "icpswap"
	DexTypeIcpExV2    DexType = "icpex-v2"
	DexTypeIcpExV1    DexType = "icpex-v1"
	DexTypeIcLight    DexType = "iclight"
	DexTypeNeutrinite DexType = "neutrinite"
	DexTypeKongSwap   DexType = "kongswap"
	DexTypeSonic      DexType = "sonic"
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
}

type DexSwap interface {
	Swap(from, to Token, amountIn, amountOutMin *big.Int) (*big.Int, error)
}
