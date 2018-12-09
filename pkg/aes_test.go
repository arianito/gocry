package cry

import "testing"


func Benchmark_Aes_Encode_Decode(b *testing.B) {
	cipher := "some password !"
	myMessage := "Lorem ipsum"
	for i := 0; i < b.N; i++ {
		enc, _ := AesEncode([]byte(myMessage), cipher)
		AesDecode(enc, cipher)
	}
}