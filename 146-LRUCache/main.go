package main

type Node struct {
	key, val  int
	pre, next *Node
}

type LRUCache struct {
	length     int
	capacity   int
	cache      map[int]*Node
	head, tail *Node
}

func Constructor(capacity int) LRUCache {
	head, tail := new(Node), new(Node)
	head.next = tail
	tail.pre = head
	return LRUCache{0, capacity, make(map[int]*Node), head, tail}
}

func (this *LRUCache) Get(key int) int {
	node, head := this.cache[key], this.head
	if node != nil {
		node.pre.next = node.next
		node.next.pre = node.pre
		node.pre = head
		node.next = head.next
		head.next.pre = node
		head.next = node
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	node, head, tail := this.cache[key], this.head, this.tail
	if node != nil {
		this.Get(key)
		node.val = value
	} else {
		node = &Node{key, value, head, head.next}
		head.next.pre = node
		head.next = node
		this.cache[key] = node
		if this.length < this.capacity {
			this.length++
		} else {
			this.cache[tail.pre.key] = nil
			tail.pre.pre.next = tail
			tail.pre = tail.pre.pre
		}
	}
}
