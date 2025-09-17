// https://leetcode.com/problems/implement-trie-prefix-tree/?envType=study-plan-v2&envId=top-interview-150

package trie

type Trie struct {
	value *byte
	nodes []*Trie
	isEnd bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	size := len(word)
	for i, node := range this.nodes {
		if hasPrefix := node.StartsWith(string(word[0])); hasPrefix && size == 1 {
			this.nodes[i].isEnd = true
			return
		} else if hasPrefix {
			node.Insert(word[1:])
			return
		}
	}

	node := Constructor()
	value := byte(word[0])
	node.value = &value
	if size > 1 {
		node.Insert(word[1:])
	} else {
		node.isEnd = true
	}
	this.nodes = append(this.nodes, &node)
}

func (this *Trie) Search(word string) bool {
	return this.matchExact(word, true)
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.matchExact(prefix, false)
}

func (this *Trie) matchExact(word string, exact bool) bool {
	size := len(word)

	first := word[0]
	if size == 1 && this.value != nil {
		return *this.value == first && (!exact || this.isEnd)
	}

	prefix := word
	if this.value != nil {
		prefix = word[1:]
	}

	for _, node := range this.nodes {
		match := (this.value == nil || *this.value == first) && node.matchExact(prefix, exact)
		if match {
			return true
		}
	}
	return false
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

/*
  ("Trie", []),
 ("insert", ["app"]),
 ("insert", ["apple"]),
 ("insert", ["beer"]),
 ("insert", ["add"]),
 ("insert", ["jam"]),
 ("insert", ["rental"]),
 ("search", ["apps"]),
 ("search", ["app"]),
 ("search", ["ad"]),
 ("search", ["applepie"]),
 ("search", ["rest"]),
 ("search", ["jan"]),
 ("search", ["rent"]),
 ("search", ["beer"]),
 ("search", ["jam"]),
 ("startsWith", ["apps"]),
 ("startsWith", ["app"]),
 ("startsWith", ["ad"]),
 ("startsWith", ["applepie"]),
 ("startsWith", ["rest"]),
 ("startsWith", ["jan"]),
 ("startsWith", ["rent"]),
 ("startsWith", ["beer"]),
 ("startsWith", ["jam"])
*/
