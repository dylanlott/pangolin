package collection

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCollectionPath(t *testing.T) {
	path, err := getCollectionPath("test")
	fmt.Printf("path: %s", path)
	assert.NoError(t, err)
}
