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
			return h.Sum(nil)
			
		case []byte:
			oldHash := make([]byte, hex.EncodedLen(32))
			hex.Encode(oldHash, v)
			h.Write(oldHash)

			return h.Sum(nil)

		default:
			return []byte{}
	}
	
}