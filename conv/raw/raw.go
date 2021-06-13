package raw

import "strconv"

func bytesToUint(b []byte) (uint64, error) {
	t, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		return 0, err
	}
	return t, nil
}


func bytesToFloat(b []byte) (float64, error) {
	t, err := strconv.ParseFloat(string(b), 64)
	if err != nil {
		return 0, err
	}
	return t, nil
}