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

func TestInsertItem(t *testing.T) {
	want := &ItemSet{
		Items: []*Item{
			{
				Name: "this",
			},
		},
	}
	have := NewItemSet()
	have.InsertItem(Item{
		Name: "this",
	})

	assert.Equal(t, have, want)
}

func TestItem(t *testing.T) {
	i := &ItemSet{
		Items: []*Item{
			{
				Name: "this",
			},
		},
	}

	want := &Item{
		Name: "this",
	}
	have := i.Item("this")

	assert.Equal(t, have, want)
}
