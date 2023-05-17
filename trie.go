package algo

type TrieNode struct {
	childrens [26]*TrieNode

	wordEnds bool
}

type trie struct {
	root *TrieNode
}

func trieData() *trie {
	t := new(trie)
	return t
}

func (t *trie) insert(word string) {
	current := t.root

	for _, wr := range word {
		index := wr - 'a'
		if current.childrens[index] == nil {
			current.childrens[index] = new(TrieNode)
		}
		current = current.childrens[index]
	}
	current.wordEnds = true
}

func (t *trie) search(word string) int {
	current := t.root
	for _, wr := range word {
		index := wr - 'a'
		if current.childrens[index] == nil {
			return 0
		}
	}
	if current.wordEnds == true {
		return 1
	}
	return 0
}
