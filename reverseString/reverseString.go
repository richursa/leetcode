package main

func main() {
	reverseString([]byte{'a', 'b', 'c', 'd', 'e'})
}
func reverseString(s []byte) {
	temp := byte('a')
	for i := 0; i <= (len(s)/2)-1; i++ {
		temp = s[i]
		s[i] = s[len(s)-i-1]
		s[len(s)-i-1] = temp
	}
	// fmt.Println(string(s))
}
