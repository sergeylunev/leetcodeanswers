package palindromenumber

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	t.Run("simple true", func(t *testing.T) {
		actual := isPalindrome(121)

		require.Equal(t, true, actual)
	})

	t.Run("simple false", func(t *testing.T) {
		actual := isPalindrome(10)

		require.Equal(t, false, actual)
	})

	t.Run("simple false with negative", func(t *testing.T) {
		actual := isPalindrome(-121)

		require.Equal(t, false, actual)
	})
}
