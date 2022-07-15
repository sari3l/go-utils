package random

import (
	"math/rand"
	"time"
)

func Choice[T any](seq []T) T {
	SetSeed(UnixSeed())
	return seq[_rand(len(seq))]
}

func Choices[T any](seq []T, k int) []T {
	var t = []T{}
	for i := 0; i < k; i++ {
		t = append(t, Choice(seq))
	}
	return t
}

func SetSeed(seed int64) {
	rand.Seed(seed)
}

func _rand(n int) int {
	return rand.Intn(n)
}

func UnixSeed() int64 {
	return time.Now().UnixNano()
}
