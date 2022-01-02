package hash

import (
	"crypto/sha512"
	"encoding/hex"
)

// SHA512 struct is a general placeholder for the SHA512
// algorithm, to implement the Hasher interface
type SHA512 struct{}

// Hash method satisfies the Hasher interface. It's a
// recursive hashing function to allow continuous hashing
func (hasher SHA512) Hash(data []byte) []byte {

	var hash []byte = make([]byte, hex.EncodedLen(64))
	var sum [64]byte = sha512.Sum512(data)

	hex.Encode(hash, sum[:])

	return hash

}

// SHA384 struct is a general placeholder for the SHA384
// algorithm, to implement the Hasher interface
type SHA384 struct{}

// Hash method satisfies the Hasher interface. It's a
// recursive hashing function to allow continuous hashing
func (hasher SHA384) Hash(data []byte) []byte {

	var hash []byte = make([]byte, hex.EncodedLen(48))
	var sum [48]byte = sha512.Sum384(data)

	hex.Encode(hash, sum[:])

	return hash

}

// SHA512 struct is a general placeholder for the SHA512
// algorithm, to implement the Hasher interface
type SHA512_256 struct{}

// Hash method satisfies the Hasher interface. It's a
// recursive hashing function to allow continuous hashing
func (hasher SHA512_256) Hash(data []byte) []byte {

	var hash []byte = make([]byte, hex.EncodedLen(32))
	var sum [32]byte = sha512.Sum512_256(data)

	hex.Encode(hash, sum[:])

	return hash

}

// SHA512 struct is a general placeholder for the SHA512
// algorithm, to implement the Hasher interface
type SHA512_224 struct{}

// Hash method satisfies the Hasher interface. It's a
// recursive hashing function to allow continuous hashing
func (hasher SHA512_224) Hash(data []byte) []byte {

	var hash []byte = make([]byte, hex.EncodedLen(28))
	var sum [28]byte = sha512.Sum512_224(data)

	hex.Encode(hash, sum[:])

	return hash

}
