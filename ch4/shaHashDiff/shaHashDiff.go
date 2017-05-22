package shaHashDiff

import "crypto/sha256"

func generateShahash(str string) [32]byte {
	return sha256.Sum256([]byte(str))
}
