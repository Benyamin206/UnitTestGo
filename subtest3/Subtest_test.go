package subtest3

import (
	"testing"

	"github.com/Benyamin206/UnitTestGo/v2/afterbeforetest"
	"github.com/stretchr/testify/assert"
)

func TestSubtest(t *testing.T) {
	var tests []struct {
		name     string
		request1 int
		request2 int
		expected int
	} = []struct {
		name     string
		request1 int
		request2 int
		expected int
	}{
		{
			name:     "Sum1",
			request1: 2,
			request2: 3,
			expected: 5,
		},
		{
			name:     "Sum2",
			request1: 3,
			request2: 5,
			expected: 8,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			value := afterbeforetest.Sum(test.request1, test.request2)
			assert.Equal(t, test.expected, value, "The result must be %d", test.expected)
		})
	}

}
