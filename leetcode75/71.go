package leetcode75

type Trie struct {
	isEnd    bool
	children []*Trie
}

func Constructor71() Trie {
	return Trie{
		isEnd:    false,
		children: make([]*Trie, 26),
	}
}

func (this *Trie) Insert(word string) {
	curr := this
	for _, ch := range word {
		idx := ch - 'a'
		if curr.children[idx] == nil {
			newChild := Constructor71()
			curr.children[idx] = &newChild
		}
		curr = curr.children[idx]
	}
	curr.isEnd = true
}

func (this *Trie) Search(word string) bool {
	curr := this
	for _, ch := range word {
		idx := ch - 'a'
		if curr.children[idx] == nil {
			return false
		} else {
			curr = curr.children[idx]
		}
	}
	return curr.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this
	for _, ch := range prefix {
		idx := ch - 'a'
		if curr.children[idx] == nil {
			return false
		} else {
			curr = curr.children[idx]
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