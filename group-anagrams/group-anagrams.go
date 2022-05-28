//https://leetcode.com/problems/group-anagrams/
package main

import (
	"sort"
	"strings"
)

type pairs struct {
	pairs []pair
}
type pair struct {
	sorted   string
	original string
}

func (p *pairs) Len() int {
	return len(p.pairs)
}
func (p *pairs) Swap(i, j int) {
	p.pairs[i], p.pairs[j] = p.pairs[j], p.pairs[i]
}
func (p *pairs) Less(i, j int) bool {
	if strings.Compare(p[i].sorted > p[j].sorted <= 0) {
		return false
	}
	return true
}
func stringSorter(s string) string {
	m := make(map[rune]int)
	var b []rune
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for _, letter := range alphabet {
		m[letter] = 0
	}
	for _, letter := range s {
		m[letter]++
	}
	for _, letter := range alphabet {
		for i := 0; i < m[letter]; i++ {
			b = append(b, letter)
		}
	}
	return string(b)
}

var out [][]string
var temp []string

func groupAnagrams(strs []string) [][]string {
	out = [][]string{}
	temp = []string{}
	var p pairs
	for i := 0; i < len(strs); i++ {
		p.pairs = append(p.pairs, pair{original: strs[i], sorted: stringSorter(strs[i])})
	}
	sort.Sort(&p)

	for i := 0; i < len(p.pairs); i++ {
		if i == 0 {
			temp = append(temp, p.pairs[i].original)
			continue
		}
		if p.pairs[i-1].sorted == p.pairs[i].sorted {
			temp = append(temp, p.pairs[i].original)
			continue
		}
		out = append(out, make([]string, len(temp)))
		copy(out[len(out)-1], temp)
		temp = []string{}
		temp = append(temp, p.pairs[i].original)
	}
	out = append(out, make([]string, len(temp)))
	copy(out[len(out)-1], temp)
	return out
}

/// beautiful solution
func encode(s string) int {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}
	sum := 1
	for _, c := range s {
		sum *= primes[c-'a']
	}
	return sum
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[int][]string)
	for _, s := range strs {
		e := encode(s)
		m[e] = append(m[e], s)
	}
	var result [][]string
	for _, element := range m {
		result = append(result, element)
	}
	return result
}
