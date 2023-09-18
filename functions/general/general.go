package general

import (
	"net/url"
	"strings"
)


func IsValidURL(input string) bool {
	_, err := url.ParseRequestURI(input)
	return err == nil
}

func IsNotValidURL(input string) bool {
	_, err := url.ParseRequestURI(input)
	return err != nil
}

func ConvertToBase64ID(counterValue int64) string {

	const base64Chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	// Define the base and initialize an empty result string.
	base := int64(len(base64Chars))
	result := ""

	// Encode the counterValue into custom Base64.
	for counterValue > 0 {
		index := counterValue % base
		result = string(base64Chars[index]) + result
		counterValue /= base
	}

	// Ensure the result is exactly 7 characters long by padding with 'A' characters.
	padding := 7 - len(result)
	if padding > 0 {
		result = strings.Repeat("A", padding) + result
	}

	return result
}
