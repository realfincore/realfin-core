package types

const (
	// ModuleName defines the module name
	ModuleName = "creditscore"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_creditscore"
)

var (
	ParamsKey = []byte("p_creditscore")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
