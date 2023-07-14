package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	expected := 5
	res := Add(2, 3)

	assert.Equal(t, expected, res)
}
