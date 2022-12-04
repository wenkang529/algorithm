package cn

// 堆，就是用数据来表示二叉树，有个前提条件就是必须是完全二叉树，极为除了叶子节点，都必须有两个
// 键值或索引总是小于（或者大于）它的父节点，这也成为最小堆和最大堆。
/*
例如一个数字，第i个数，它的父，子节点如下
parent(i) = floor((i - 1)/2)
left(i)   = 2i + 1
right(i)  = 2i + 2
*/

// 堆排序步骤
// 1. 构建大顶推/小顶堆 （如果按照生序排序，那就是大顶堆）
// 2. 交换收尾数据，并重新构建n-1的大顶堆
// 3. 重复n-1次，直到完成

// 如何构建大顶堆呢？
// 1. 找到最后一个非叶子节点 （i-1）/2 其中i为 len -1
// 2. 对比节点和两个子节点的最大值，放父节点位置（交换）
// 重复所有非叶子节点

func buildHead(l []int) {
	if len(l) < 2 {
		return
	}
	last := (len(l) - 1 - 1) / 2
	for i := last; i >= 0; i-- {
		left := 2*i + 1
		right := 2*i + 2
		if right == len(l) {
			right = i
		}
		if l[left] >= l[right] {
			if l[left] > l[i] {
				l[left], l[i] = l[i], l[left]
			}
		} else {
			if l[right] > l[i] {
				l[right], l[i] = l[right], l[i]
			}
		}
	}
}

func HeapSort(l []int) {
	length := len(l)
	if length < 2 {
		return
	}
	for i := length; i > 0; i-- {
		buildHead(l[:i])
		l[0], l[i-1] = l[i-1], l[0]
	}
	return
}
