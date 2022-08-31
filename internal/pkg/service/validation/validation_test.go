package validation

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsTitleOk(t *testing.T) {
	t.Parallel()
	t.Run("normal", func(t *testing.T) {
		str := strings.Repeat("a", maxNameLen-1)
		assert.NoError(t, IsTitleOk(str))
	})

	t.Run("large", func(t *testing.T) {
		str := strings.Repeat("a", maxNameLen+1)
		assert.Error(t, IsTitleOk(str))
	})

	t.Run("empty", func(t *testing.T) {
		assert.Error(t, IsTitleOk(""))
	})
}

func TestDescriptionOk(t *testing.T) {
	t.Parallel()
	t.Run("normal", func(t *testing.T) {
		str := strings.Repeat("a", maxDescriptionLen-1)
		assert.NoError(t, IsDescriptionOk(str))
	})

	t.Run("large", func(t *testing.T) {
		str := strings.Repeat("a", maxDescriptionLen+1)
		assert.Error(t, IsDescriptionOk(str))
	})

	t.Run("empty", func(t *testing.T) {
		assert.NoError(t, IsDescriptionOk(""))
	})
}
