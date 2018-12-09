package cry

/*
#include "b64c.h"
*/
import "C"
import (
	"unsafe"
	"errors"
)

const b64cipher = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_."

func Base64Cipher() []byte {
	data := make([]byte, 65)
	copy(data, b64cipher)
	C.shuffle((*C.char)(unsafe.Pointer(&data[0])), 65)
	return data
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
