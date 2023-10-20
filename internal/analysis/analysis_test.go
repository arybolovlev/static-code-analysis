package analysis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewItemSet(t *testing.T) {
	want := &ItemSet{}
	have := NewItemSet()
	assert.Equal(t, have, want)
}
