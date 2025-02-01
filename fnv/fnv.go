package fnv

import "github.com/stevezaluk/fnv-hash/prime"

/*
NewFNV0Hash - Generates a new FNV0 Hash based on the bitSize you pass to it.
This is deprecated and should not be used unless you are computing an FNV offset basis
*/
func NewFNV0Hash(bitSize uint16, data []byte) uint64 {
	var hash uint64

	for _, _byte := range data {
		hash = hash * prime.GetPrime(bitSize)
		hash = hash ^ uint64(_byte)
	}

	return hash
}
