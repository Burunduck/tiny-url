package utils

import (
	"hash/fnv"
	"math/rand"
	"time"
)


const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_"

func Hash(salt string) string {
	h := fnv.New64()
	h.Write([]byte(salt))
	rand.Seed(time.Now().UTC().UnixNano() + int64(h.Sum64()))
	b := make([]byte, 10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
