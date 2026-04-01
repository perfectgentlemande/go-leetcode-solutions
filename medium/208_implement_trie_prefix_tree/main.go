package main

const AlphabetSize = 26

type TrieNode struct {
	children [AlphabetSize]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{
		root: &TrieNode{},
	}
}

func (this *Trie) Insert(word string) {
	n := len(word)
	currentNode := this.root
	for i := 0; i < n; i++ {
		charIndex := word[i] - 'a'

		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &TrieNode{}
		}

		currentNode = currentNode.children[charIndex]
	}

	currentNode.isEnd = true
}

func (this *Trie) Search(word string) bool {
	lastNode := this.searchWord(word)
	if lastNode == nil {
		return false
	}

	return lastNode.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.searchWord(prefix) != nil
}

func (this *Trie) searchWord(word string) *TrieNode {
	n := len(word)
	currentNode := this.root
	for i := 0; i < n; i++ {
		charIndex := word[i] - 'a'

		if currentNode.children[charIndex] == nil {
			return nil
		}

		currentNode = currentNode.children[charIndex]
	}

	return currentNode
}

func main() {

}
