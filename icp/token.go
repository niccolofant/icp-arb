package icp

import "math/big"

type TokenStandard string

const (
	TokenStandardICP   TokenStandard = "ICP"
	TokenStandardICRC1 TokenStandard = "ICRC1"
	TokenStandardICRC2 TokenStandard = "ICRC2"
)

func (ts TokenStandard) String() string {
	return string(ts)
}

type Token interface {
	Canister
	Metadata() Metadata
}

type Metadata struct {
	Name   string
	Symbol string
	Fee    *big.Int
}
