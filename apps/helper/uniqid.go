package helper

import (
	"crypto/rand"
	"encoding/hex"
)

func UniqID() string {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}
