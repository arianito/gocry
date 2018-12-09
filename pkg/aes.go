package cry

import (
	"crypto/aes"
	"crypto/cipher"
	"io"
	"crypto/rand"
	"crypto/md5"
	"encoding/hex"
)



func createHash(key string) string {
	hsh := md5.New()
	hsh.Write([]byte(key))
	return hex.EncodeToString(hsh.Sum(nil))
}

// Encode blob of data with a password
func AesEncode(message []byte, pass string) ([]byte, error) {
	block, _ := aes.NewCipher([]byte(createHash(pass)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	ciphertext := gcm.Seal(nonce, nonce, message, nil)
	return ciphertext, nil
}

// Decode blob of data with a password
func AesDecode(encoded []byte, pass string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(createHash(pass)))
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	nonce, cpt := encoded[:nonceSize], encoded[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, cpt, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil

}
