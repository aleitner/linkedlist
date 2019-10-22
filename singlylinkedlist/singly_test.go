package singlylinkedlist_test

import (
	"ll/singlylinkedlist"
	"testing"

	"github.com/stretchr/testify/require"
)

func SinglyLinkedListTest(t *testing.T) {
	id := 1
	singly := singlylinkedlist.New(1)
	require.Equal(t, id, singly.ID)

	err := singly.Add("item #1")
	require.NoError(t, err)
	err = singly.Add("item #2")
	require.NoError(t, err)
	err = singly.Add("item #3")
	require.NoError(t, err)

	currentItem, err := singly.Current()
	require.NoError(t, err)
	require.Equal(t, currentItem, "item #1")

	currentItem, err = singly.Next()
	require.NoError(t, err)
	require.Equal(t, currentItem, "item #2")
}
