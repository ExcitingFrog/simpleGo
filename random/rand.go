package random

import (
	"math/rand"
	"time"
)

var (
	LetterNumberRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	LetterRunes       = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	NumberRunes       = []rune("0123456789")
)

func RandStringRunes(n int, runes []rune) string {
	var (
		b = make([]rune, n)
	)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}
