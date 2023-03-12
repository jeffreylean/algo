package algo

type Node struct {
	Prev  *Node
	Next  *Node
	Key   int
	Value int
}

type LRU struct {
	cache map[int]*Node
	head  *Node
	tail  *Node
	size  int
}

func NewLRU(size int) LRU {
	lru := LRU{
		cache: make(map[int]*Node, size),
		size:  size,
		head:  &Node{},
		tail:  &Node{},
	}
	lru.head.Next = lru.tail
	lru.tail.Prev = lru.head
	return lru
}

func (this *LRU) Get(key int) int {
	if node, ok := this.cache[key]; !ok {
		return -1
	} else {
		if len(this.cache) == this.size {
			this.removeFromBehind()
		}
		this.addToFront(node)
		return node.Value
	}
}

func (this *LRU) Put(key, value int) {
	// Overwrite exising value
	if node, ok := this.cache[key]; ok {
		node.Value = value
		this.addToFront(node)
		return
	}

	if len(this.cache) == this.size {
		this.removeFromBehind()
	}
	node := &Node{Key: key, Value: value}
	this.cache[key] = node
	this.addToFront(node)
}

func (this *LRU) addToFront(node *Node) {
	temp := this.head.Next
	this.head.Next = node
	temp.Prev = node
	node.Next = temp
	node.Prev = this.head
}

func (this *LRU) removeFromBehind() {
	temp := this.tail.Prev
	delete(this.cache, temp.Key)
	temp.Prev.Next = this.tail
	this.tail.Prev = temp.Prev
}
