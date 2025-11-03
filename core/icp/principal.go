package icp

import (
	"github.com/aviate-labs/agent-go/principal"
)

var LedgerPrincipal = MustDecodePrincipal("ryjl3-tyaaa-aaaaa-aaaba-cai")

type Principal struct {
	raw principal.Principal
}

func NewPrincipal(raw principal.Principal) Principal {
	return Principal{
		raw: raw,
	}
}

func MustDecodePrincipal(s string) Principal {
	return Principal{
		raw: principal.MustDecode(s),
	}
}

func (p Principal) Raw() principal.Principal {
	return p.raw
}

func (p Principal) String() string {
	return p.raw.String()
}

func (p Principal) Equal(other Principal) bool {
	return p.raw.Equal(other.Raw())
}

func (p Principal) Blob() []byte {
	principalBytes := p.Raw().Raw
	result := make([]byte, 32)
	result[0] = byte(len(principalBytes))
	copyLen := min(len(principalBytes), 31)
	copy(result[1:], principalBytes[:copyLen])

	return result
}
