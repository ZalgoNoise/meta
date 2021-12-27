package hash

// Hasher interface defines the behavior of a recursive hasher
// which takes in an empty interface (to support multiple formats)
// and returns a slice of bytes (the hashed seed)
type Hasher interface {
	Hash(interface{}) []byte
}
