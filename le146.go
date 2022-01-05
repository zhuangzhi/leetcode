package leetcode

// https://leetcode-cn.com/problems/lru-cache/
// 146. LRU 缓存
// 请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
// 实现 LRUCache 类：
// LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
// 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

type LRUCache struct {
	cap, size  int
	cache      map[int]*LRUCacheNode
	head, tail *LRUCacheNode
}

type LRUCacheNode struct {
	key, val  int
	pre, next *LRUCacheNode
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{cap: capacity, size: 0}
	cache.head = &LRUCacheNode{}
	cache.tail = &LRUCacheNode{}
	cache.head.next = cache.tail
	cache.tail.pre = cache.head
	cache.cache = make(map[int]*LRUCacheNode, capacity)
	return cache
}

func (this *LRUCache) removeNode(node *LRUCacheNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) addToHead(node *LRUCacheNode) {
	node.next = this.head.next
	node.pre = this.head
	this.head.next.pre = node
	this.head.next = node
}

func (this *LRUCache) moveToHead(node *LRUCacheNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.cache[key]; ok {
		this.moveToHead(v)
		return v.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.cache[key]; ok {
		this.moveToHead(v)
		v.val = value
		return
	}

	this.size++
	node := &LRUCacheNode{key: key, val: value}
	this.cache[key] = node
	this.addToHead(node)
	if this.size > this.cap {
		tail := this.removeTail()
		delete(this.cache, tail.key)
		this.size--
	}
}

func (this *LRUCache) removeTail() *LRUCacheNode {
	node := this.tail.pre
	this.removeNode(node)
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
