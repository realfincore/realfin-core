package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// PriceKeyPrefix is the prefix to retrieve all Price
	PriceKeyPrefix = "Price/value/"
)

// PriceKey returns the store key to retrieve a Price from the index fields
func PriceKey(
	symbol string,
) []byte {
	var key []byte

	symbolBytes := []byte(symbol)
	key = append(key, symbolBytes...)
	key = append(key, []byte("/")...)

	return key
}
