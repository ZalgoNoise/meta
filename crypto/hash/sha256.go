package hash

import (
	"crypto/sha256"
	"encoding/hex"
)

// SHA256 struct is a general placeholder for the SHA256
// algorithm, to implement the Hasher interface
type SHA256 struct{}

// Hash method satisfies the Hasher interface. It's a
// recursive hashing function to allow continuous hashing
func (hasher SHA256) Hash(input interface{}) []byte {

	var hash []byte = make([]byte, hex.EncodedLen(32))
	var sum [32]byte

	switch v := input.(type) {
	case string:
		sum = sha256.Sum256([]byte(v))

	case []byte:
		sum = sha256.Sum256(v)

	default:
		return []byte{}
	}

	hex.Encode(hash, sum[:])

	return hash

}
