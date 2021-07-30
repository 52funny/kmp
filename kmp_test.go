package kmp

import (
	"math/rand"
	"testing"
	"time"
)

var letterRunes = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func BenchmarkKMP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s, pattern := RandStringRunes(10), RandStringRunes(1)
		// s, pattern := "abdsaf", "saf"
		New(pattern)
		b.StartTimer()
		Index(s, pattern)
	}
}

func RandStringRunes(n int) string {
	rand.Seed(time.Now().Unix())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
