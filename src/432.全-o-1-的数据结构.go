/*
 * @lc app=leetcode.cn id=432 lang=golang
 *
 * [432] 全 O(1) 的数据结构
 *
 * https://leetcode-cn.com/problems/all-oone-data-structure/description/
 *
 * algorithms
 * Hard (35.06%)
 * Likes:    24
 * Dislikes: 0
 * Total Accepted:    1.6K
 * Total Submissions: 4.6K
 * Testcase Example:  '["AllOne","getMaxKey","getMinKey"]\n[[],[],[]]'
 *
 * 实现一个数据结构支持以下操作：
 * 
 * 
 * Inc(key) - 插入一个新的值为 1 的 key。或者使一个存在的 key 增加一，保证 key 不为空字符串。
 * Dec(key) - 如果这个 key 的值是 1，那么把他从数据结构中移除掉。否者使一个存在的 key 值减一。如果这个 key
 * 不存在，这个函数不做任何事情。key 保证不为空字符串。
 * GetMaxKey() - 返回 key 中值最大的任意一个。如果没有元素存在，返回一个空字符串""。
 * GetMinKey() - 返回 key 中值最小的任意一个。如果没有元素存在，返回一个空字符串""。
 * 
 * 
 * 挑战：以 O(1) 的时间复杂度实现所有操作。
 * 
 */

// @lc code=start

type AllOne struct {
	data map[string]*LinkedListNode
	list *LinkedList
}

func Constructor() AllOne {
	return AllOne{
		data: make(map[string]*LinkedListNode),
		list: NewLinkedList(),
	}
}

func (ao *AllOne) Inc(key string) {
	if n, ok := ao.data[key]; ok {
		n.Val++
		ao.rightAdjust(n)
	} else {
		n := &LinkedListNode{Key: key, Val: 1}
		ao.list.LPush(n)
		ao.data[key] = n
	}
}

func (ao *AllOne) rightAdjust(n *LinkedListNode) {
	for cur := n.Next; cur != nil; cur = cur.Next {
		if cur.Val > n.Val {
			ao.list.Unlink(n)
			ao.list.InsertBefore(cur, n)
			break
		} else if cur == ao.list.tail {
			ao.list.Unlink(n)
			ao.list.RPush(n)
			break
		}
	}
}

func (ao *AllOne) Dec(key string) {
	if node, ok := ao.data[key]; ok {
		if node.Val == 1 {
			delete(ao.data, key)
			ao.list.Unlink(node)
		} else {
			node.Val--
			ao.leftAdjust(node)
		}
	}
}

func (ao *AllOne) leftAdjust(n *LinkedListNode) {
	for cur := n.Next; cur != nil; cur = cur.Prev {
		if cur.Val < n.Val {
			ao.list.Unlink(n)
			ao.list.InsertAfter(cur, n)
			break
		}else if cur == ao.list.head {
			ao.list.Unlink(n)
			ao.list.LPush(n)
			break
		}
	}
}

func (ao *AllOne) GetMaxKey() string {
	if len(ao.data) == 0 {
		return ""
	}
	return ao.list.tail.Key
}

func (ao *AllOne) GetMinKey() string {
	if len(ao.data) == 0 {
		return ""
	}
	return ao.list.head.Key
}

type LinkedList struct {
	head *LinkedListNode
	tail *LinkedListNode
	size int
}

type LinkedListNode struct {
	Key  string
	Val  int
	Next *LinkedListNode
	Prev *LinkedListNode
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) LPush(n *LinkedListNode) {
	if l.head == nil {
		l.head, l.tail = n, n
	} else {
		n.Next = l.head
		l.head.Prev = n
		l.head = n
	}
	l.size++
}

func (l *LinkedList) RPush(n *LinkedListNode) {
	if l.head == nil {
		l.head, l.tail = n, n
	} else {
		l.tail.Next = n
		n.Prev = l.tail
		l.tail = n
	}
	l.size++
}

func (l *LinkedList) LPop() *LinkedListNode {
	if l.size == 0 {
		return nil
	}

	if l.head == l.tail {
		l.tail = nil
	}

	n := l.head
	l.head = n.Next

	n.Next = nil

	l.size--
	return n
}

func (l *LinkedList) RPop() *LinkedListNode {
	if l.size == 0 {
		return nil
	}

	if l.head == l.tail {
		l.head = nil
	}

	n := l.tail
	l.tail = n.Prev

	n.Prev = nil

	l.size--
	return n
}

func (l *LinkedList) Unlink(n *LinkedListNode) {
	if n == l.head {
		l.LPop()
		return
	}

	if n == l.tail {
		l.RPop()
		return
	}

	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev

	n.Prev = nil
	n.Next = nil

	l.size--
}

func (l *LinkedList) InsertBefore(pos, n *LinkedListNode) {
	if pos == l.head {
		l.LPush(n)
		return
	}

	prev := pos.Prev
	prev.Next = n
	n.Prev = prev
	n.Next = pos
	pos.Prev = n

	l.size++
}

func (l *LinkedList) InsertAfter(pos, n *LinkedListNode) {
	if pos == l.tail {
		l.RPush(n)
		return
	}

	next := pos.Next
	pos.Next = n
	n.Prev = pos
	n.Next = next
	next.Prev = n

	l.size++
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
// @lc code=end

