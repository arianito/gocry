package cry

/*
#include "b64c.h"
*/
import "C"
import (
	"unsafe"
	"errors"
	"bytes"
)

const b64cipher = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_."

func Base64Cipher() []byte {
	data := make([]byte, 65)
	copy(data, b64cipher)
	C.shuffle((*C.char)(unsafe.Pointer(&data[0])), 65)
	return data
}

func Base64Scrabmle(data []byte) string {
	n := len(data)
	v := make([]byte, 13)
	copy(v, "0123456789ABC")
	C.shuffle((*C.char)(unsafe.Pointer(&v[0])), 13)
	a := new(bytes.Buffer)
	b := new(bytes.Buffer)
	b.WriteString("const _")
	b.WriteString(RandomString(5))
	b.WriteString( " = ")
	i := 0
	j := 0

	bc := make([]byte, 5)

	for i < n {
		bc[0] = data[i]
		i++
		bc[1] = data[i]
		i++
		bc[2] = data[i]
		i++
		bc[3] = data[i]
		i++
		bc[4] = data[i]
		i++

		bcp := RandomString(5)
		a.WriteString("const _")
		a.WriteString(bcp)
		a.WriteString(" = \"")
		a.Write(bc)
		a.WriteString("\"\n")
		//
		b.WriteString("_")
		b.WriteString(bcp)
		b.WriteString(" + ")

		j++
	}
	a.WriteString(string(b.Bytes()[:b.Len()-2]))

	return string(a.Bytes())
}

func Base64Encode(message []byte, cipher []byte) []byte {
	n := len(message)
	op := make([]byte, (n+2)/3*4)
	C.base64_encode(
		(*C.char)(unsafe.Pointer(&message[0])),
		C.int(n),
		(*C.char)(unsafe.Pointer(&cipher[0])),
		(*C.char)(unsafe.Pointer(&op[0])),
	)
	return op
}

func Base64Decode(encoded []byte, cipher []byte) ([]byte, error) {
	n := len(encoded)
	op := make([]byte, n/4*3)

	result := C.base64_decode(
		(*C.char)(unsafe.Pointer(&encoded[0])),
		(C.int)(n),
		(*C.char)(unsafe.Pointer(&cipher[0])),
		(*C.char)(unsafe.Pointer(&op[0])),
	)

	if result == 0 {
		return nil, errors.New("something went wrong")
	}

	return op, nil
}
