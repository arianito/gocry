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
or if you have webpack environment, you can use obfuscate plugin,
and use scramble function to generate a js cipher for you
```
cipher := cry.Base64Cipher()
cry.Base64Scrabmle(cipher)
// output
// const _IHRIu = "jFfX0"
// const _BwIUb = "Ghun2"
// const _nAC3q = "4EmTI"
// const _qZEwV = "QV8Or"
// const _PXLsk = "PlxS9"
// const _aQ2Tr = "YKtep"
// const _k7134 = "vW6Mo"
// const _3dTud = "yUDJ."
// const _NODaV = "NcdL1"
// const _11KEO = "7z3Bb"
// const _bPnwW = "sZ_-i"
// const _3lwMI = "qakwH"
// const _6ZJMq = "ACg5R"
// const _ITLzh = _IHRIu + _BwIUb + _nAC3q + _qZEwV + _PXLsk + _aQ2Tr + _k7134 + _3dTud + _NODaV + _11KEO + _bPnwW + _3lwMI + _6ZJMq
// _ITLzh is your cipher

``` 
and here is how after obfuscation:
```javascript
var _0x3cc7=["iterator","xECgk","wrzSo","kGfFo","vlpqU","label","ops",
"mkkFR","nCCaS","RFWbO","trys","sObIo","SohCv","isSynced","diff","Yntqv",
"_instance","sync","oILbb","CgNDw","now","sent","cJcuW","QaNcH","schedule",
"LfTQk","zMswW","Time","UdjVA","iKURW","tVvsi","2|3|4|1|0","extend","merge",
"isCancel","all","spread","ICpDd","XjiuN","Okgcq","ZmJOn","readFloatLE","JGtcS",
"3|4|0|8|2|5|6|1|7","vnYki","defaults","request","qkIOg","method","sUeVo","EfsUQ",
"MprnO","fulfilled","response","rejected","xeyQI","sRhOp","HuXXu",
"clearTimeout has not been defined","browser","setTimeout has not been defined","nQMak",
"CnDTK","vMwiQ","HmHyn","xiQoi","qtnzq","vKImk","nextTick","SpOfT","xNDeR","fmYOi","NxLIq",
"fun","array","run","title","zYULS","env","argv","version","addListener","once","removeListener",
"removeAllListeners","prependListener","prependOnceListener","listeners","binding",
"process.binding is not supported","cwd","qRMum","umask","XhCDl","HKvUM","wvFKT","ltlpC",
"PKVzC","config","validateStatus","zDUKB","cxFMr","KpGDX","Request failed with status code ",
"4|2|1|3|0","DAlUW","code","XtWMb","dgNdA","iwIzp","PmQZh","xZOoB","toISOString","isObject","AaVqe",
"elvIG","gVtOg","VpQeO","set-cookie","age","authorization","content-length","expires","proxy-authorization",
"retry-after","user-agent","yRNpC","zasjE","ADWAI","fGinw","FrvQC","ImnUj","tvChl","YDxVD","fSNll","asYFk","OJxvQ",
"ZZTqP","from","last-modified","NBXAh","max-forwards","vbDLF","referer","BQCSK","href","EPFDk","isStandardBrowserEnv",
"gkFjw","userAgent","lpZvp","protocol","hostname","port","isString","iKxRn","TMGQv","host",...];
```

well now, it's just harder to find your cipher, but thief can


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
// obfuscate cipher to build in 
func Base64Scrabmle(data []byte) string;
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