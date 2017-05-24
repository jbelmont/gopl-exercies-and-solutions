package shaHashDiff

import (
	"crypto/sha256"
	"fmt"
)

func generateShahash(str string) string {
	return fmt.Sprintf("%X", sha256.Sum256([]byte(str)))
}

func compareBits(str1, str2 string) int {
	var count int
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			count++
		}
	}
	return count
}
