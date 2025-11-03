package ntnswap

import (
	"math/big"

	"github.com/niccolofant/ic-arb/core/icp"
)

type Tx struct {
	Token          icp.Token
	Memo           []byte
	FromSubaccount *[]byte
	To             icp.Principal
	Amount         *big.Int
}
