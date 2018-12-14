package cry

/*
#include "b64c.h"
*/
import "C"
import (
	"unsafe"
	"bytes"
)

const b64cipher = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_="

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
	b.WriteString(" = ")
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


func Base64Set(cipher []byte) bool {
	return C.set_cipher((*C.char)(unsafe.Pointer(&cipher[0]))) != 0
}

func Base64Encode(message []byte) []byte {
	n := len(message)
	olen := (n + 2) / 3 * 4
	op := make([]byte, olen)
	C.base64_encode(
		(*C.char)(unsafe.Pointer(&message[0])),
		C.int(n),
		(*C.char)(unsafe.Pointer(&op[0])),
		C.int(olen),
	)
	return op
}

func Base64Decode(encoded []byte) []byte {
	n := len(encoded)
	olen := n / 4 * 3

	pad := byte(C.get_pad())
	if encoded[n-1] == pad {
		olen--
	}
	if encoded[n-2] == pad {
		olen--
	}
	op := make([]byte, olen)

	result := C.base64_decode(
		(*C.char)(unsafe.Pointer(&encoded[0])),
		(C.int)(n),
		(*C.char)(unsafe.Pointer(&op[0])),
		C.int(olen),
	)

	if result == 0 {
		return nil
	}

	return op
}
