package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertTolowerCaseWithoutSpace(t *testing.T) {
	t.Run("should convert string to lower case and replace string", func(t *testing.T) {
		expected := "a_capital_word"
		action := ConvertTolowerCaseWithoutSpace("A Capital Word")
		assert.Equal(t, expected, action)
	})
}
