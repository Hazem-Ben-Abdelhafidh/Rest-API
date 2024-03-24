package utils

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz01234565789"

func RandomString(n int) string {
	var sb strings.Builder
	alphabetLength := len(alphabet)
	for i := 0; i < n; i++ {
		randomNumber := rand.Intn(alphabetLength)
		sb.WriteByte(alphabet[randomNumber])
	}
	return sb.String()
}
