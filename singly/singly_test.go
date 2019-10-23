package singly_test

import (
	"linkedlist/singly"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSinglyLinkedList(t *testing.T) {
	id := 1
	singly := singly.NewSinglyList(1)
	require.Equal(t, id, singly.ID)

	err := singly.Add("item #1")
	require.NoError(t, err)
	err = singly.Add("item #2")
	require.NoError(t, err)
	err = singly.Add("item #3")
	require.NoError(t, err)

	currentItem, err := singly.Current()
	require.Error(t, err)

	err = singly.Beginning()
	require.NoError(t, err)

	currentItem, err = singly.Current()
	require.NoError(t, err)
	require.Equal(t, currentItem, "item #1")

	currentItem, err = singly.Next()
	require.NoError(t, err)
	require.Equal(t, currentItem, "item #2")
}
