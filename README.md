# golang_lfu_impls

LFU(Least Frequency Used) 是一種演算法把最近最不常使用的 Cache 值除掉


## 參考

https://halfrost.com/lru_lfu_interview/

## 觀察

首先是 cache 有一個最大容量 K 

並且在超過這個容量 K 時 ， 有值要插入需要把最近最小少存取的值移除

再在把這個新的值放入

## 關鍵點

需要紀錄目前最小的存取值透過這個值找到對應的 Cache

透過一個 value map nodes 去紀錄每個 key 對應的 node

透過一個 list map lists 去紀錄每個 frequency 對應的 list

每次取 key 的 valaue 時，如果 key 值存在， 則取出 key 對應的 node

透過 lists 查訊對應 frequency list 的最後一個 node 刪除

把 frequency + 1 並且放入對應 frequency + 1 list 的最前面

檢查原本 frequency 是否為最小值 min ， 如果是則把最小值更新為 原本 frequency + 1

## 實作

```golang
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

```