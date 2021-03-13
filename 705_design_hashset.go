package leetcode

type MyHashSet struct {
	keys []int
}

/** Initialize your data structure here. */
func Constructor5() MyHashSet {
	return MyHashSet{
		keys: make([]int, 0),
	}
}

func (this *MyHashSet) Add(key int) {
	for _, k := range this.keys {
		if k == key {
			return
		}
	}

	this.keys = append(this.keys, key)
}

func (this *MyHashSet) Remove(key int) {
	for i, k := range this.keys {
		if k == key {
			this.keys = append(this.keys[:i], this.keys[i+1:]...)
			return
		}
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	for _, k := range this.keys {
		if k == key {
			return true
		}
	}

	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
