package afterbeforetest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	value := Sum(2, 2)
	assert.Equal(t, 4, value, "Must be equal")
}
