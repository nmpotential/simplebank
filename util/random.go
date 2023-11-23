package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	// Seed the random number generator with the current Unix timestamp in nanoseconds
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max (inclusive)
func RandomInt(min, max int64) int64 {
	// Generate a random integer in the specified range using rand.Int63n
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n using characters from the 'alphabet'
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		// Select a random character from the 'alphabet' and append it to the string builder
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name by creating a random string of length 6
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money (integer between 0 and 1000)
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency code from a predefined list
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// RandomEmail generates a random email address with a random string and a common domain
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
