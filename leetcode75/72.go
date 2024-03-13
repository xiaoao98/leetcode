package leetcode75

type trieNode struct {
	isEnd    bool
	children []*trieNode
}

func newTrieNode() *trieNode {
	return &trieNode{
		isEnd:    false,
		children: make([]*trieNode, 26),
	}
}

type trie struct {
	root         *trieNode
	resultBuffer []string
}

func (t *trie) insert(word string) {
	curr := t.root
	for _, ch := range word {
		idx := ch - 'a'
		if curr.children[idx] == nil {
			curr.children[idx] = newTrieNode()
		}
		curr = curr.children[idx]
	}
	curr.isEnd = true
}

func (t *trie) search(word string) [][]string {
	curr := t.root
	ret := make([][]string, len(word))
	wordNow := ""
	for i, ch := range word {
		idx := ch - 'a'
		if curr.children[idx] == nil {
			return ret
		}
		curr = curr.children[idx]
		wordNow += string(ch)
		t.dfsWithPrefix(curr, wordNow)
		for _, result := range t.resultBuffer {
			ret[i] = append(ret[i], result)
		}
		t.resultBuffer = t.resultBuffer[:0]
	}
	return ret
}

func (t *trie) dfsWithPrefix(node *trieNode, word string) {
	if len(t.resultBuffer) == 3 {
		return
	}
	if node.isEnd == true {
		t.resultBuffer = append(t.resultBuffer, word)
	}
	for i, ele := range node.children {
		if ele != nil {
			ch := 'a' + i
			t.dfsWithPrefix(ele, word+string(ch))
		}
	}
}

func suggestedProducts(products []string, searchWord string) [][]string {
	newTrie := &trie{root: newTrieNode(), resultBuffer: make([]string, 0)}
	for _, product := range products {
		newTrie.insert(product)
	}
	return newTrie.search(searchWord)
}
