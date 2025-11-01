package icp

type Canister interface {
	CanisterID() Principal
	Equal(other Canister) bool
}
