package flattree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Index(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		depth, offset, expected uint64
	}{
		{0, 0, 0},
		{0, 1, 2},
		{0, 2, 4},
		{0, 3, 6},
		{0, 4, 8},
		{1, 0, 1},
		{1, 1, 5},
		{1, 2, 9},
		{1, 3, 13},
		{1, 4, 17},
		{2, 0, 3},
		{2, 1, 11},
		{2, 2, 19},
		{2, 3, 27},
		{2, 4, 35},
		{3, 0, 7},
		{3, 1, 23},
		{3, 2, 39},
		{3, 3, 55},
		{3, 4, 71},
	}

	for _, tc := range testCases {
		i := Index(tc.depth, tc.offset)
		assert.Equal(t, tc.expected, i, "Index of %d with offset %d expected value of %d; got %d", tc.depth, tc.offset, tc.expected, i)
	}
}

func Test_Depth(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		depth, expected uint64
	}{
		{0, 0},
		{1, 1},
		{2, 0},
		{3, 2},
		{4, 0},
		{5, 1},
		{6, 0},
		{7, 3},
		{8, 0},
	}

	for _, tc := range testCases {
		d := Depth(tc.depth)
		assert.Equal(t, tc.expected, d, "Depth of %d with expected value of %d; got %d", tc.depth, tc.expected, d)
	}
}

func Test_Offset(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		offset, expected uint64
	}{
		{0, 0},
		{1, 0},
		{2, 1},
		{3, 0},
		{4, 2},
		{5, 1},
		{6, 3},
		{7, 0},
		{8, 4},
		{9, 2},
		{10, 5},
		{11, 1},
		{12, 6},
	}

	for _, tc := range testCases {
		o := Offset(tc.offset)
		assert.Equal(t, tc.expected, o, "Offset of %d with expected value of %d; got %d", tc.offset, tc.expected, o)
	}
}
