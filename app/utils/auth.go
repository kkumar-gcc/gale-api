package utils

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateResetToken creates a secure token for password resets.
func GenerateResetToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}
