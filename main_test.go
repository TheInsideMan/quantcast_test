package main

import (
	"testing"
	"fmt"
	"os"
	// "github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"
)

func TestFileExists(t *testing.T) {
	fmt.Println("running AssertFileExists")
	path := "cookie_log.csv"
	if _, err := os.Stat(path); err != nil {
	  assert.Fail(t, fmt.Sprintf("unable to find file with path %v", path))
	}
}