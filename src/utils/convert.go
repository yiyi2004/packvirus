package utils

import "strconv"

// Convert2BinaryString -
func Convert2BinaryString(num int) string {
	s := ""

	if num == 0 {
		return "0000"
	}

	for ; num > 0; num /= 2 {
		lsb := num % 2
		s = strconv.Itoa(lsb) + s
	}

	for len(s) < 4 {
		s = "0" + s
	}

	return s
}
