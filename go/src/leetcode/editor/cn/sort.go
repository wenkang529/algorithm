package cn

import "fmt"

func quickSort(l []int, left, right int) {
	if len(l) < 1 {
		return
	}
	if right - left < 1 {
		return
	}
	ston := l[left]
	subleft := left
	subright := right
	for subleft < subright {
		for l[subleft] < ston && subleft < subright {
			subleft++
		}
		for l[subright] > ston && subleft < subright {
			subright--
		}
		l[subleft], l[subright] = l[subright], l[subleft]
	}
	if l[subleft] < l[left] {
		l[left], l[subleft] = l[subleft], l[left]
	} else {

	}
	quickSort(l, left, subleft-1)
	quickSort(l, subleft+1, right)
	return
}

func SortRun() {
	l := []int{1, 3, 2, 4, 9, 6}
	quickSort(l, 0, 5)
	fmt.Println(l)
}
