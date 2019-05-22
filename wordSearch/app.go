package main

func main() {
	
}

func exist(board [][]byte, word string) bool {

	x, y := len(board), len(board[0])

	visit := make([][]bool, x)

	for i := 0; i < x; i++ {
		visit[i] = make([]bool, y)
	}

	for i := 0; i <x; i++ {
		for j:= 0; j<y; j++ {
			if searched(board, visit, word, i,j,0) {
				return true
			}
		}
	}

	return false
}


func searched(board [][]byte, visited [][] bool, word string, i, j, start int) bool  {

	if start ==  len(word) {
		return true
	}

	// out of index false
	// visited false
	// not equal false
	if i<0 || i >=len(board) || j <0 || j>= len(board[0]) || visited[i][j] || board[i][j] != byte([]rune(word)[start])  {

		return false
	}
	// visited
	visited[i][j] = true

	// search all direction

	result := searched(board, visited, word, i+1, j, start +1)  ||
		searched(board, visited, word, i, j+1, start+1) ||
		searched(board, visited, word, i-1, j, start+1) ||
		searched(board, visited, word, i, j-1, start+1)

	// visited false since we don't get the result from this point
	// result true end of search
	// result false roll back the visited point
	visited[i][j] = false

	return result

}

