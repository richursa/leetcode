package main

func main() {

}

func isAnagram(s string, t string) bool {
	var firstMap = make(map[rune]int)
	var secondMap = make(map[rune]int)
	for _, uniCodeChar := range s {
		firstMap[uniCodeChar]++
	}
	for _, uniCodechar := range t {
		secondMap[uniCodechar]++
	}
	if len(firstMap) != len(secondMap) {
		return false
	}
	for uniCodeChar, count := range firstMap {
		if secondMap[uniCodeChar] != count {
			return false
		}
	}
	return true
}
