package main

import "fmt"

type TrieNode struct {
	children map[byte]*TrieNode
	isWord   bool
}

type Trie struct {
	root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	root := &TrieNode{
		children: map[byte]*TrieNode{},
	}
	return Trie{root}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if nextNode, ok := node.children[c]; ok {
			node = nextNode
		} else {
			node.children[c] = &TrieNode{
				children: map[byte]*TrieNode{},
			}
			node = node.children[c]
		}
	}

	node.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.root
	for i := 0; i < len(word); i++ {
		if nextNode, ok := node.children[word[i]]; ok {
			node = nextNode
		} else {
			return false
		}
	}

	return node.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for i := 0; i < len(prefix); i++ {
		if nextNode, ok := node.children[prefix[i]]; ok {
			node = nextNode
		} else {
			return false
		}
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
	trie := Constructor()
	words := []string{"sam", "john", "tim", "jose", "rose",
		"cat", "dog", "dogg", "roses"}
	for i := 0; i < len(words); i++ {
		trie.Insert(words[i])
	}
	wordsToFind := []string{"sam", "john", "tim", "jose", "rose",
		"cat", "dog", "dogg", "roses", "rosess", "ans", "san"}
	for i := 0; i < len(wordsToFind); i++ {
		found := trie.Search(wordsToFind[i])
		if found {
			fmt.Printf("Word \"%s\" found in trie\n", wordsToFind[i])
		} else {
			fmt.Printf("Word \"%s\" not found in trie\n", wordsToFind[i])
		}
	}
}
