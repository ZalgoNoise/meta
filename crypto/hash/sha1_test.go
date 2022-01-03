package hash

import (
	"testing"
)

func TestSHA1Hash(t *testing.T) {
	type input struct {
		seed       string
		iterations int
	}
	type test struct {
		input input
		ok    string
	}

	var expectedResults []string = []string{
		"2ef7bde608ce5404e97d5f042f95f89f1c232871",
		"1d82b8fb700e2ca794ac0c94c6b0d8a709b2c849",
		"09d09cafbec7adffc61981a1e8f88e8b4954d77c",
		"f2140cd51196c4063afa830a168d69edb892f511",
		"8fd962dfe55c76668620cfd5653acf802a294716",
		"9e039636e3d7becaa1fc45d0d2dd6271988cdeb7",
		"0d18c61269043a8f6d501738a4e8780ec7a469cd",
		"e9f585d24b0fa1538ef86cdce2d467639e7160ea",
		"78966f253b665176692e6ca8247106234ae71916",
		"259db186b1fe877ea9a435c256ead03f56473a9d",
		"6a59e3381dc721f0bd333c55fab53cc339c626f0",
		"0eaa0387e95e74cfaf4d5668ad49a5d5e6d8c5f5",
		"21ce424bea712c8ca6affcf347f88c546d021b48",
		"64317d5c81c62447f1e21a9a3b0feef814a7e3ed",
		"49316e888ec6d199d6f019de55746ee9ed1cb610",
		"bf27527cb44b65ab82f57fa1594ede56b42eb374",
		"82d4dadd1b2b13fb5e4219d04563a6974e0eff16",
		"8ca51ce0576718a7f3c29e8207bbc0df973870ab",
		"f410bc514f2a54369922920c8fcbcaaa0e8bff3f",
		"2d862ec148245f8761d6884481bcf5bd56e30ef5",
	}

	tests := []test{}

	counter := 0
	for s := 0; s < len(MockStrings); s++ {
		for i := 1; i <= MockIter; i++ {
			tests = append(tests, test{
				input: input{
					seed:       MockStrings[s],
					iterations: i,
				},
				ok: expectedResults[counter],
			})
			counter++
		}
	}

	h := SHA1{}

	for _, test := range tests {
		var hash []byte

		for i := 1; i <= test.input.iterations; i++ {
			if i == 1 {
				hash = h.Hash([]byte(test.input.seed))
			} else {
				hash = h.Hash(hash)
			}
		}

		result := string(hash)

		if test.ok != result {
			t.Errorf(`[sha1] Hash([]byte(%s)) x %v = %s ; expected %s`,
				test.input.seed,
				test.input.iterations,
				result,
				test.ok,
			)
		}
	}
}
