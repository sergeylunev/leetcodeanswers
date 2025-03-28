package longestpalindromicsubstring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongestPalindromicSubstring(t *testing.T) {
	t.Run("simple success", func(t *testing.T) {
		actual := longestPalindrome("babad")

		require.Equal(t, "bab", actual)
	})

	t.Run("success one letter", func(t *testing.T) {
		actual := longestPalindrome("a")

		require.Equal(t, "a", actual)
	})
	t.Run("success two letters", func(t *testing.T) {
		actual := longestPalindrome("aa")

		require.Equal(t, "aa", actual)
	})
	t.Run("simplest success", func(t *testing.T) {
		actual := longestPalindrome("aba")

		require.Equal(t, "aba", actual)
	})
	t.Run("fail with two letters", func(t *testing.T) {
		actual := longestPalindrome("ab")

		require.Equal(t, "", actual)
	})
}
