package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	for i := 0; i < 4; i++ {
		fmt.Println(rand.Int64())
	}
	r1 := rand.New(rand.NewPCG(1, 0))
	for i := 0; i < 5; i++ {
		fmt.Println(r1.Int64())
	}

	chacha8seed := [32]byte([]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ123456"))
	r2 := rand.NewChaCha8(chacha8seed)
	for i := 0; i < 5; i++ {
		fmt.Println(r2.Uint64())
	}

	r3 := rand.NewZipf(r1, 3, 10, 100)
	for i := 0; i < 5; i++ {
		fmt.Println(r3.Uint64())
	}
}
