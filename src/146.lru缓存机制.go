/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU缓存机制
 *
 * https://leetcode-cn.com/problems/lru-cache/description/
 *
 * algorithms
 * Medium (43.74%)
 * Likes:    271
 * Dislikes: 0
 * Total Accepted:    20.6K
 * Total Submissions: 47.1K
 * Testcase Example:  '["LRUCache","put","put","get","put","get","put","get","get","get"]\n' +
  '[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'
 *
 * 运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
 * 
 * 获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
 * 写入数据 put(key, value) -
 * 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。
 * 
 * 进阶:
 * 
 * 你是否可以在 O(1) 时间复杂度内完成这两种操作？
 * 
 * 示例:
 * 
 * LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );
 * 
 * cache.put(1, 1);
 * cache.put(2, 2);
 * cache.get(1);       // 返回  1
 * cache.put(3, 3);    // 该操作会使得密钥 2 作废
 * cache.get(2);       // 返回 -1 (未找到)
 * cache.put(4, 4);    // 该操作会使得密钥 1 作废
 * cache.get(1);       // 返回 -1 (未找到)
 * cache.get(3);       // 返回  3
 * cache.get(4);       // 返回  4
 * 
 * 
 */

// @lc code=start

func Constructor(capacity int) LRUCache {
	m := make(map[int]*CacheNode)
	c := LRUCache{Cap: capacity, Map: m}
	return c
}

func (this *LRUCache) Get(key int) int {
	found, ok := this.Map[key]
	if !ok {
		return -1
	}
	if this.Head == found {
		return found.Val
	}
	if this.Tail == found {
		this.Tail = found.Prev
	}
	// move found to head
	if found.Next != nil {
		found.Next.Prev = found.Prev
	}
	if found.Prev != nil {
		found.Prev.Next = found.Next
	}
	this.Head.Prev, found.Next = found, this.Head
	this.Head = found
	return found.Val
}

func (this *LRUCache) Put(key int, value int) {
	found, ok := this.Map[key]
	if ok {
		found.Val = value
		_ = this.Get(found.Key) // to move found node to head
		return
	}

	// add to head
	n := &CacheNode{Key: key, Val: value}

	if len(this.Map) == 0 {
		this.Tail = n
	} else {
		this.Head.Prev, n.Next = n, this.Head
	}
	this.Map[key], this.Head = n, n
	if this.Cap == this.Len {
		delete(this.Map, this.Tail.Key)
		this.Len, this.Tail = this.Len-1, this.Tail.Prev
	}
	this.Len++
}

type CacheNode struct {
	Next *CacheNode
	Prev *CacheNode
	Key  int
	Val  int
}

type LRUCache struct {
	Cap  int
	Len  int
	Head *CacheNode
	Tail *CacheNode
	Map  map[int]*CacheNode
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

