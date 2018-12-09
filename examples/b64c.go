// written by "xeuus" aka aryan

package main

import (
	"github.com/xeuus/gocry/pkg"
	"fmt"
	"log"
)

func main() {

	// this method is nothing but a shuffled version of base64 encoding,
	// but also, works in Javascript or Typescript or any other language,
	// implementation is easy and needs not effort.
	// make sure you share this cipher with other client.
	cipher := cry.Base64Cipher()
	fmt.Println("Cipher: ", string(cipher))

	// we have a message and we want to send it for someone!
	myMessage := "Hello World!"
	fmt.Println("Message: ", myMessage)

	// first encrypt it, this method takes an []byte array
	// so message could be a blob of data, whatever
	enc := cry.Base64Encode([]byte(myMessage), cipher)

	fmt.Println("Encoded: ", string(enc))
	// assume we have received a message from other client,
	// which we know it's encrypted using the same cipher,
	// we just gonna decrypt it
	dec, err := cry.Base64Decode(enc, cipher)
	if err != nil {
		log.Fatal(err)
	}

	// here we get our message back.
	fmt.Println("Decoded: ", string(dec))

	// let's go and run it, go run $GOPATH/src/github.com/xeuus/gocry/examples/b64c.go
}
