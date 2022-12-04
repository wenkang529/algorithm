package cn

/*
请你设计并实现一个满足 LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value)如果关键字key 已经存在，则变更其数据值value ；如果不存在，则向缓存中插入该组key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
*/

/*
典型的 双向链表 + 哈希map， 这两个组合起始就是有序的map，其中map中key的value是链表的节点
考察点：
1. 双向链表的增，删，移动（应该先删再增）
2. 执行的顺序，由于链表在增/删/移动的时候会改变值，所以其他执行动作要考虑到这个
3. 在双向链表的实现中，使用一个伪头部（dummy head）和伪尾部（dummy tail）标记界限，
这样在添加节点和删除节点的时候就不需要检查相邻的节点是否存在。在初始化的时候可以将伪头部和伪尾部链接起来。
*/


type LinkList struct {
	pre   *LinkList
	post  *LinkList
	key   int
	value int
}

type LRUCache struct {
	capacity int
	head     *LinkList
	tail     *LinkList
	cacheMap map[int]*LinkList
}

func Constructor(capacity int) LRUCache {
	res := LRUCache{
		capacity: capacity,
		head:     &LinkList{},
		tail:     &LinkList{},
	}
	res.head.post = res.tail
	res.tail.pre = res.head
	res.cacheMap = map[int]*LinkList{}
	return res
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.cacheMap[key]; ok {
		this.Del(v)
		this.Insert(this.head, v)
		return v.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.cacheMap[key]; ok {
		this.cacheMap[key].value = value
		this.Del(v)
		this.Insert(this.head, v)
		return
	}
	tmp := &LinkList{
		key:   key,
		value: value,
	}
	//不应该插入尾部，应该插入头部
	this.Insert(this.head, tmp)
	this.cacheMap[key] = tmp
	if len(this.cacheMap) > this.capacity {
		// 删词典由于是要根据tail.pre获取，所以要先于del执行，否则tai.pre会变化
		delete(this.cacheMap, this.tail.pre.key)
		this.Del(this.tail.pre)
	}
}

func (this *LRUCache) Insert(pre *LinkList, cur *LinkList) {
	cur.post = pre.post
	cur.pre = pre
	pre.post = cur
	cur.post.pre = cur
}

func (this *LRUCache) Del(cur *LinkList) {
	cur.pre.post = cur.post
	cur.post.pre = cur.pre
}

func Runcc() {
	t := Constructor(2)
	t.Put(1, 1)
	t.Put(2, 2)
	t.Get(1)
	t.Put(3, 3)
	t.Get(2)
	t.Put(4, 4)

}
