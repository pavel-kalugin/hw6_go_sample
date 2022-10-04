package main

import (
	"testing"
	"github.com/stretchr/testify/require"
)


func TestFakeTest1(t *testing.T) {
	require.Equal(t, 2*2, 4, "2*2 == 4, Success")
}

func TestFakeTest2(t *testing.T) {
	require.Equal(t, 2*4, 8, "2*2 == 8, Success")
}
