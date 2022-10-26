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
func convertByteToString(byteTochnage []byte) string {
	return hex.EncodeToString(byteTochnage)
}

func convertStringToByte(stringTochange string) ([]byte, error) {
	b, err := hex.DecodeString(stringTochange)
	if err != nil {
		panic(err)
	}
	return b, nil
}

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

	encryptedText := aesGCM.Seal(nonce, nonce, plaintext, nil)

	return encryptedText

}

// decrypt
func decrypt(encryptedString string, keyString string) string {
	// decode string to byte
	key, _ := hex.DecodeString(keyString)
	encryptedText, _ := hex.DecodeString(encryptedString)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	aesGCM, err := cipher.NewGCM(block)

	if err != nil {
		panic(err)

	}

	nonceSize := aesGCM.NonceSize()

	nonce, ciphertext := encryptedText[:nonceSize], encryptedText[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)

	return convertByteToString(plaintext)

}

func main() {

}
