package cry

import (
	"testing"
	"encoding/base64"
)


func Benchmark_Base_64_Cipher_Generation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Base64Cipher()
	}
}



func Benchmark_Base_64_Encode_Decode(b *testing.B) {
	cipher := Base64Cipher()
	Base64Set(cipher)
	myMessage := "Lorem ipsum"
	for i := 0; i < b.N; i++ {
		enc := Base64Encode([]byte(myMessage))
		Base64Decode(enc)
	}
}

func Benchmark_GO_Base_64_Encode_Decode(b *testing.B) {
	myMessage := "Lorem ipsum"
	for i := 0; i < b.N; i++ {
		enc := base64.URLEncoding.EncodeToString([]byte(myMessage))
		base64.URLEncoding.DecodeString(enc)
	}
}