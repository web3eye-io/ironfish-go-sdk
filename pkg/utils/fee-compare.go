package utils

import (
	"encoding/binary"
	"encoding/hex"
)

func OverFeeLimit(maxFee uint64, transaction string) (bool, error) {
	hexStr, err := hex.DecodeString(transaction)
	if err != nil {
		return false, err
	}

	fee := binary.LittleEndian.Uint64(hexStr)
	if fee > maxFee {
		return true, nil
	}
	return false, nil
}
