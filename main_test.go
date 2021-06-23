package main

import (
	"testing"
	// "github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {

  // assert equality
	assert.Equal(t, 123, 124, "they should be equal")
}


// func TestEcho(t *testing.T) {
// 	// Test happy path
// 	err := echo([]string{"bin-name", "hello", "world!"})
// 	require.NoError(t, err)
// }

// func TestEchoErrorNoArgs(t *testing.T) {
// 	// Test empty arguments
// 	err := echo([]string{})
// 	require.Error(t, err)
// }

// func assertFileExists(t *testing.T) {
// 	_, err := os.Lstat(path)
// 	if err != nil {
// 		if os.IsNotExist(err) {
// 			assert.Fail(t, "unable to find file %q", path)
// 		}
// 		assert.Fail(t, "error when running os.Lstat(%q): %s", path, err)
// 	}
// }