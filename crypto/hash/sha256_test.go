package hash

import (
	"testing"
)

var inputStrings []string = []string{
	"Hello World!",
	"Hello, 世界!",
}

var inputIterations []int = []int{
	1,
	2,
	10,
}

var expectedResultsSHA256 []string = []string{
	"7f83b1657ff1fc53b92dc18148a1d65dfc2d4b1fa3d677284addd200126d9069",
	"4163fb4ab9e1e0a51709a51bc7e13ab6792907905960145c722d2c1479caac42",
	"1fada6a9b8084eff6347baeae0812aac1c14fcc0a97759a6bfa2d8a2ea087705",
	"7de2f06498b5b4d53b170000c311101b55046a3c889efd54351cb3697fcf57cc",
	"8888cd035a543f3c28b335c2afe8a09fc49530f37380271a5b1ba1925b5f355a",
	"eb98b6704c17b6add5642ea05e12e5ad2d3c94c39b9950dacfb9bea8882b9a98",
}

func TestSHA256Hash(t *testing.T) {
	type input struct {
		seed       string
		iterations int
	}

	tests := []struct {
		input input
		ok    string
	}{
		{
			input: input{
				seed:       inputStrings[0],
				iterations: inputIterations[0],
			},
			ok: expectedResultsSHA256[0],
		}, {
			input: input{
				seed:       inputStrings[0],
				iterations: inputIterations[1],
			},
			ok: expectedResultsSHA256[1],
		}, {
			input: input{
				seed:       inputStrings[0],
				iterations: inputIterations[2],
			},
			ok: expectedResultsSHA256[2],
		}, {
			input: input{
				seed:       inputStrings[1],
				iterations: inputIterations[0],
			},
			ok: expectedResultsSHA256[3],
		}, {
			input: input{
				seed:       inputStrings[1],
				iterations: inputIterations[1],
			},
			ok: expectedResultsSHA256[4],
		}, {
			input: input{
				seed:       inputStrings[1],
				iterations: inputIterations[2],
			},
			ok: expectedResultsSHA256[5],
		},
	}

	h := SHA256{}

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
			t.Errorf(`[sha256] Hash([]byte(%s)) x %v = %s ; expected %s`,
				test.input.seed,
				test.input.iterations,
				result,
				test.ok,
			)
		}
	}
}
