package raw

import (
	"bytes"
	"strconv"
)

// BytesToUint function will convert the input byte array
// into a uint64; returning a uint64 and an error
func BytesToUint(b []byte) (uint64, error) {
	t, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		return 0, err
	}
	return t, nil
}

// BytesToFloat function will convert the input byte array
// into a float64; returning a float64 and an error 
func BytesToFloat(b []byte) (float64, error) {
	t, err := strconv.ParseFloat(string(b), 64)
	if err != nil {
		return 0, err
	}
	return t, nil
}


// ByteJoin function will join variable amounts of byte slices
// to make it easier (and faster) to build content
func ByteJoin(input ...[]byte) []byte {
	var empty []byte
	array := bytes.Join(input, empty)
	return array
}
