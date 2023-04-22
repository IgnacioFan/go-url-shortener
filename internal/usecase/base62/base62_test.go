package base62

import (
	"errors"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestBase62Encode(t *testing.T) {
	tests := []struct {
		name        string
		input       uint64
		expectedRes string
	}{
		{
			"Return encoded string",
			10000,
			"SlC",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedRes, Encode(test.input))
		})
	}
}

func TestBase62Decode(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectedRes uint64
		expectedErr error
	}{
		{
			"Return decoded ID",
			"SlC",
			10000,
			nil,
		},
		{
			"Return invalid ID",
			"[",
			0,
			errors.New("Invalid character: ["),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			id, err := Decode(test.input)
			if test.expectedErr != nil {
				assert.Equal(t, test.expectedRes, id)
				assert.Equal(t, test.expectedErr, err)
			} else {
				assert.Equal(t, test.expectedRes, id)
				assert.Equal(t, test.expectedErr, err)
			}
		})
	}
}
