package lfu

import "container/list"

type node struct {
	K, V, frequency int
}
type LFUCache struct {
	capacity int
	min      int
	nodes    map[int]*list.Element // record all node by freq
	lists    map[int]*list.List    // record all list by freq
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		nodes:    make(map[int]*list.Element),
		lists:    make(map[int]*list.List),
		capacity: capacity,
		min:      0,
	}
}

func (cache *LFUCache) Get(key int) int {
	value, ok := cache.nodes[key]
	if !ok {
		return -1
	}
	currentNode := value.Value.(*node)
	// remove node from origin frequency
	cache.lists[currentNode.frequency].Remove(value)
	currentNode.frequency++
	if _, ok := cache.lists[currentNode.frequency]; !ok {
		cache.lists[currentNode.frequency] = list.New()
	}
	newList := cache.lists[currentNode.frequency]
	newNode := newList.PushFront(currentNode)
	cache.nodes[key] = newNode
	if currentNode.frequency-1 == cache.min && cache.lists[currentNode.frequency-1].Len() == 0 {
		cache.min++
	}
	return currentNode.V
}

func (cache *LFUCache) Put(key int, value int) {
	if cache.capacity == 0 {
		return
	}
	// 有則更新值跟頻率
	if currentValue, ok := cache.nodes[key]; ok {
		currentNode := currentValue.Value.(*node)
		currentNode.V = value
		cache.Get(key) // update frequency
		return
	}
	// 如果沒有且滿了則刪除
	if cache.capacity == len(cache.nodes) {
		currentList := cache.lists[cache.min]
		backNode := currentList.Back()
		delete(cache.nodes, backNode.Value.(*node).K)
		currentList.Remove(backNode)
	}
	// 輸入
	cache.min = 1
	currentNode := &node{
		K:         key,
		V:         value,
		frequency: 1,
	}
	if _, ok := cache.lists[1]; !ok {
		cache.lists[1] = list.New()
	}
	newList := cache.lists[1]
	newNode := newList.PushFront(currentNode)
	cache.nodes[key] = newNode
}
