// https://leetcode.com/problems/unique-paths
package main

func main() {

}

var pathMap [101][101]int

func uniquePaths(m, n int) int {
	if pathMap[m][n] != 0 {
		return pathMap[m][n]
	}
	for i := 1; i <= m; i++ {
		pathMap[i][1] = 1
	}
	for i := 1; i <= n; i++ {
		pathMap[1][i] = 1
	}
	for i := 2; i <= m; i++ {
		for j := 2; j <= n; j++ {
			pathMap[i][j] = pathMap[i-1][j] + pathMap[i][j-1]
		}
	}
	return pathMap[m][n]
}
