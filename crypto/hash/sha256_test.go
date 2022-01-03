package hash

import (
	"testing"
)

func TestSHA256Hash(t *testing.T) {
	type input struct {
		seed       string
		iterations int
	}
	type test struct {
		input input
		ok    string
	}

	var expectedResults []string = []string{
		"7f83b1657ff1fc53b92dc18148a1d65dfc2d4b1fa3d677284addd200126d9069",
		"4163fb4ab9e1e0a51709a51bc7e13ab6792907905960145c722d2c1479caac42",
		"04d516e1731d60e8bccf12843f9f7a3f065e137ed0d06804b16c6fd1461f7c85",
		"4239c3ec2221ee9ad2269551189c43f99feb2d9a79377844887b53af531d49bc",
		"58150621e8b8fec56ea75d2e1f354c7b636e0e170d712e14c8fcb5af533eb273",
		"415e2f87fb18c06fba57a214147905a7a2c22ec727c5340a23a21b7197015760",
		"c81498adc67c4b79d254169326e15dbc4da53c6a97d61efe7d32ea75ca5cb4a6",
		"44767119f3bbf94a9d45e3fd4c51c15e3d8f2287dc126ad3646e54a32f55962c",
		"a7cfcb61212352ce5e96e65e5a4d5998787410686735e5597b864303f7a96984",
		"1fada6a9b8084eff6347baeae0812aac1c14fcc0a97759a6bfa2d8a2ea087705",
		"7de2f06498b5b4d53b170000c311101b55046a3c889efd54351cb3697fcf57cc",
		"8888cd035a543f3c28b335c2afe8a09fc49530f37380271a5b1ba1925b5f355a",
		"47475db6736c25d5cdbb743ccc46fcc78bd39d46a3c4f27b79544f44fba4e356",
		"742e8c0d65d84c00dde3ae261b55eddba38b3d44e31a2816408eb24232771ca2",
		"a42fd5b823f229fe4d04cc43844d019ef389fe74e7cce66cd1addb11016a732a",
		"e6c66fcfa30c76c822e140a0bf6c9442eeffb68593b93bd9b1c855eed9f81008",
		"a2d837d6ddfefe127840f8f22104e8fc8fcae13b40d23c665fbcc07aadbf5ac1",
		"b5bf1e91ce14bb0384d970cf80b9021268e276e52d09de201acc7e8e191cc919",
		"3e8b4ca2f8206ff2b325073775437eeda8a1d30a6fa4b76339f992a7734fd9aa",
		"eb98b6704c17b6add5642ea05e12e5ad2d3c94c39b9950dacfb9bea8882b9a98",
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

func TestSHA224Hash(t *testing.T) {
	type input struct {
		seed       string
		iterations int
	}
	type test struct {
		input input
		ok    string
	}

	var expectedResults []string = []string{
		"4575bb4ec129df6380cedde6d71217fe0536f8ffc4e18bca530a7a1b",
		"91d70884daa33d890748aab078779e181a1f4543b3f6e2b7a1fe1e98",
		"55c0d229a82352c4145610ef132f87c90b0bd335d319a2ba43db8545",
		"efc8fb322c218891bcc6c97d5a292a721d333b07675b8d9d9ea98b53",
		"2324c0564beaa0c097528fffcb5e97ed092df97952660ff6221f3f6b",
		"0d41fa622f2fbb419767d65b2aefbde1c1035c6431b3a0bdaafa5dac",
		"1b4ca592c6fd7828d8190144d6a88ce29f5b02a777086991bbff1c77",
		"940fee95f71631243c637937a76ca320c70208eeac392e4185da2104",
		"515691ff92870a63786e3f70bdbbe76136ea4b752bc0f835e75cc10d",
		"ec014e7eab99b14277c94e7c9ac5f899af2e3f990fca2e3ccfbdb58f",
		"757d172373ff6679c334bbe0896a5270ed931719819b19a3ecc12faf",
		"8a5b20ccfc3bf4e04a75eab4939c19b1b8ecb943733512d0c8ba7a21",
		"f4f308d120b5d76cfe5ed3058b9daa1b74b989ab82994b86965667e6",
		"57811c0c5d494aff84f8adc95919e15b02bf69e436e4d2e3702d109f",
		"aa3bed715831ab53d017a76d24e78409d82571b219f004018d127b90",
		"f61d3e9da5f1242e72ad9f99299e954fcb0aa0c360464b11daa1cf53",
		"9b194a5afb9f90bcfefeeca280b474c0e032ae6955c9378e5d71249a",
		"35a3775764baba577dbbc04d91f1cc73f4bb58a6aa46f095c0109f65",
		"9c0ad0a706d694a9bc862fb1916bea14105cef8d88c92200d96931f4",
		"fb894fa4deaa72b0a6770d9efbafa771edb8aca3ce92073b44d4c823",
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

	h := SHA224{}

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
			t.Errorf(`[sha224] Hash([]byte(%s)) x %v = %s ; expected %s`,
				test.input.seed,
				test.input.iterations,
				result,
				test.ok,
			)
		}
	}
}
