package types

const (
	// ModuleName defines the module name
	ModuleName = "blog"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_blog"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	HelpKey      = "Help/value/"
	HelpCountKey = "Help/count/"
)

const (
	PostKey      = "Post/value/"
	PostCountKey = "Post/count/"
)
