package utils

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/goravel/framework/facades"
)

// GenerateResetToken creates a secure token for password resets.
func GenerateResetToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}

// AppURL returns the URL of the application.
func AppURL() string {
	config := facades.Config()
	defaultHost := config.GetString("http.host")
	defaultPort := config.GetString("http.port")
	return defaultHost + ":" + defaultPort
}
