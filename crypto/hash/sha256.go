package sha256

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hash(input interface{}) []byte {
	h := sha256.New()

	switch v := input.(type) {
	case string:
		h.Write([]byte(v))

	case []byte:
		h.Write(v)

	default:
		return []byte{}
	}

	hash := h.Sum(nil)
	hexHash := make([]byte, hex.EncodedLen(32))
	hex.Encode(hexHash, hash)
	return hexHash

}
