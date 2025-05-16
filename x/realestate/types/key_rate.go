package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// RateKeyPrefix is the prefix to retrieve all Rate
	RateKeyPrefix = "Rate/value/"
)

// RateKey returns the store key to retrieve a Rate from the index fields
func RateKey(
	symbol string,
) []byte {
	var key []byte

	symbolBytes := []byte(symbol)
	key = append(key, symbolBytes...)
	key = append(key, []byte("/")...)

	return key
}
