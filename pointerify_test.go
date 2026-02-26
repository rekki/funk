package funk

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPointerify(t *testing.T) {
	t.Run("float32", func(t *testing.T) {
		p := Pointerify[float32](0.1)
		require.NotNil(t, p)
		assert.Equal(t, float32(0.1), *p)
	})

	t.Run("int", func(t *testing.T) {
		p := Pointerify(42)
		require.NotNil(t, p)
		assert.Equal(t, 42, *p)
	})

	t.Run("string", func(t *testing.T) {
		p := Pointerify("hello")
		require.NotNil(t, p)
		assert.Equal(t, "hello", *p)
	})

	t.Run("struct", func(t *testing.T) {
		type S struct{ X int }
		p := Pointerify(S{X: 1})
		require.NotNil(t, p)
		assert.Equal(t, 1, p.X)
	})

	t.Run("zero value", func(t *testing.T) {
		p := Pointerify[float32](0)
		require.NotNil(t, p)
		assert.Equal(t, float32(0), *p)
	})
}
