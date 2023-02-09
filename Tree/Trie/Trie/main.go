package main

type TrieNode struct {
	Children map[rune]*TrieNode
}
type Trie struct {
	Root *TrieNode
}

var endsymbol = '*'

func main() {

}

func (t *Trie) insertToTrie(str string) {
	t.populateSufixTrie(str)
}

func (t *Trie) populateSufixTrie(str string) {
	tempNode := t.Root

	for i := 0; i < len(str); i++ {
		charecter := rune(str[i])

		child, exist := tempNode.Children[charecter]
		if !exist {
			newTrieNode := &TrieNode{Children: make(map[rune]*TrieNode)}
			child = newTrieNode
		}
		tempNode = child
	}
	tempNode.Children[endsymbol] = nil
}

func (t *Trie) condains(str string) bool {
	tempNode := t.Root
	var charecter rune

	for i := 0; i < len(str); i++ {
		charecter = rune(str[i])

		child, exist := tempNode.Children[charecter]
		if !exist {
			return false
		}
		tempNode = child
	}
	_, exist := tempNode.Children[endsymbol]
	return exist
}
