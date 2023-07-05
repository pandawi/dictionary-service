package trie_test

import (
	"testing"

	tr "dictionary-service/app/lib/trie"

	"github.com/stretchr/testify/assert"
)

func TestTrie_InsertAndSearch(t *testing.T) {
	words := []string{"apple", "banana", "car", "dog", "cat"}
	trie := tr.NewTrie().FromSource(words)
	for _, word := range words {
		assert.True(t, trie.Search(word))
	}
	nonExistingWords := []string{"ape", "ball", "care", "dogs", "category"}
	for _, word := range nonExistingWords {
		assert.False(t, trie.Search(word))
	}
}
