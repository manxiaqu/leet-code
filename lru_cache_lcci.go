package leetcode

/*
type LRUCache struct {
	cache map[int]int

	// store key in this array
	lru  []int
	caps int
}

func Constructor4(capacity int) LRUCache {
	return LRUCache{
		cache: make(map[int]int, 0),
		lru:   make([]int, 0, capacity),
		caps:  capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	// fmt.Println("try to get key", key)
	v, ok := this.cache[key]
	if !ok {
		return -1
	}

	this.update(key)
	return v
}

func (this *LRUCache) Put(key int, value int) {
	// set value if key exist
	_, ok := this.cache[key]
	if ok {
		this.cache[key] = value
		this.update(key)
	} else {
		// remove one
		if len(this.lru) == this.caps {
			delete(this.cache, this.lru[0])
			this.lru = this.lru[1:]
		}

		this.lru = append(this.lru, key)
		this.cache[key] = value
	}
}

func (this *LRUCache) update(key int) {
	// try to add this to lru end
	for i := 0; i < len(this.lru); i++ {
		if this.lru[i] == key && i != len(this.lru)-1 {
			this.lru = append(this.lru[:i], this.lru[i+1:]...)
			this.lru = append(this.lru, key)
			break
		}
	}
}
*/
