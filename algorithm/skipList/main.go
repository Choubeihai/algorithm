package main

import (
	"math/rand"
	"time"
)

var MaxLevel = 15 // 最多只有15，index从0开始，共16层。层数越小，节点越多。
var levelIndex = 0
var dummy = &SkipListNode{}

func Get(val int) *SkipListNode {
	p := dummy
	for i := levelIndex; i >= 0; i-- {
		for p.Next[i] != nil && p.Next[i].Val < val {
			p = p.Next[i]
		}
	}
	if p.Next[0] != nil && p.Next[0].Val == val {
		return p.Next[0]
	} else {
		return nil
	}
}

func Insert(val int) {
	nodeLevelIndex := RandomLevel()
	node := &SkipListNode{
		Val:   val,
		level: nodeLevelIndex,
	}
	var next = make([]*SkipListNode, nodeLevelIndex+1)
	p := dummy
	for i := nodeLevelIndex; i >= 0; i-- {
		for p.Next[i] != nil && p.Next[i].Val < val {
			p = p.Next[i]
		}
		next[i] = p
	}
	for i := 0; i <= nodeLevelIndex; i++ {
		node.Next[i] = next[i].Next[i]
		next[i].Next[i] = node
	}
}

func Delete(val int) {

}

func RandomLevel() int {
	var level = 0
	for i := 0; i < MaxLevel; i++ {
		rand.Seed(time.Now().Unix())
		num := rand.Int31n(100)
		if num%2 == 0 {
			level++
		}
	}
	return level
}

type SkipListNode struct {
	Val   int
	Next  []*SkipListNode
	level int
}
