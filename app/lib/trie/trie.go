package trie

type Trie interface {
	Insert(word string)
	Search(word string) bool
	FromSource(words []string) Trie
}

type TrieNode struct {
	children map[rune]*TrieNode
	isWord   bool
}

type TrieImpl struct {
	root *TrieNode
}

func NewTrie() Trie {
	return &TrieImpl{
		root: &TrieNode{children: make(map[rune]*TrieNode)},
	}
}

func (t *TrieImpl) Insert(word string) {
	node := t.root
	for _, char := range word {
		if _, ok := node.children[char]; !ok {
			node.children[char] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[char]
	}
	node.isWord = true
}

func (t *TrieImpl) Search(word string) bool {
	node := t.root
	for _, char := range word {
		if _, ok := node.children[char]; !ok {
			return false
		}
		node = node.children[char]
	}
	return node.isWord
}

func (t *TrieImpl) FromSource(words []string) Trie {
	for _, word := range words {
		t.Insert(word)
	}
	return t
}
