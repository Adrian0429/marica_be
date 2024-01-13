package utils

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateOTP() string {
	// Seed the random number generator with the current Unix timestamp
	rand.Seed(time.Now().UnixNano())

	// Generate a random 5-digit OTP as an integer
	otp := rand.Intn(90000) + 10000

	// Convert the integer OTP to a string
	otpString := strconv.Itoa(otp)

	return otpString
}
