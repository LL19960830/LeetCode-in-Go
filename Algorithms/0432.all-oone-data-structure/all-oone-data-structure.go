package Problem0432

import "container/heap"

// AllOne 是解题所需的结构
type AllOne struct {
	m   map[string]*item
	max maxPQ
	min minPQ
}

// Constructor initialize your data structure here.
func Constructor() AllOne {
	return AllOne{m: make(map[string]int),
		max: -1 << 63,
		min: 1<<63 - 1,
	}
}

// Inc inserts a new key <Key> with value 1. Or increments an existing key by 1.
func (a *AllOne) Inc(key string) {
	if _, ok := a.m[key]; ok {
		a.m[key]++
		if a.max < a.m[key] {
			a.max = a.m[key]
			a.maxKey = key
		}
	} else {
		a.m[key] = 1
		if a.min > 1 {
			a.min = 1
			a.minKey = key
		}
	}

}

// Dec decrements an existing key by 1. If Key's value is 1, remove it from the data structure.
func (a *AllOne) Dec(key string) {
	if _, ok := a.m[key]; !ok {
		return
	}

	a.m[key]--

	if a.min > a.m[key] {
		a.min = a.m[key]
		a.minKey = key
	}
}

// GetMaxKey returns one of the keys with maximal value.
func (a *AllOne) GetMaxKey() string {
	return a.maxKey
}

// GetMinKey returns one of the keys with Minimal value.
func (a *AllOne) GetMinKey() string {
	return a.minKey
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */

// item 是 priorityQueue 中的元素
type item struct {
	key   string
	value int
	// index 是 item 在 heap 中的索引号
	// item 加入 Priority Queue 后， Priority 会变化时，很有用
	// 如果 item.priority 一直不变的话，可以删除 index
	index int
}

// priorityQueue implements heap.Interface and holds items.
type minPQ []*item

func (pq minPQ) Len() int { return len(pq) }

func (pq minPQ) Less(i, j int) bool {
	return pq[i].value < pq[j].value
}

func (pq minPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push 往 pq 中放 item
func (pq *minPQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(*item)
	item.index = n
	*pq = append(*pq, item)
}

// Pop 从 pq 中取出最优先的 item
func (pq *minPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an item in the queue.
func (pq *minPQ) update(item *item, value string, priority int) {
	item.key = value
	item.value = priority
	heap.Fix(pq, item.index)
}

type maxPQ minPQ

func (pq maxPQ) Less(i, j int) bool {
	return pq[i].value > pq[i].value
}
