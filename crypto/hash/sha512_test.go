package hash

import (
	"testing"
)

func TestSHA512Hash(t *testing.T) {
	type input struct {
		seed       string
		iterations int
	}
	type test struct {
		input input
		ok    string
	}

	var expectedResults []string = []string{
		"861844d6704e8573fec34d967e20bcfef3d424cf48be04e6dc08f2bd58c729743371015ead891cc3cf1c9d34b49264b510751b1ff9e537937bc46b5d6ff4ecc8",
		"597a546472ed4f6a2cf8b67a1ae21fb84705bcb6cfc14098f91e60b2e041cde5905f146cf3ff73d634468867d50b550e70a9e2ee1effa185abf7214cbaa6231d",
		"5ad946cecaa1964f86cabdf4daf08321e8a1d88755d4bd931c63684f40e5b61e0c1ba9c504769cfa5c2e34e559b3368f68d0d52a675b2f5ca60eef22d33ba8e3",
		"3d965d197e0b0cdb241d66ad0776d357688994ffcac4f5810e1d2c1bfe5d373aeb0edc210ce4c8d178e2ed710f70b7735f478be657f1fb4f29400d043a0085ab",
		"bcab01e461c54ef89c197c744938a8514aa69f57d65f24f74da986582beef3e4e2d7c6dc5452bea9034d7104252ec2aaada034f84ce4434b833ef364e303b5df",
		"7edfd21f50676dc60d11d1513db34bc3222737041a02b11f72ee5feb8ac2805606f4ca12635108633d384f0c1d852d51c24fd5711460f457afaeff80aec8a098",
		"b42cdf3fed4a45edf312aae838c8a9b3a32ae77aeab5005a93480c1634ca8faecbc4ae978e9bece512a168d0cc2f311a089be31bbbb3fedecca221f1d1c09112",
		"b579e9faff3ab5383a1507f547aa3129f7f2ceb30baf4d2823dca775dcee8a331cf535b045be9e3e8cd8fa44642b178fd7f0867e5393357162c5fb6f1f74a393",
		"f651c62c9465ef99cc04e6fe6749731720143166bc0677f3e94e126b13a82dff86448690d7d3c7116c56abc48a0e38321fcba6d7357e5a2022ad6176c83f0ca5",
		"4d28ca363687a77cc144cdc3acf14c595b622c2b94988ea7568b6179d6ed767e8c149a12332cd39e0e51b0b1109c60d8f82d1bb2d5afd29e585bc4866ba8825b",
		"ac0b54775e8457dd24c19e5f167dfc1ee29bf0c8df2cfe97fab207108243e665e55418af5a981078ca3a5e36a01106a4f1727929c8022d118a25fb182ef8842e",
		"f50aef5b4191dd536321b47b818e426c3882cb27536581d4e2e895a56aff98221e76f00df0c885848e6898a922c795db59e2590372d9b7f638f2c3d773f555bb",
		"da9f10831ca45ce664c7de8144433d7361c90e368aaa72125e4bfb44c2d1408b793005c041dba2359a9fe2921ec9a425a2944b0bc6152c4f67f5f60c4d96f26e",
		"64b6a38e6d2485d45c1a08d059c0c444c715dcf3670f9fd7866f15a9d0c1e475c9c1765a21d9c0a85f17bd71e440c1908e52c990c93f213e311223daae62e719",
		"6c9b9d015e2acd01968b0dab03ea805200fe1fbbd488e663ae5cbe54668efcdf2ed15b901577234b9c405d8acdc486122f2ecf4fe61c6e76e97e2d514a8c989c",
		"db190ee728ddaa17e11c66b2f924d3828619bdcb6bffac0e46c4de4c2094325af64fb04ea29aa4673adafa26e259c1bbdb85d5145df97a1cf997b351a0c22da2",
		"8abae11414ae6901ea3afaccd70131a7c19b48b3cba062c93e4fab344a8e258f29e1073872029f0470758dec8cba6e56d3ce69ccc0b8faf4389295cb65c16948",
		"89d125d0440df8c7a9d2ac00b380e8ec5c0f55f0a19ef23959601cf34a5c24f3088d87044b4cd5f906791890c7868b81d72265d1dd5d77ff462726fa4b7fc595",
		"7d6acf7ef6ec1ca7a2cb8ae2fc7cf5f7c9fbe40533ed609fa5f68eb2f1011cc3a65d811966b51ae9035d44257ae9670bbc52ae24fdbcbfde22bc50bc00f66410",
		"ffbacca7b48b87d2c17050c81ba43024ffd1b29feeac2a72d6602d8aa62a995bc9dc55915a84c21fa63419734b4af7529e2686b31195cd39ffedb6d8d5d983ae",
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

	h := SHA512{}

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
			t.Errorf(`[sha512] Hash([]byte(%s)) x %v = %s ; expected %s`,
				test.input.seed,
				test.input.iterations,
				result,
				test.ok,
			)
		}
	}
}

