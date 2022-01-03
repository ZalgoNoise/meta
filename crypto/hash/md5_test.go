package hash

import (
	"testing"
)

func TestMD5Hash(t *testing.T) {
	type input struct {
		seed       string
		iterations int
	}
	type test struct {
		input input
		ok    string
	}

	var expectedResults []string = []string{
		"ed076287532e86365e841e92bfc50d8c",
		"153163e20c7dd03b131fe2bf21927e1e",
		"95b9b73c7a8abe6d5bf5ad8067a53cee",
		"d7b6778d149c5a085f2675b959503c2b",
		"4a3dd453c637dabbdce7433653e6d7cb",
		"88648d19a031133d87b093997813a740",
		"a07578c3f0bebca3f2cca2b1d0b8f9fe",
		"6134e6ee02e25ea0faf7a348ef599cdd",
		"188ee3fb92bb86c91a08198100fa167a",
		"dc06bc1fb41cae72a4b60e8480817085",
		"301e2149e9bfbe6c596aa631f6924697",
		"ce5e3afd4fb3b5e21a5e1ce99ab21acf",
		"19375359d7dd2f2d5b43888766cfcf05",
		"0e78e896041ecbd4b0f28cb0a2a98de6",
		"7b0c20a257e724db0e0e9be88dd67131",
		"e981affc8c88d89af7685a6d686cfe02",
		"bdddadc0e8b1b288da9f371955a8e7b7",
		"1e620b8a662ed76559b12d796b2aa51b",
		"23373334ceda8b84f9e2b2339d24c234",
		"fac2e231da55a96a9c2bfe0c2604dbab",
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

	h := MD5{}

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
			t.Errorf(`[md5] Hash([]byte(%s)) x %v = %s ; expected %s`,
				test.input.seed,
				test.input.iterations,
				result,
				test.ok,
			)
		}
	}
}
