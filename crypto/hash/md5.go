package hash

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 struct is a general placeholder for the MD5
// algorithm, to implement the Hasher interface
type MD5 struct{}

// Hash method satisfies the Hasher interface. It's a
// recursive hashing function to allow continuous hashing
func (hasher MD5) Hash(data []byte) []byte {

	var hash []byte = make([]byte, hex.EncodedLen(16))
	var sum [16]byte = md5.Sum(data)

	hex.Encode(hash, sum[:])

	return hash

}
