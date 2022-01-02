package hash

import (
	"crypto/sha1"
	"encoding/hex"
)

// SHA1 struct is a general placeholder for the SHA1
// algorithm, to implement the Hasher interface
type SHA1 struct{}

// Hash method satisfies the Hasher interface. It's a
// recursive hashing function to allow continuous hashing
func (hasher SHA1) Hash(data []byte) []byte {

	var hash []byte = make([]byte, hex.EncodedLen(20))
	var sum [20]byte = sha1.Sum(data)

	hex.Encode(hash, sum[:])

	return hash

}
