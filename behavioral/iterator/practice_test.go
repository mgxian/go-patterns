package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getNumbers(n int) []interface{} {
	nums := make([]interface{}, n)
	for i := 0; i < n; i++ {
		nums[i] = i
	}
	return nums
}

func TestSimplePaginator(t *testing.T) {
	var p paginator = newSimplePaginator(getNumbers(50), 10)
	expected := []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.Equal(t, expected, p.first())

	expected = []interface{}{40, 41, 42, 43, 44, 45, 46, 47, 48, 49}
	assert.Equal(t, expected, p.last())

	expected = []interface{}{30, 31, 32, 33, 34, 35, 36, 37, 38, 39}
	assert.Equal(t, expected, p.prev())

	expected = []interface{}{20, 21, 22, 23, 24, 25, 26, 27, 28, 29}
	assert.Equal(t, expected, p.prev())

	expected = []interface{}{30, 31, 32, 33, 34, 35, 36, 37, 38, 39}
	assert.Equal(t, expected, p.next())

	p = newSimplePaginator(getNumbers(50), 8)
	expected = []interface{}{0, 1, 2, 3, 4, 5, 6, 7}
	assert.Equal(t, expected, p.first())

	expected = []interface{}{48, 49}
	assert.Equal(t, expected, p.last())

	expected = []interface{}{40, 41, 42, 43, 44, 45, 46, 47}
	assert.Equal(t, expected, p.prev())
}
