package doubly_test

import (
	"linkedlist/doubly"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDoublyLinkedList(t *testing.T) {
	id := 1
	doubly := doubly.NewDoublyList(1)
	require.Equal(t, id, doubly.ID)

	err := doubly.Add("item #1")
	require.NoError(t, err)
	err = doubly.Add("item #2")
	require.NoError(t, err)
	err = doubly.Add("item #3")
	require.NoError(t, err)

	err = doubly.Beginning()
	require.NoError(t, err)

	currentItem, err := doubly.Current()
	require.NoError(t, err)
	require.Equal(t, currentItem, "item #1")

	currentItem, err = doubly.Next()
	require.NoError(t, err)
	require.Equal(t, currentItem, "item #2")

	currentItem, err = doubly.Previous()
	require.NoError(t, err)
	require.Equal(t, currentItem, "item #1")
}