func TestSHA384Hash(t *testing.T) {
	type input struct {
		seed       string
		iterations int
	}
	type test struct {
		input input
		ok    string
	}

	var expectedResults []string = []string{
		"bfd76c0ebbd006fee583410547c1887b0292be76d582d96c242d2a792723e3fd6fd061f9d5cfd13b8f961358e6adba4a",
		"aa690c4090ad23942f7a7426fd3174494ddadde1a26ddaf251a9ef4ec2c133c831b085dffafbd2c3b7eaf880869851ed",
		"5e169a1e30ee4525aee2acf05311544738cb3694da7ef60a9f956390178234d6b8233bf980bf6874ceb5ebf6ed62a894",
		"c3e4b47a5a120bc1acee4f69af92a1326d3bca4d3e4d5979077ebe6e1fe14ab0ccf2facc88f63d613ebbc63ce65deda1",
		"c5db6babe0600000120bba197eda755eb057d72a4ddfb9d719726f46b1b4e498aecd6d99f6e485b7bee8f8d347aa96ca",
		"d9284304fa228db07b49b8676eba26b9bfc6062962b5e094d306cf20a8e64ab361eea45cf161994eb4ce51420951b8ff",
		"79b28b728412e1e0af5cfd6de2e4a73604a25c85140dd7cf71a111b4329b847cc2fcf617f62e846dcd6de0f525a00fc0",
		"a7c4abb91fb5bfd17a7b580dc11f3ba78c61d24d77f6db967c3a651e6f9923a86df26a48bdad692fcca1645e94d15b80",
		"744a96fdd29efe850b03d4a3836952f5100972f1b3ea8babfba381f28ef674ecbce36c76f33ab7a04136b616ddaca70c",
		"ba6d6658c0fa6a88acc39a7e18ce2872ffc77a1c6693f587592b8109951ef8ec2c3c179207807f3ac52a1608c76f45cd",
		"585c4094891e7a11072fd0f02c347bfdf10d5043f460ac48db2699193129385bae50a8904d29708eb98c09085010805c",
		"b6e1e95ed4fa1a24774076d285e01b92c8e709d9d3a8316872bf3dd07be34b4bbe5cfbb13193818091b4d358ecba8c57",
		"7f7f1a8af9484a2a1428325cef0400dbea45e411778608f19698a92ffb8c0fd73a9f6bddc06a5cc502da9ade65c61c27",
		"79ab599c5bf285d38890eaaf2a36a2652ce86b00cbd2685367ca4187d7a8c511317b3b272454c99feb1d01d31f8ced13",
		"bc0451f67a0f64eb0b70831c0b71429487fd608268ce6d13855038f7d0de1aef8195a4d84c8d22e635e5628f38b90b4f",
		"c7134584a12bc1c4d0b66c0781f013302f5ebe534ec0973bf7671347ea89f076229991ccc626937d4ebb0d028becf957",
		"52f5a8f3c85060c5b4a24d30c0b68d1f8efd7986b2a1519759368846d333d5af40ba9c5e3f239b613255624c305bcc10",
		"f824a0d00bfaa1aed9a2157d091fa20d5002b73009fdb1e0a449bbad2b776cf35be0c3e08ed9d1be7cb2209ec63e517e",
		"4f5da5757b4ad68b6af3e372f1a6fb2a9011ff223faf99dafdbe73c4a81dc5d65d184b698f16db6006790e44761b96d8",
		"0a42c68fa59325c4f6ee5fa0124d7edfafc3d50e6498583f8135b16159aae7464076df8aa832a517683893cac1491035",
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

	h := SHA384{}

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
			t.Errorf(`[sha384] Hash([]byte(%s)) x %v = %s ; expected %s`,
				test.input.seed,
				test.input.iterations,
				result,
				test.ok,
			)
		}
	}
}

