package leetcode

type Trie struct {
	root *node
}

/** Initialize your data structure here. */
func Constructor7() Trie {
	return Trie{
		root: &node{},
	}
}

type node struct {
	val  []byte
	flag bool

	children []*node
}

func newNode(val []byte, flag bool) *node {
	return &node{
		children: make([]*node, 0),
		val:      val,
		flag:     flag,
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	insertToTrie(this.root, []byte(word))
}

// root is parent, children is the place to search
func insertToTrie(root *node, bs []byte) {
	if len(bs) == 0 {
		return
	}
	// the first time insert word
	if len(root.children) == 0 {
		root.children = append(root.children, newNode(bs, true))
		return
	}

	find := false
	for _, nd := range root.children {
		i := 0

		for j := 0; j < len(nd.val); j++ {
			if i == len(bs) || bs[i] != nd.val[j] {
				break
			}
			find = true
			i++
		}

		if !find {
			continue
		}

		// at this place find the same prefix
		// nd.val is search over
		if i == len(nd.val) {
			// bs is not search over
			if i < len(bs) {
				insertToTrie(nd, bs[i:])
				return
			}
			// bs is search over
			// the word is already a node, set flag to true
			if i == len(bs) {
				nd.flag = true
				return
			}
		}

		// at the place, i < len(val), i can't > len(val)
		// find same prefix [0: i)
		// nd should break into 2 node
		nf := newNode(nd.val[i:], nd.flag)
		nf.children = nd.children

		nd.children = make([]*node, 0)
		nd.children = append(nd.children, nf)
		nd.val = nd.val[:i]
		nd.flag = true

		if i < len(bs) {
			ns := newNode(bs[i:], true)
			nd.children = append(nd.children, ns)
			// set flag to false if bs isn't over
			nd.flag = false
		}
		return
	}

	// at this place, not find same prefix in this node
	root.children = append(root.children, newNode(bs, true))
}

// len(bs) always > 0
func searchInTrie(root *node, bs []byte) (startWith, equal bool) {
	if len(root.children) == 0 {
		return
	}

	find := false
	for _, nd := range root.children {
		i := 0

		for j := 0; j < len(nd.val); j++ {
			if i == len(bs) || bs[i] != nd.val[j] {
				break
			}
			find = true
			i++
		}

		if !find {
			continue
		}

		// at this place find the same prefix
		// nd.val is search over
		if i == len(nd.val) {
			// bs is not search over
			if i < len(bs) {
				return searchInTrie(nd, bs[i:])
			}
			// bs is search over
			// the word is already a node, set flag to true
			if i == len(bs) {
				startWith = true
				equal = nd.flag
				return
			}
		}

		// at the place, i < len(val), i can't > len(val)
		// find same prefix [0: i)
		// nd should break into 2 node
		if i == len(bs) {
			startWith = true
			equal = false
			return
		}

		return
	}

	return
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	_, equal := searchInTrie(this.root, []byte(word))

	return equal
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	startWith, _ := searchInTrie(this.root, []byte(prefix))

	return startWith
}
