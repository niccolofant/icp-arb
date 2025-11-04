package icp

import "math/big"

type TokenStandard string

const (
	TokenStandardICP   TokenStandard = "ICP"
	TokenStandardICRC1 TokenStandard = "ICRC-1"
	TokenStandardICRC2 TokenStandard = "ICRC-2"
)

func (ts TokenStandard) String() string {
	return string(ts)
}

type Token interface {
	Canister
	Metadata() TokenMetadata
	IsICP() bool
}

type TokenMetadata struct {
	Name     string
	Symbol   string
	Fee      *big.Int
	Standard TokenStandard
	Decimals int
}
