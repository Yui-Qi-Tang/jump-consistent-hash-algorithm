package jchash

import (
	"testing"
)

func TestHash(t *testing.T) {
	t.Log("start testing simple jump consistent hashing...")
	testcases := []struct {
		key        uint64
		numBuckets int32
		want       int32
	}{
		{1, 1, 0},
		{2, 1, 0},
		{3, 1, 0},
		{1, 2, 0},
		{2, 2, 0},
		{10000, 2, 1},
		{123, 123, 106},
		{100, 1000, 169},
		{0, -10, 0},
	}

	for _, tt := range testcases {
		result := Hash(tt.key, tt.numBuckets)
		if tt.want != result {
			t.Log("it sould be:", tt.want, "but got:", result)
		}
	}

	t.Log("... Passed")
}
