package main

/**
使用哈希表和双向链表
双向链表按照使用顺序存储键值对，靠近头部的键值对是最近使用的，而靠近尾部的键值对是最久未使用的
哈希表即为普通的哈希映射（HashMap），通过缓存数据的键映射到其在双向链表中的位置。
这样以来，我们首先使用哈希表进行定位，找出缓存项在双向链表中的位置，随后将其移动到双向链表的头部

*/

type LRUCache struct {
	size     int
	capacity int
	m        map[int]*DLinkedNode
	head     *DLinkedNode // 虚头部
	tail     *DLinkedNode // 虚尾部
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		m:        make(map[int]*DLinkedNode),
		head:     &DLinkedNode{key: -1, value: -1},
		tail:     &DLinkedNode{key: -1, value: -1},
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.m[key]; ok {
		// 删除node
		node.prev.next = node.next
		node.next.prev = node.prev
		// 将node添加到头部
		node.prev = lru.head
		node.next = lru.head.next
		lru.head.next.prev = node
		lru.head.next = node
		return node.value
	} else {
		return -1
	}
}

func (lru *LRUCache) Put(key int, value int) {
	if node, ok := lru.m[key]; ok { // 如果有
		node.value = value
		// 删除node
		node.prev.next = node.next
		node.next.prev = node.prev
		// 将node添加到头部
		node.prev = lru.head
		node.next = lru.head.next
		lru.head.next.prev = node
		lru.head.next = node
	} else { // 如果没有
		node = &DLinkedNode{
			key:   key,
			value: value,
		}
		lru.m[key] = node
		// 将node添加到头部
		node.prev = lru.head
		node.next = lru.head.next
		lru.head.next.prev = node
		lru.head.next = node
		lru.size++
		if lru.size > lru.capacity {
			// 删除最后面的节点
			node = lru.tail.prev
			// 删除节点
			node.prev.next = node.next
			node.next.prev = node.prev
			delete(lru.m, node.key)
			lru.size--
		}
	}
}

type DLinkedNode struct {
	key   int
	value int
	prev  *DLinkedNode
	next  *DLinkedNode
}
