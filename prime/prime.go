package prime

import (
	"fmt"
	"math"
	"strconv"
)

const (
	PrimeBit32   uint16 = 32
	PrimeBit64   uint16 = 64
	PrimeBit128  uint16 = 128
	PrimeBit256  uint16 = 256
	PrimeBit512  uint16 = 512
	PrimeBit1024 uint16 = 1024
)

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

/*
isValid - Return's true if the prime number is valid, and false if it is not
*/
func isValid(prime uint64) bool {
	left := prime % (uint64(math.Pow(2, 40) - math.Pow(2, 24) - 1))
	right := uint64(math.Pow(2, 24) + math.Pow(2, 8) + 7)
	if left > right {
		return true
	}

	return false
}

/*
GenerateFNVPrime - Generate the first FNV prime number that meets the constraints outlined
in this article: https://datatracker.ietf.org/doc/html/draft-eastlake-fnv-17.html#section-2.1
*/
func GenerateFNVPrime(bitSize uint16) uint64 {
	t := math.Floor(float64((5 + bitSize) / 12))

	var prime float64
	for b := 1; b < 255; b++ {
		oneBits := countBitValues(uint64(b))
		if oneBits == 4 || oneBits == 5 {
			prime = math.Pow(256.00, t) + math.Pow(2.00, 8.00) + float64(b)
			fmt.Println(prime, b)
			if isValid(uint64(prime)) {
				return uint64(prime)
			}
		}
	}

	return uint64(0)
}
