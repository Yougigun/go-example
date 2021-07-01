package main

import (
	"encoding/hex"
	"math/big"
)

func main() {
	
}
func getHashInt(hashBytes []byte) uint64 {
	hexSha1 := hex.EncodeToString(hashBytes[:])
	// Integer base16 conversion
	intBase16, success := new(big.Int).SetString(hexSha1, 16)
	if success {
		return intBase16.Uint64()
	}
	return 0
}