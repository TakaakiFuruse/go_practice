package largestNumber

import "sort"

func largestOfFour(ar [][]int) []int {
	resAr := []int{}
	for i := 0; i <= len(ar)-1; i++ {
		sort.Sort(retLargest(ar[i]))
		resAr = append(resAr, ar[i][0])
	}
	return resAr
}

type retLargest []int

func (el retLargest) Len() int {
	return len(el)
}

func (el retLargest) Swap(i, j int) {
	el[i], el[j] = el[j], el[i]
}

func (el retLargest) Less(i, j int) bool {
	return el[i] > el[j]
}
