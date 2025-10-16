package zigzagconversion

import "strings"

func convert(s string, numRows int) string {
	slices := make([][]byte, numRows)

	if numRows == 1 {
		return s
	}

	idx := 0
	direction := "forward"
	for i := 0; i < len(s); i++ {
		switch idx {
		case numRows:
			idx = idx - 2
			direction = "backward"
		case -1:
			idx = idx + 2
			direction = "forward"
		}

		slices[idx] = append(slices[idx], s[i])

		if direction == "forward" {
			idx = idx + 1
		} else {
			idx = idx - 1
		}

	}
	var builder strings.Builder

	for _, s := range slices {
		for _, b := range s {
			builder.WriteByte(b)
		}
	}

	return builder.String()
}
