package icp

type TokenStandard string

const (
	TokenStandardICP   TokenStandard = "ICP"
	TokenStandardICRC1 TokenStandard = "ICRC1"
	TokenStandardICRC2 TokenStandard = "ICRC2"
)

func (ts TokenStandard) String() string {
	return string(ts)
}
