package cry

import (
	"crypto/rsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto"
	"encoding/pem"
	"crypto/x509"
	"os"
	"log"
	"errors"
)

type EncryptedMessage struct {
	Signature []byte
	Encoded   []byte
}

type KeyPair struct {
	PublicKey  []byte
	PrivateKey []byte
}

func RsaKeyPair() *KeyPair {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	privBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	})
	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&key.PublicKey),
	})

	return &KeyPair{
		pubBytes,
		privBytes,
	}
}

func readPublicKey(bytes []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(bytes)
	if block == nil {
		return nil, errors.New("error in parsing bytes, invalid rsa public key")
	}
	newKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return newKey, nil
}

func readPrivateKey(bytes []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(bytes)
	if block == nil {
		return nil, errors.New("error in parsing bytes, invalid rsa private key")
	}
	newKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return newKey, nil
}




func RsaEncode(message []byte, privateKey []byte, destinationPublicKey []byte) (*EncryptedMessage, error) {
	other, err := readPublicKey(destinationPublicKey)
	if err != nil {
		return nil, err
	}

	label := []byte("")
	hash := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(
		hash,
		rand.Reader,
		other,
		message,
		label,
	)

	if err != nil {
		return nil, err
	}

	mine, err := readPrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto
	newhash := crypto.SHA256
	pssh := newhash.New()
	pssh.Write(message)
	hashed := pssh.Sum(nil)

	signature, _ := rsa.SignPSS(
		rand.Reader,
		mine,
		newhash,
		hashed,
		&opts,
	)

	return &EncryptedMessage{
		signature,
		ciphertext,
	}, nil
}

func RsaDecode(message *EncryptedMessage, privateKey []byte, destinationPublicKey []byte) ([]byte, error) {

	mine, err := readPrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	hash := sha256.New()
	label := []byte("")

	plainText, err := rsa.DecryptOAEP(
		hash,
		rand.Reader,
		mine,
		message.Encoded,
		label,
	)
	if err != nil {
		return nil, err
	}

	other, err := readPublicKey(destinationPublicKey)
	if err != nil {
		return nil, err
	}

	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto
	newhash := crypto.SHA256
	pssh := newhash.New()
	pssh.Write(plainText)
	hashed := pssh.Sum(nil)

	err = rsa.VerifyPSS(
		other,
		newhash,
		hashed,
		message.Signature,
		&opts,
	)

	if err != nil {
		return nil, err
	}

	return plainText, nil
}