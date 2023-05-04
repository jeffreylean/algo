package algo

type LRUNode struct {
	Prev  *LRUNode
	Next  *LRUNode
	Key   int
	Value int
}

type LRU struct {
	cache map[int]*LRUNode
	head  *LRUNode
	tail  *LRUNode
	size  int
}

func NewLRU(size int) LRU {
	lru := LRU{
		cache: make(map[int]*LRUNode, size),
		size:  size,
		head:  &LRUNode{},
		tail:  &LRUNode{},
	}
	lru.head.Next = lru.tail
	lru.tail.Prev = lru.head
	return lru
}

func (this *LRU) Get(key int) int {
	if LRUNode, ok := this.cache[key]; !ok {
		return -1
	} else {
		if len(this.cache) == this.size {
			this.removeFromBehind()
		}
		this.addToFront(LRUNode)
		return LRUNode.Value
	}
}

func (this *LRU) Put(key, value int) {
	// Overwrite exising value
	if LRUNode, ok := this.cache[key]; ok {
		LRUNode.Value = value
		this.addToFront(LRUNode)
		return
	}

	if len(this.cache) == this.size {
		this.removeFromBehind()
	}
	LRUNode := &LRUNode{Key: key, Value: value}
	this.cache[key] = LRUNode
	this.addToFront(LRUNode)
}

func (this *LRU) addToFront(LRUNode *LRUNode) {
	temp := this.head.Next
	this.head.Next = LRUNode
	temp.Prev = LRUNode
	LRUNode.Next = temp
	LRUNode.Prev = this.head
}

func (this *LRU) removeFromBehind() {
	temp := this.tail.Prev
	delete(this.cache, temp.Key)
	temp.Prev.Next = this.tail
	this.tail.Prev = temp.Prev
}
