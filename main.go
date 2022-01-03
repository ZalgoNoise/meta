package main

import (
	"fmt"

	"github.com/ZalgoNoise/meta/crypto/hash"
)

func main() {
	// charmatcher example
	//
	//
	// matcher, err := cm.NewCharMatcher(`abcdefg`)

	// if err != nil {
	// 	panic(err)
	// }

	// if matcher.Contains(`abc`) {
	// 	fmt.Println(`yep!`)
	// } else {
	// 	fmt.Println(`nope!`)
	// }

	// recursive hashers example
	//
	seed := "hashfunction"
	var h hash.Hasher

	// md5
	header("md5")
	h = hash.MD5{}
	fmt.Println(recHash(seed, h))

	// sha1
	header("sha1")
	h = hash.SHA1{}
	fmt.Println(recHash(seed, h))

	// sha224
	header("sha224")
	h = hash.SHA224{}
	fmt.Println(recHash(seed, h))

	// sha256
	header("sha256")
	h = hash.SHA256{}
	fmt.Println(recHash(seed, h))

	// sha384
	header("sha384")
	h = hash.SHA384{}
	fmt.Println(recHash(seed, h))

	// sha512
	header("sha512")
	h = hash.SHA512{}
	fmt.Println(recHash(seed, h))

	// sha512_224
	header("sha512_224")
	h = hash.SHA512_224{}
	fmt.Println(recHash(seed, h))

	// sha512_256
	header("sha512_256")
	h = hash.SHA512_256{}
	fmt.Println(recHash(seed, h))

}

func header(hf string) {
	fmt.Printf("#\tHashing %s\t#\n\nIndex   |  Hash\n----------------\n", hf)
}

func recHash(seed string, h hash.Hasher) string {
	counter := 0
	out := fmt.Sprintf("#%v\t| %s\n", counter, seed)
	counter++
	hash := h.Hash([]byte(seed))
	out += fmt.Sprintf("#%v\t| %s\n", counter, string(hash))

	for counter = 2; counter <= 10; counter++ {
		hash = h.Hash(hash)
		out += fmt.Sprintf("#%v\t| %s\n", counter, string(hash))
	}
	return out
}
