package afterbeforetest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDot(t *testing.T) {
	value := Dot(5, 5)
	assert.Equal(t, 45, value, "Must be equal")
}
