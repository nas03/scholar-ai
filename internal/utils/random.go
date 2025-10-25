package utils

import (
	"crypto/rand"
	"math/big"
)

func GenerateSixDigitOtp() int {
	// Generate a random number between 100000 and 999999 (6 digits)
	max := big.NewInt(900000) // 999999 - 100000 + 1
	n, _ := rand.Int(rand.Reader, max)
	return int(n.Int64()) + 100000
}
