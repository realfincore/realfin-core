package types

const (
	// ModuleName defines the module name
	ModuleName = "realfin"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_realfin"
)

var (
	ParamsKey = []byte("p_realfin")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
