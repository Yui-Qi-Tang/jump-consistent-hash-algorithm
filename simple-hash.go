// Package jchash implements jump consistent hash in go
// HINT: This is a toy
package jchash

// Hash returns hash value; it is the implementation of jump consistent hash
func Hash(key uint64, numBuckets int32) int32 {

	if numBuckets <= 0 {
		return 0
	}

	var b, j int64 = -1, 0

	for j < int64(numBuckets) {
		b = j
		key = key*2862933555777941757 + 1
		j = int64(float64(b+1) * (float64(int64(1)<<31) / float64((key>>33)+1)))
	}
	return int32(b)
}
