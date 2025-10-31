package icp

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
