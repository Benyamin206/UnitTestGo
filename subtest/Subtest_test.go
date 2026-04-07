package subtest

import (
	"testing"

	"github.com/Benyamin206/UnitTestGo/v2/afterbeforetest"
	"github.com/stretchr/testify/assert"
)

func TestSubtest(t *testing.T) {
	t.Run("Sum", func(t *testing.T) {
		value := afterbeforetest.Sum(2, 2)
		assert.Equal(t, 4, value)
	})

	t.Run("Dot", func(t *testing.T) {
		value := afterbeforetest.Dot(2, 2)
		assert.Equal(t, 4, value)
	})

}
