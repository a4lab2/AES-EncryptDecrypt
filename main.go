package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
)

// String version of the key
func generateKey_String() string {
	keyB := make([]byte, 32)
	_, err := rand.Read(keyB)
	if err != nil {
		panic(err)
	}
	key := hex.EncodeToString(keyB)
	return key
}

// Overkill not really neccesary
// func convertKeyToByte() {

// }

func encryptS(stringToEncrypt string) []byte {
	keyString := generateKey_String()
	key, _ := hex.DecodeString(keyString)
	plaintext := []byte(stringToEncrypt)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)

	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}
	nonce := make([]byte, aesGCM.NonceSize())

	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		panic(err)
	}

	encryptextText := aesGCM.Seal(nonce, nonce, plaintext, nil)

	return encryptextText

}

// decrypt
func decrypt(cipherText []byte) {

}

func main() {

}
