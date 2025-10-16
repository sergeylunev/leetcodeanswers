// zigzagconversion_test.go
package zigzagconversion

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		name        string
		inputS      string
		numRows     int
		want        string
		expectPanic bool
	}{
		{
			name:    "Empty string",
			inputS:  "",
			numRows: 3,
			want:    "",
		},
		{
			name:    "Single character",
			inputS:  "A",
			numRows: 1,
			want:    "A",
		},
		{
			name:    "NumRows 1",
			inputS:  "PAYPALISHIRING",
			numRows: 1,
			want:    "PAYPALISHIRING",
		},
		{
			name:    "NumRows greater than string length",
			inputS:  "ABC",
			numRows: 5,
			want:    "ABC",
		},
		{
			name:    "Standard zigzag cjsonversion",
			inputS:  "PAYPALISHIRING",
			numRows: 3,
			want:    "PAHNAPLSIIGYIR",
		},
		{
			name:    "Four row conversion",
			inputS:  "PAYPALISHIRING",
			numRows: 4,
			want:    "PINALSIGYAHRPI",
		},
		{
			name:        "Zero rows panic",
			inputS:      "TEST",
			numRows:     0,
			expectPanic: true,
		},
		{
			name:        "Negative rows panic",
			inputS:      "TEST",
			numRows:     -2,
			expectPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectPanic {
				require.Panics(t, func() {
					convert(tt.inputS, tt.numRows)
				})
				return
			}

			got := convert(tt.inputS, tt.numRows)
			assert.Equal(t, tt.want, got)
		})
	}
}
