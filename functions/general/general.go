package general

import (
	"encoding/base64"
	"fmt"
)

// TODO: make this proper
func ConvertToBase64ID(counterValue int64) string {
	// Convert the counter value to a 7-character Base64-encoded ID.
	counterBytes := []byte(fmt.Sprintf("%07d", counterValue))
	encodedID := base64.StdEncoding.EncodeToString(counterBytes)

	// Trim any trailing padding characters "=" from the Base64 encoding.
	return encodedID[:7]
}