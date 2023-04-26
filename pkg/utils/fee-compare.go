package utils

import (
	"encoding/binary"
	"encoding/hex"
)

// reference ironfish/src/primitives/rawTransation.ts>deserialize
func GetFee(transaction string) (uint64, error) {
	hexStr, err := hex.DecodeString(transaction)
	if err != nil {
		return 0, err
	}

	fee := binary.LittleEndian.Uint64(hexStr)
	return fee, nil
}
