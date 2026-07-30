package hash

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateHash() (string, error) {
	byteSlice := make([]byte, 32)
	_, err := rand.Read(byteSlice)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(byteSlice), nil
}
