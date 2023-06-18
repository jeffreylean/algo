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
		this.remove(LRUNode)
		this.addToFront(LRUNode)
		return LRUNode.Value
	}
}

func (this *LRU) Put(key, value int) {
	// Overwrite exising value
	if LRUNode, ok := this.cache[key]; ok {
		LRUNode.Value = value
		this.remove(LRUNode)
		this.addToFront(LRUNode)
		return
	}

	if len(this.cache) == this.size {
		delete(this.cache, this.tail.Prev.Key)
		this.remove(this.tail.Prev)
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

func (this *LRU) remove(node *LRUNode) {
	prev := node.Prev
	next := node.Next

	prev.Next = next
	next.Prev = prev
}
