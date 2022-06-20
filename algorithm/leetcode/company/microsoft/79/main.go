package main

import "fmt"

func main() {
	var board = [][]byte{{'a', 'a'}}
	var word = "aaa"
	fmt.Println(exist(board, word))
}

var direction [][]int
var myBoard [][]byte
var visit [][]bool
var flag bool

func exist(board [][]byte, word string) bool {
	flag = false
	direction = [][]int{{1, 0}, {0, 1}, {0, -1}, {-1, 0}}
	myBoard = board
	visit = make([][]bool, len(board))
	for i := 0; i < len(visit); i++ {
		visit[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			visit[i][j] = true
			dfs(i, j, 0, word)
			visit[i][j] = false
			if flag {
				return true
			}
		}
	}
	return false
}

func dfs(x, y int, index int, word string) {
	if myBoard[x][y] != word[index] {
		return
	} else {
		if index == len(word)-1 {
			flag = true
			return
		}
	}
	for i := 0; i < len(direction); i++ {
		if flag {
			return
		}
		newX := x + direction[i][0]
		newY := y + direction[i][1]
		if newX >= 0 && newX < len(myBoard) && newY >= 0 && newY < len(myBoard[0]) && !visit[newX][newY] {
			visit[newX][newY] = true
			dfs(newX, newY, index+1, word)
			visit[newX][newY] = false
		}
	}
}
