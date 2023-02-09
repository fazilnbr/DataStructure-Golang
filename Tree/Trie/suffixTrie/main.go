package main

import "fmt"

type TrieNode struct {
	Children map[rune]*TrieNode
}

type Trie struct {
	Root *TrieNode
}

var endsymbol = '*'

func main() {
	Trie := new(Trie)
	Trie.Root = &TrieNode{
		Children: map[rune]*TrieNode{},
	}

	Trie.insertToTrie("fazil")
	fmt.Println(Trie.Root.Children)

	fmt.Println(Trie.condains("azil"))

}

func (t *Trie) insertToTrie(str string) {
	t.populateSufixTrie(str)
}

func (t *Trie) populateSufixTrie(str string) {
	for i := 0; i < len(str); i++ {
		t.insertSubstringStartingAt(i, str)
	}
}

func (t *Trie) insertSubstringStartingAt(index int, str string) {
	tempNode := t.Root
	for i := index; i < len(str); i++ {
		charecter := str[i]

		if _, exist := tempNode.Children[rune(charecter)]; !exist {
			newTrieNode := &TrieNode{Children: make(map[rune]*TrieNode)}
			tempNode.Children[rune(charecter)] = newTrieNode
		}
		tempNode = tempNode.Children[rune(charecter)]

	}
	tempNode.Children[endsymbol] = nil

}

func (t *Trie) condains(str string) bool {
	tempNode := t.Root
	var charecter rune
	fmt.Println(tempNode)
	for i := 0; i < len(str); i++ {
		charecter = rune(str[i])

		_, exist := tempNode.Children[rune(charecter)]
		if !exist {
			return false
		}
		tempNode = tempNode.Children[rune(charecter)]
	}
	_, exist := tempNode.Children[rune(endsymbol)]
	return exist
}
