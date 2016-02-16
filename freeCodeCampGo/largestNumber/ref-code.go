package main

import (
	"fmt"
	"sort"
)

// 実際使う場面があるかどうかは置いといて2次元配列の整数値を全てソートして、元の2次元配列の要素数に合うように並び替えるという場合、下のように書けます。
// ここでlen(a1)としてしまうと4となってしまい、実際のソートする範囲と合いませんし、sort関数は、実際ソートする要素数がわかりません。
// なのでLen関数で実際ソートする要素数を求めます。

type Data [][]int

func (t Data) Len() int {
	total := 0
	fmt.Println(t)
	for _, e := range t {
		total += len(e)
	}
	return total
}

func (t Data) Swap(i, j int) {
	Ai, Bi := t.getIndex(i)
	Aj, Bj := t.getIndex(j)
	t[Ai][Bi], t[Aj][Bj] = t[Aj][Bj], t[Ai][Bi]
}

func (t Data) Less(i, j int) bool {
	Ai, Bi := t.getIndex(i)
	Aj, Bj := t.getIndex(j)
	return t[Ai][Bi] < t[Aj][Bj]
}

func (t Data) getIndex(index int) (int, int) {
	current := 0

	for i, e := range t {
		nxt := current + len(e)
		if index < nxt {
			return i, index - current
		}
		current = nxt
	}
	return -1, -1
}

func main() {

	a1 := Data{
		[]int{13, 26},
		[]int{4, 5, 1, 3},
		[]int{32, 35, 37, 39, 10000},
		[]int{1000, 1001, 1},
	}

	sort.Sort(a1)
	fmt.Println(a1)

	//[[1 1] [3 4 5 13] [26 32 35 37 39] [1000 1001 10000]]
}
