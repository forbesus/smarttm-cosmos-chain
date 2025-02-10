package types

const (
	// ModuleName defines the module name
	ModuleName = "smarttm"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_smarttm"
)

var (
	ParamsKey = []byte("p_smarttm")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
