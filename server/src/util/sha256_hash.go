package util

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashSHA256(plainText string) string {
	hash := sha256.Sum256([]byte(plainText))
	return hex.EncodeToString(hash[:])
}
