package treap

// Treap interface
type Treap interface {
	Add(...int)
	Find(x int) bool
}

// Config generic configs for all Treaps
type Config struct {
	AllowDuplicates bool //allow only unique keys in the treap
}

// DefaultConfig configuration for most of the cases
var DefaultConfig = Config{AllowDuplicates: true}
