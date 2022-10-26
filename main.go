package main

import (
	"crypto/rand"
	"encoding/hex"
)

func generateKey() string {
	keyB := make([]byte, 32)
	_, err := rand.Read(keyB)
	if err != nil {
		panic(err)
	}
	key := hex.EncodeToString(keyB)
	return key
}
func main() {

}
