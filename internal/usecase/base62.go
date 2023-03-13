package usecase

import (
	"errors"
	"math"
	"strings"
)

const (
	alphanumeric = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	size         = uint64(len(alphanumeric))
)

func Encode(id uint64) string {
	var encoded strings.Builder
	encoded.Grow(7)

	for ; id > 0; id = id / size {
		encoded.WriteByte(alphanumeric[id%size])
	}
	return encoded.String()
}

func Decode(encoded string) (uint64, error) {
	var id uint64
	for i := 0; i < len(encoded); i++ {
		index := findIndex(encoded[i])
		if index == -1 {
			return 0, errors.New("Invalid character: " + string(encoded[i]))
		}
		id += uint64(index) * uint64(math.Pow(float64(size), float64(i)))
	}
	return id, nil
}

// The order of A-Z, a-z and 0-9 are contiguous
func findIndex(char byte) int {
	if char >= 'A' && char <= 'Z' {
		return int(char - 'A')
	} else if char >= 'a' && char <= 'z' {
		return int(char-'a') + 26
	} else if char >= '0' && char <= '9' {
		return int(char-'0') + 52
	} else {
		return -1
	}
}
