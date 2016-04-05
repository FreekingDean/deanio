package deanio

import (
	"math/rand"
	"time"
)

func calcProb() float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()
}
