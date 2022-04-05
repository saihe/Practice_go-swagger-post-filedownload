package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFileService(t *testing.T) {
	assert.NotNil(t, NewFileService(), "Fail initialize.")
}

func TestFileServiceImpl_Generate(t *testing.T) {
	n := fmt.Sprintf("%s_%s.csv", "0001", "test")
	fn, f := NewFileService().Generate("0001", "test")
	assert.NotNil(t, f, "Fail generate.")
	assert.NotEqualValues(t, n, fn, "Not equal names")
}
