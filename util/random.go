package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomInt(min, max int) int64 {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	return int64(min + r.Intn(max-min+1))
}

func RandomString(n int) string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	var sb strings.Builder
	k := len(alphabet)

	for range n {
		c := alphabet[r.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	currencies := []string{"EUR", "USD"}

	n := len(currencies)
	return currencies[r.Intn(n)]
}
