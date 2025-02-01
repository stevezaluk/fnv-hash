package prime

import (
	"math/rand/v2"
	"strconv"
)

const (
	primeBit32   uint16 = 32
	primeBit64   uint16 = 64
	primeBit128  uint16 = 128
	primeBit256  uint16 = 256
	primeBit512  uint16 = 512
	primeBit1024 uint16 = 1024
)

/*
randomRange - Generate a pseudo-random number between the min and max provided
*/
func randomRange(max uint64, min uint64) uint64 {
	return rand.Uint64N(max-min) + min
}

/*
countBitValues - Count's the number of 1-bits in an 64-bit unsigned integer
*/
func countBitValues(number uint64) uint64 {
	oneBit := 0

	binary := strconv.FormatInt(int64(number), 2)

	for i := 0; i < len(binary); i++ {
		if string(binary[i]) == "1" {
			oneBit++
		}

	}

	return uint64(oneBit)
}
