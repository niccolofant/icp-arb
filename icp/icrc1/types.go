package icrc1

import "math/big"

type TransferResponse struct {
	Amount   *big.Int
	BlockIdx *big.Int
}
