package main

import (
	"fmt"
	"sort"
)

type pairs []pair
type pair struct {
	val       int
	frequency int
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	var p pairs
	for key := range m {
		// fmt.Println(key)
		p = append(p, pair{key, m[key]})
	}
	sort.Sort(p)
	p = p[len(p)-k:]
	var array []int
	for _, item := range p {
		array = append(array, item.val)
	}
	return array
}

func (p pairs) Less(i, j int) bool {
	return p[i].frequency < p[j].frequency
}

func (p pairs) Len() int {
	return len(p)
}

func (p pairs) Swap(i, j int) {
	p[i].val, p[j].val = p[j].val, p[i].val
	p[i].frequency, p[j].frequency = p[j].frequency, p[i].frequency
}
func main() {
	fmt.Println(topKFrequent([]int{1, 2, 2, 3, 3, 4, 5, 6, 6, 6, 6, 6, 7, 7, 8}, 4))

}
