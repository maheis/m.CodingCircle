package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHottest(t *testing.T) {
	v := []int{27, 9, 17, 2, 12, 8}
	k := 3
	result := []int{27, 17, 17, 12}

	resI, err := iterative(v, k)
	assert.NoError(t, err)
	assert.Equal(t, resI, result)

	resQ, err := que(v, k)
	assert.NoError(t, err)
	assert.Equal(t, resQ, result)

	resS, err := slice(v, k)
	assert.NoError(t, err)
	assert.Equal(t, resS, result)
}
