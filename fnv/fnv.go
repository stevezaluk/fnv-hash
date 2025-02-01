package fnv

import "github.com/stevezaluk/fnv-hash/prime"

const (
	OffsetStr        string = "chongo <Landon Curt Noll> /\\../\\"
	OffsetBasis64Bit uint64 = 14695981039346656037
	OffsetBasis32Bit uint64 = 2166136261
)

/*
GenerateOffsetBasis - Generates an offset basis for use in the FNV1 and FNV1a algorithms. The offset basis
is an FNV0 hash of the OffsetStr. Note: If you are generating a 32-bit offset basis, then you need to re-cast
the type to a 32-bit unsigned integer, or you will get an invalid offset basis. This can be done with the
following: uint32(GenerateOffsetBasis(prime.BitSize32))
*/
func GenerateOffsetBasis(bitSize uint16) uint64 {
	return NewFNV0Hash(bitSize, []byte(OffsetStr))
}

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

/*
NewFNV1Hash - Generate a new FNV1 hash
*/
func NewFNV1Hash(bitSize uint16, data []byte) uint64 {
	var hash uint64 = OffsetBasis64Bit
	if bitSize == prime.BitSize32 {
		hash = OffsetBasis32Bit
	}

	for _, _byte := range data {
		hash = hash * prime.GetPrime(bitSize)
		hash = hash ^ uint64(_byte)
	}

	return hash
}

/*
NewFNV1AHash - Generate a new FNV1A hash
*/
func NewFNV1AHash(bitSize uint16, data []byte) uint64 {
	var hash uint64 = OffsetBasis64Bit
	if bitSize == prime.BitSize32 {
		hash = OffsetBasis32Bit
	}

	for _, _byte := range data {
		hash = hash ^ uint64(_byte)
		hash = hash * prime.GetPrime(bitSize)
	}

	return hash
}
