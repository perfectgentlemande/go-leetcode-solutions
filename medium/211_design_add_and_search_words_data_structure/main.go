package main

const AlphabetSize = 26

type WordDictionary struct {
	children [AlphabetSize]*WordDictionary
	isEnd    bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	n := len(word)
	currentNode := this
	for i := 0; i < n; i++ {
		charIndex := word[i] - 'a'

		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &WordDictionary{}
		}

		currentNode = currentNode.children[charIndex]
	}

	currentNode.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	currentNode := this

	for i := 0; i < len(word); i++ {
		if word[i] == '.' {
			for j := 0; j < 26; j++ {
				if currentNode.children[j] != nil && currentNode.children[j].Search(word[i+1:]) {
					return true
				}
			}

			return false
		}

		c := int(word[i] - 'a')
		if currentNode.children[c] == nil {
			return false
		}

		currentNode = currentNode.children[c]
	}

	return currentNode.isEnd
}
