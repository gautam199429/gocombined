package utilities

import (
	"crypto/sha256"
	"fmt"
)

func HashSHA256(input string) string {
	hash := sha256.Sum256([]byte(input))
	return fmt.Sprintf("%x", hash)
}
