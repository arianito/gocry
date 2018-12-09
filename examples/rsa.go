// written by "xeuus" aka aryan

package main

import (
	"github.com/xeuus/gocry/pkg"
	"fmt"
	"log"
)

func main() {
	// scenario: bob want's to send a message to alice,
	// first they have to exchange their public keys,
	// private key must kept private,
	bob := cry.RsaKeyPair()
	alice := cry.RsaKeyPair()

	// it's bob now speaking :)
	bobMessage := "It's bob, just wanted to say hello 8)"
	fmt.Println("Message: ", bobMessage)

	// bob have to encrypt his message, let's see
	// it's bob actually, his doing this, see bob only have alice's public key, no harm!
	enc, err := cry.RsaEncode([]byte(bobMessage), bob.PrivateKey, alice.PublicKey)

	if err != nil {
		// in case something goes wrong, shit happens, you know.
		log.Fatal(err)
	}

	fmt.Println("Signature: ", string(enc.Signature))
	fmt.Println("Encoded: ", string(enc.Encoded))
	// bob could use any protocol to send his message, could be base64?
	// it's already included in this package.


	// well on the other side, i'm alice :) i'm the other one
	// bob sent me a encrypted message, let's see,
	// i have bob's public key, he gave it to me, so:
	dec, err := cry.RsaDecode(enc, alice.PrivateKey, bob.PublicKey)

	if err != nil {
		// in case something goes wrong
		log.Fatal(err)
	}

	// so finaly, we could read our message
	fmt.Println("Decoded: ", string(dec))


	// let's go and run it, go run $GOPATH/src/github.com/xeuus/gocry/examples/rsa.go


}
