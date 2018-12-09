package cry

import "testing"

func Benchmark_Rsa_Key_Generation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RsaKeyPair()
	}
}

func Benchmark_Rsa_Encode_Decode(b *testing.B) {
	bob := RsaKeyPair()
	alice := RsaKeyPair()
	myMessage := "Lorem ipsum"

	for i := 0; i < b.N; i++ {
		enc, _ := RsaEncode([]byte(myMessage), bob.PrivateKey, alice.PublicKey)
		RsaDecode(enc, alice.PrivateKey, bob.PublicKey)
	}
}
