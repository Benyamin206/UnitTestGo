package app2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	value := Sum(1, 2)
	assert.Equal(t, 3, value, "Salah")
	fmt.Println("Testing Sum 1 selesai")
}
