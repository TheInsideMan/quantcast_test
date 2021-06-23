package main

import (
	"testing"
	"fmt"
	"os"
	// "github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	fmt.Println("running TestSomething")
  // assert equality
	assert.Equal(t, 123, 123, "they should be equal")
}

func assertFileExists(t *testing.T) {
	path := "cookie_log2.csv"
	_, err := os.Lstat(path)
	if err != nil {
		if os.IsNotExist(err) {
			assert.Fail(t, "unable to find file %q", path)
		}
		assert.Fail(t, "error when running os.Lstat(%q): %s", path, err)
	}
}