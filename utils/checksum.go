package utils

import (
	"crypto/sha256"
	"fmt"
)

func SHA256Checksum(data []byte) string {
	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash)
}
