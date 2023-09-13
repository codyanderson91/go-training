package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(Add(1, 2), 3, "Values should be the same")
	assert.NotEqual(Add(1, 3), 3, "Values should not be the same")
}
