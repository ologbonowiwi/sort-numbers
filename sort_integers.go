package sort_integers

import "sort"

type SortableInt64 []int64

func (arr SortableInt64) Len() int {
	return len(arr)
}

func (arr SortableInt64) Less(a, b int) bool {
	return arr[a] < arr[b]
}

func (arr SortableInt64) Swap(a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}

func SortInt64(arr []int64) []int64 {
	sortable := append(make(SortableInt64, 0), arr...)

	sort.Sort(sortable)

	return sortable
}
