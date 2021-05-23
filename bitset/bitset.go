package bitset

import "math/bits"

// Set sets the bit at index to 1.
func Set(bitset *uint64, index int) {
	*bitset |= 1 << index
}

// Clear sets the bit at index to 0.
func Clear(bitset *uint64, index int) {
	*bitset &= ^(1 << index)
}

// EachSetBit calls the callback for each set bit in the bitset.
func EachSetBit(bitset uint64, callback func(i int)) {
	for i := bits.TrailingZeros64(bitset); bitset != 0; i = bits.TrailingZeros64(bitset) {
		bitset &= ^(1 << i)
		callback(i)
	}
}

// Count returns the amount of set bits in the bitset.
func Count(bitset uint64) int {
	return bits.OnesCount64(bitset)
}

// IsSet returns true if the bit at index is set.
func IsSet(bitset uint64, index int) bool {
	return (bitset & (1 << index)) > 0
}
