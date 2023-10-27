package analysis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewItemSet(t *testing.T) {
	want := &NodesSet{}
	have := NewNodeSet()

	assert.Equal(t, have, want)
}

func TestInsertItem(t *testing.T) {
	want := &NodesSet{
		Nodes: []*Node{
			{
				Name: "this",
			},
		},
	}
	have := NewNodeSet()
	have.InsertNode(Node{
		Name: "this",
	})

	assert.Equal(t, have, want)
}

func TestItem(t *testing.T) {
	i := &NodesSet{
		Nodes: []*Node{
			{
				Name: "this",
			},
		},
	}

	want := &Node{
		Name: "this",
	}
	have := i.GetNode("this")

	assert.Equal(t, have, want)
}
