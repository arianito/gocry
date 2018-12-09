// written by "xeuus" aka aryan

package main

import (
	"github.com/xeuus/gocry/pkg"
	"fmt"
	"log"
)


type SomeMessage struct {
	Field1 string
	Field2 int
	Field3 map[string]string
}


func main() {
	bob := cry.RsaKeyPair()
	alice := cry.RsaKeyPair()
	bobMessage := &SomeMessage{
		Field1: "hello world!",
		Field2: 12423,
		Field3: map[string]string{
			"are":     "you",
			"kidding": "me ?",
			"why ?":   "i'm exhausted writing this",
		},
	}
	fmt.Println("Message: ", bobMessage)

	// first we have to serialize bob's message
	serialized, err := cry.Serialize(bobMessage)

	if err != nil {
		log.Fatal("we cant serialize this message due to ", err)
	}
	enc, err := cry.RsaEncode(serialized, bob.PrivateKey, alice.PublicKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Signature: ", string(enc.Signature))
	fmt.Println("Encoded: ", string(enc.Encoded))
	dec, err := cry.RsaDecode(enc, alice.PrivateKey, bob.PublicKey)

	if err != nil {
		log.Fatal(err)
	}
	// here we got the message, we have to deserialize it
	deserialized := new(SomeMessage)
	err = cry.Deserialize(dec, deserialized)

	if err != nil {
		log.Fatal(err)
	}

	// so finaly, we could read our message
	fmt.Println("Decoded: ", deserialized)


	// let's go and run it, go run $GOPATH/src/github.com/xeuus/gocry/examples/serialize.go


}
