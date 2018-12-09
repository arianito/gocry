// written by "xeuus" aka aryan

package main

import (
	"github.com/xeuus/gocry/pkg"
	"fmt"
	"log"
)

func main() {
	// scenario: bob want's to send a message to alice,
	// so all they need is a common password
	shared := "some random password !"
	// it's bob now speaking :)
	bobMessage := "It's bob, just wanted to say hello 8)"
	fmt.Println("Message: ", bobMessage)

	// bob have to encrypt his message, let's see
	enc, err := cry.AesEncode([]byte(bobMessage), shared)

	if err != nil {
		// in case something goes wrong
		log.Fatal(err)
	}

	fmt.Println("Encoded: ", string(enc))


	// well on the other side, i'm alice :) i'm the other one
	// bob sent me a encrypted message, let's see,
	dec, err := cry.AesDecode(enc, shared)

	if err != nil {
		// in case something goes wrong
		log.Fatal(err)
	}

	// so finaly, we could read our message
	fmt.Println("Decoded: ", string(dec))


	// let's go and run it, go run $GOPATH/src/github.com/xeuus/gocry/examples/aes.go


}
