package utils_test

import (
	"testing"

	"github.com/MaksimSvinin/gui-fresheye/internal/utils"
	"github.com/go-playground/assert/v2"
)

func TestSortReverse(t *testing.T) {
	arr := []int{4, 33, 20, 28, 36, 23, 11, 3, 8, 32, 9, 4, 21, 16, 15, 1, 31}

	utils.Rsshell2(arr)

	assert.Equal(t, arr, []int{36, 33, 32, 31, 28, 23, 21, 20, 16, 15, 11, 9, 8, 4, 4, 3, 1})
}
