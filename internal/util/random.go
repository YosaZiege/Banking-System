package util

import (
	crand "crypto/rand"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var randSrc *rand.Rand

func init() {
	randSrc = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomString(length int) string {
	result := make([]byte, length)
	for i := range result {
		index := randSrc.Intn(len(alphabet))
		result[i] = alphabet[index]
	}
	return string(result)
}

func RandomInit(min, max int64) int64 {
	return min + randSrc.Int63n(max-min+1)
}

func RandomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// RandomUsername returns a random username + 4-digit number
func RandomUsername() string {
	usernames := []string{
		"ninja", "shadow", "warrior", "falcon", "panther",
		"samurai", "ghost", "hunter", "viper", "phoenix",
	}
	base := usernames[randSrc.Intn(len(usernames))]
	suffix := randSrc.Intn(10000)
	return fmt.Sprintf("%s%04d", base, suffix)
}

// RandomEmail returns a random email like ninja2341@gmail.com
func RandomEmail() string {
	name := RandomUsername()
	domains := []string{"gmail.com", "yahoo.com", "protonmail.com", "hotmail.com"}
	domain := domains[randSrc.Intn(len(domains))]
	return fmt.Sprintf("%s@%s", strings.ToLower(name), domain)
}

// RandomPasswordHash returns a random 128-bit hex string (simulating hashed passwords)
func RandomPasswordHash() string {
	bytes := make([]byte, 16)
	_, _ = crand.Read(bytes)
	return hex.EncodeToString(bytes)
}

// RandomProvider returns a random auth provider
func RandomProvider() string {
	providers := []string{"google", "github", "none", "facebook"}
	return providers[randSrc.Intn(len(providers))]
}
