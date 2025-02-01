package prime

import (
	"fmt"
	"math"
	"strconv"
)

const (
	BitSize32 uint16 = 32
	BitSize64 uint16 = 64

	FNVPrime32Bit uint64 = 16777619
	FNVPrime64Bit uint64 = 1099511628211
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

/*
GetPrime - Utility function for returning a pre-computer FNV Prime number
based on the bitSize that you pass to it. If the unsigned 16-bit integer that is
passed does not match one of the constants defined above, then the 32-bit prime is returned
*/
func GetPrime(bitSize uint16) uint64 {
	if bitSize == BitSize32 {
		return FNVPrime32Bit
	}

	if bitSize == BitSize64 {
		return FNVPrime64Bit
	}

	return FNVPrime32Bit
}