func TestSHA512_256Hash(t *testing.T) {
	type input struct {
		seed       string
		iterations int
	}
	type test struct {
		input input
		ok    string
	}

	var expectedResults []string = []string{
		"f371319eee6b39b058ec262d4e723a26710e46761301c8b54c56fa722267581a",
		"d5c30a09095eaa3edfe61e3519105c7ba15b9065787d192320a20a3894ae4c73",
		"09857a891600380ff53341bd3de25a57af133a45405cf8200172415948463dec",
		"74e21ecc53bc7f9edaa0ed2f86a83c69b63ef2aed17d7577f12a5e745654e56b",
		"07b35c74b6507d970a076a6cd6b5126d25664b6ff4446ec0a4be5646af7d6dcd",
		"09d6ec438f61bc1bfb802ebbbe52d4c19fb2d7381a071db99e1b43874896fdee",
		"58e564a6444cac00665a886b492adbe8dddc2be4b675e222a9174c338ad702b4",
		"ea8601967a5fbac169ac279ab8210cb3266b0d3db6cd9099415e55df96a8e133",
		"cb3080ae722e0eca974d66014789aba138e8cc584ee600bc624c4eb65aa7dffc",
		"26a41f5bd0bfa72a62a931b47dc44f075701704ec905dea1fceaa486afc96e21",
		"40cb4309fe37ff197a7b101b8f578ca03d8940c549deeebfe512453ca8328618",
		"a0844e29e58e14ccab09844b5587366839b7903101c3ddb4d96628707a8a6e43",
		"b1c8598080860683c11c952d8a2c179902b1b09afca525d636b693f9e390f7e4",
		"51d7bb18253ae4cea347904bdb08767fbe3f6cd658a273608d30ed6dc6851003",
		"78beddf16561ac4ac6a755911e6b17fc8ba9891ed192a032e267fac2de5b1e89",
		"7692189dc01ea8a29d14d806009c9d3ae685e0fac87657316b879f5bad247c4f",
		"741c5cdb3606893926cc8cf975241d11579ebc4d85bdb61efade6cd9cb507779",
		"b96c492cb1478647eb1513e0a10297aa5d9b88b5eeb5331f8839624395e53094",
		"74efa9014ad98a2214bc392dea9e1e3965ab6c489339bd27d27ab3a3fe65e4e1",
		"445bd9f4a78693bddbb57f7cf254e4b7ae31617be95317a6c106fb9b3339fb1d",
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

	h := SHA512_256{}

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
			t.Errorf(`[sha512_256] Hash([]byte(%s)) x %v = %s ; expected %s`,
				test.input.seed,
				test.input.iterations,
				result,
				test.ok,
			)
		}
	}
}

func TestSHA512_224Hash(t *testing.T) {
	type input struct {
		seed       string
		iterations int
	}
	type test struct {
		input input
		ok    string
	}

	var expectedResults []string = []string{
		"ba0702dd8dd23280b617ef288bcc7e276060b8ebcddf28f8e4356eae",
		"cf9f4f1192eedc0b99f504308e402c381dfe2607b04043ad6c536fc8",
		"e129f34a6e821cbf5e47a4d4fd03438c0f2895e5e6d1c5ce17f036e4",
		"400f6cf1c5592d7018f3ae741de2de9d2091299afc84508889050e3d",
		"0d308d6a3d8a770b32352d32d2cb0b25ad663ee691ac156eca70639c",
		"e4a23032051650c154908ffce2f00cf46a7ee92d85370f5cc8ebc8b5",
		"e082b96e4351e40aa0a3d164b1192f380fb0fc1e44f25cc62facce59",
		"ce5c154d63ec0676f83f42d77326d536fb64dc82c05a55c1705efe61",
		"486df12e5251b4551ebb7380a95b14a243e71d5ac493604f70a1b505",
		"92cdbdd6ad910ce9f0d773018e82cf72ecf3d3af157e0b2374919c7c",
		"33128956342949e073786bfd288f8a810575934c7bc6d1573f3aafe2",
		"f1b29d9f2e273ab346949f7584eed059e2e284a33f723e02322ab17f",
		"88dcf5c82037ea0b854d02bcb742f910a7cbd900d7ac29f40df840e9",
		"1d6efacd7361e6be49f467c9f62a496e29b654717e75edb291db02cf",
		"7b6efc42c1970ec086c794739138603b3fad28198e5eb29664a66aaa",
		"ac397d0d1f42de7fca08d119d8f4f2b57689fe6d9d4c5ab63bddbb5f",
		"7e2cad96c0d4f8d2c866dd6c350b13be4436bc5d1901dfa881850610",
		"6292a6134d2e1b61cb2cb573f7da278efd496802224b3ed4d437c0b0",
		"3eec164d262ffbc94b28ae8957715984d1db68975031961b98c6eb61",
		"0228d17becbeeaee1f6f15957c19b957dfc972b3b54d0e6ad1d2ad92",
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

	h := SHA512_224{}

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
			t.Errorf(`[sha512_224] Hash([]byte(%s)) x %v = %s ; expected %s`,
				test.input.seed,
				test.input.iterations,
				result,
				test.ok,
			)
		}
	}
}
