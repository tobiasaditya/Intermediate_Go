package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrint(t *testing.T) {
	result := helloWorld()

	// Vanilla
	// if result != "hello world 1" {
	// 	t.Fail()
	// }

	// Using external package
	require.Equal(t, "hello world", result)
}
