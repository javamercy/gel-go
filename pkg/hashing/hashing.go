package hashing

import "crypto/sha1"

func ComputeSha1Hash(data []byte) []byte {
	sum := sha1.Sum(data)
	return sum[:]
}
