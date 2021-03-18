package flattree

import "math/bits"

// Index returns the flat-tree of the node at the provided depth and offset
func Index(depth, offset uint64) uint64 {
	return (offset << (depth + 1)) | ((1 << depth) - 1)
}

// Depth returns the depth of a given node
func Depth(n uint64) uint64 {
	return uint64(bits.TrailingZeros64(^n))
}

// Offset returns the offset of a given node
// The offset is the distance from the left edge of the tree
func Offset(n uint64) uint64 {
	if isEven(n) {
		return n / 2
	}
	return n >> (Depth(n) + 1)
}

func isEven(n uint64) bool {
	return n%2 == 0
}
