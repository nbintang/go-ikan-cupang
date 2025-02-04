package helper

import (
	"crypto/rand"
	"time"
)

const otpChars = "0123456789"

func GenerateOTP() (string, error) {
	maxOTPLength := 6
	buffer := make([]byte, maxOTPLength)
	_, err := rand.Read(buffer)

	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < maxOTPLength; i++ {
		buffer[i] = otpChars[int(buffer[i]%byte(otpCharsLength))]
	}
	return string(buffer), nil
}

func GenerateExpiration() int64 {
	return time.Now().Add(15 * time.Minute).Unix()	
}
