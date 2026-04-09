package subtest2

import (
	"testing"

	"github.com/Benyamin206/UnitTestGo/v2/afterbeforetest"
	"github.com/Benyamin206/UnitTestGo/v2/app2"
	"github.com/stretchr/testify/assert"
)

func TestSubtest(t *testing.T) {
	t.Run("A", func(t *testing.T) {
		value := app2.Sum(2, 2)
		assert.Equal(t, 5, value, "Must be equal")
	})

	t.Run("B", func(t *testing.T) {
		value := afterbeforetest.Dot(2, 2)
		assert.Equal(t, 4, value)
	})
}
