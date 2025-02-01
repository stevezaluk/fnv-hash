package prime

import (
	"math/rand/v2"
	"strconv"
)

/*
randomRange - Generate a pseudo-random number between the min and max provided
*/
func randomRange(max uint64, min uint64) uint64 {
	return rand.Uint64N(max-min) + min
}

/*
countBitValues - Count's the number of 1 and 0 bits in an unsigned 64-bit integer.
*/
func countBitValues(number uint64) (uint64, uint64) {
	oneBit := 0
	zeroBit := 0

	binary := strconv.FormatInt(int64(number), 2)

	for i := 0; i < len(binary); i++ {
		if string(binary[i]) == "1" {
			oneBit++
		}

		if string(binary[i]) == "0" {
			zeroBit++
		}
	}

	return uint64(oneBit), uint64(zeroBit)
}
