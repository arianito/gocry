# GOCRY (Painless Encryption)
Minimal encoding/decoding wrapper written for go
including base64, AES, RSA, serialize/deserializing toolkits.

### HOW TO INSTALL
```
go get github.com/xeuus/gocry
```

### HOW TO RUN EXAMPLES
```
go run $GOPATH/src/github.com/xeuus/gocry/examples/serialize.go
go run $GOPATH/src/github.com/xeuus/gocry/examples/rsa.go
go run $GOPATH/src/github.com/xeuus/gocry/examples/b64c.go
go run $GOPATH/src/github.com/xeuus/gocry/examples/aes.go
```

### HOW TO RUN BENCHMARK
```
cd $GOPATH/src/github.com/xeuus/gocry && make benchmark
```
if you're looking for results:
```
go test ./pkg -bench=.
goos: darwin
goarch: amd64
pkg: github.com/xeuus/gocry/pkg
Benchmark_Aes_Encode_Decode-4           	  300000	      3609 ns/op
Benchmark_Base_64_Cipher_Generation-4   	 2000000	       855 ns/op
Benchmark_Base_64_Encode_Decode-4       	 3000000	       427 ns/op
Benchmark_Rsa_Key_Generation-4          	      10	 299213612 ns/op
Benchmark_Rsa_Encode_Decode-4           	     100	  11808003 ns/op
Benchmark_Serialize_Deserialize-4       	   50000	     35508 ns/op
PASS
ok  	github.com/xeuus/gocry/pkg	14.745s

```

### HOW TO ADD IT TO YOUR WEBSITE
plugin located in dist/b64.min.js
```
<script src="b64.min.js"></script>
<script>
base64encode(message: string, cipher: string): string;
base64decode(encoded: string, cipher: string): string;
</script>
```

## WHATS IN THE BOX
```go
// RSA signature and encoded data block
type cry.EncryptedMessage struct {
	Signature []byte
	Encoded   []byte
}
// pair of keys in bytes
type cry.KeyPair struct {
	PublicKey  []byte
	PrivateKey []byte
}
// aes encryption/decryption
func cry.AesEncode(message []byte, pass string) ([]byte, error){}
func cry.AesDecode(encoded []byte, pass string) ([]byte, error){}

// generates a new random cipher for base64 encodings
// use []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_.") for standard b64 encoding.
func cry.Base64Cipher() []byte{}
func cry.Base64Encode(message []byte, cipher []byte) []byte{}
func cry.Base64Decode(encoded []byte, cipher []byte) ([]byte, error){}

// a wrapper for bcrypt
func cry.Bcrypt(password string) string{}
func cry.Check(hashed string, password string) bool{}

// rsa wrapper for encoding/decoding
func cry.RsaKeyPair() *KeyPair{}
func cry.RsaEncode(message []byte, privateKey []byte, destinationPublicKey []byte) (*EncryptedMessage, error){}
func cry.RsaDecode(message *EncryptedMessage, privateKey []byte, destinationPublicKey []byte) ([]byte, error){}

// in case you want to send or receive any type of struct
func cry.Serialize(a interface{}) ([]byte, error){}
func cry.Deserialize(a []byte, o interface{}) error{}

```

## BASE 64
here is the example:
```go
cipher := cry.Base64Cipher()
message := "Hello World!"
enc := cry.Base64Encode([]byte(message), cipher)
dec, _ := cry.Base64Decode(enc, cipher)
fmt.Println("Decoded: ", string(dec))
```
i also wrote this function for client side for web applications, same functionality works for web too,
```html
<script>
var receivedCipher = "...";
const encoded = base64encode("some message.", receivedCipher);
send_data_back_to_server(encoded);
</script>
```

## AES
here is the example:
```go
pass := "some random password !"
message := "Hello world"
enc, _ := cry.AesEncode([]byte(message), pass)
dec, _ := cry.AesDecode(enc, pass)
fmt.Println("Decoded: ", string(dec))
```

## RSA
here is the example:
```go
bob := cry.RsaKeyPair()
alice := cry.RsaKeyPair()
bobMessage := "It's bob"
enc, _ := cry.RsaEncode([]byte(bobMessage), bob.PrivateKey, alice.PublicKey)
dec, _ := cry.RsaDecode(enc, alice.PrivateKey, bob.PublicKey)
fmt.Println("Decoded: ", string(dec))
```

## Serialize/Deserialize
here is the example:
```go
type SomeMessage struct {
	Field1 string
	Field2 int
	Field3 map[string]string
}
message := &SomeMessage{
  Field1: "hello world!",
  Field2: 32982,
  Field3: map[string]string{
    "author": "someone",
    "presence": "no",
  },
}
serialized, err := cry.Serialize(bobMessage)
cipher := cry.Base64Cipher()
enc := cry.Base64Encode(serialized, cipher)
dec, _ := cry.Base64Decode(enc, cipher)

receivedMessage := new(SomeMessage)
cry.Deserialize(dec, receivedMessage)

```