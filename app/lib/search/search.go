package search

import (
	t "dictionary-service/app/lib/trie"
)

// searchWordPermutations searches for words in the word list that match the permutations of the given word.
func SearchWordPermutations(word string, trie t.Trie) []string {
	matchedWords := []string{}

	// Count the frequency of characters in the given word.
	wordFreq := make(map[rune]int)
	for _, char := range word {
		wordFreq[char]++
	}

	// Perform a depth-first search with backtracking to find permutations of the given word that exist in the Trie.
	searchPermutations(trie, word, wordFreq, []rune{}, &matchedWords)
	return matchedWords
}

// searchPermutations performs a depth-first search with backtracking to find permutations of the given word that exist in the Trie.
func searchPermutations(
	trie t.Trie,
	word string,
	wordFreq map[rune]int,
	currentWord []rune,
	matchedWords *[]string,
) {
	// If the current permutation matches a word in the Trie and its length is the same as the given word, add it to the matchedWords list.
	if trie.Search(string(currentWord)) && len(currentWord) == len(word) {
		*matchedWords = append(*matchedWords, string(currentWord))
	}

	// Explore each character and its count in the word frequency map.
	for char, count := range wordFreq {
		if count > 0 {
			// Decrease the count of the character in the word frequency map to mark its usage in the current permutation.
			wordFreq[char]--

			// Add the current character to the currentWord slice.
			currentWord = append(currentWord, char)

			// Recursively search for permutations using the updated currentWord and wordFreq.
			searchPermutations(trie, word, wordFreq, currentWord, matchedWords)

			// Backtrack by removing the last character from the currentWord slice and restoring its count in the word frequency map.
			currentWord = currentWord[:len(currentWord)-1]
			wordFreq[char]++
		}
	}
}
