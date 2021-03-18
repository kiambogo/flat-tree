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

// Parent returns the parent node of the provided node
func Parent(n uint64) uint64 {
	return Index(Depth(n)+1, Offset(n)/2)
}

// Sibling returns the sibling of the provided node
func Sibling(n uint64) uint64 {
	return Index(Depth(n), Offset(n)^1)
}

// Uncle returns the parent's sibling of the provided node
func Uncle(n uint64) uint64 {
	return Index(Depth(n)+1, Offset(Parent(n))^1)
}

// Children returns the children of the provided node, if it exists
// Returns the children and a bool indicating if they exist
func Children(n uint64) (left uint64, right uint64, exists bool) {
	if isEven(n) {
		return 0, 0, false
	}

	depth := Depth(n)
	offset := Offset(n) * 2

	return Index(depth-1, offset), Index(depth-1, offset+1), true
}

// LeftChild returns the left child of the provided node, if it exists
// Returns the left child and a bool indicating if it exists
func LeftChild(n uint64) (uint64, bool) {
	if isEven(n) {
		return 0, false
	}

	return Index(Depth(n)-1, Offset(n)*2), true
}

// RightChild returns the left child of the provided node, if it exists
// Returns the right child and a bool indicating if it exists
func RightChild(n uint64) (uint64, bool) {
	if isEven(n) {
		return 0, false
	}

	return Index(Depth(n)-1, (Offset(n)*2)+1), true
}

// Spans returns the left and right most nodes in the tree which the provided node spans
func Spans(n uint64) (left uint64, right uint64) {
	return LeftSpan(n), RightSpan(n)
}

// LeftSpan returns the left most node in the tree which the provided node spans
func LeftSpan(n uint64) uint64 {
	depth := Depth(n)
	if depth == 0 {
		return n
	}
	return Offset(n) * (2 << depth)
}

// RightSpan returns the right most node in the tree which the provided node spans
func RightSpan(n uint64) uint64 {
	depth := Depth(n)
	if depth == 0 {
		return n
	}
	return (Offset(n)+1)*(2<<depth) - 2
}

func isEven(n uint64) bool {
	return n%2 == 0
}
