# Overview
This service efficiently searches for all words matching the permutations of the given word using a Trie data structure.

# Algorithm:

* The solution uses a Trie data structure to efficiently search for words matching the permutations of a given word.
* The Trie is constructed by inserting all words from the word list.
* The algorithm then performs a depth-first search to find permutations of the given word that exist in the Trie.
Backtracking is used to explore all possible permutations.

### Reasons for using this algorithm:

The Trie data structure provides efficient storage and lookup for words.
By storing the words in a Trie, we can leverage prefix-based searching, which reduces the number of comparisons needed.
The depth-first search with backtracking allows us to explore all possible permutations without redundancy.

# Runtime Complexity:

* Building the Trie: O(n * k), where n is the number of words in the list and k is the average length of the words.
* Searching for permutations: O(m), where m is the length of the given word.
Since the length of the given word is typically small compared to the word list size, the search operation's complexity is considered linear.

# Running the service
Running the service should be done with Docker.

* Clone or Download the project
* Run `docker build -t dict-service .` in the root directory of the project in order to build the image
* Run `docker run -p 8000:8000 dict-service`

# Example calls
`curl -X GET http://localhost:8000/api/v1/stats`

`curl -X GET http://localhost:8000/api/v1/similar?word=apple`
