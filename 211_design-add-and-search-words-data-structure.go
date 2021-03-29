package leetcode

type WordDictionary struct {
	root *wordNode
}

type wordNode struct {
	flag bool
	next [26]*wordNode
}

/** Initialize your data structure here. */
func Constructor8() WordDictionary {
	return WordDictionary{
		root: new(wordNode),
	}
}

func (this *WordDictionary) AddWord(word string) {
	node := this.root
	for _, char := range []byte(word) {
		if node.next[char-'a'] == nil {
			node.next[char-'a'] = new(wordNode)
		}
		node = node.next[char-'a']
	}

	node.flag = true
}

func (this *WordDictionary) Search(word string) bool {
	return searchForWord(this.root, []byte(word))
}

func searchForWord(node *wordNode, word []byte) bool {
	for i := 0; i < len(word); i++ {
		if node == nil {
			return false
		}

		if word[i] == '.' {
			ok := false
			for _, n := range node.next {
				if n != nil {
					ok = searchForWord(n, word[i+1:])
				}
				if ok {
					return true
				}
			}
			return false
		}

		node = node.next[word[i]-'a']
	}

	return node != nil && node.flag
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
