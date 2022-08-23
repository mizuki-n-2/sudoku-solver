package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"
)

type Board [9][9]int

func parseInput(input string) Board {
	var board Board
	row := 0
	for i := 0; i < len(input); i++ {
		if i != 0 && i%9 == 0 {
			row++
		}

		currentChar := string(input[i])

		if currentChar == "." {
			board[row][i%9] = 0
			continue
		}

		// TODO: エラー処理
		board[row][i%9], _ = strconv.Atoi(currentChar)
	}

	return board
}

func printBoard(b Board) {
	fmt.Println("+---+---+---+")
	for row := 0; row < 9; row++ {
		fmt.Print("|")
		for col := 0; col < 9; col++ {
			fmt.Print(b[row][col])
			if col%3 == 2 {
				fmt.Print("|")
			}
			if col%9 == 8 {
				fmt.Print("\n")
			}
		}
		if row%3 == 2 {
			fmt.Println("+---+---+---+")
		}
	}
}

func isValid(b *Board) bool {
	// 行チェック
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[b[row][col]]++
		}

		if duplicated(counter) {
			return false
		}
	}

	// 列チェック
	for col := 0; col < 9; col++ {
		counter := [10]int{}
		for row := 0; row < 9; row++ {
			counter[b[row][col]]++
		}

		if duplicated(counter) {
			return false
		}
	}

	// 3×3チェック
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			counter := [10]int{}
			for row := 0; row < 3; row++ {
				for col := 0; col < 3; col++ {
					counter[b[row+i][col+j]]++
				}
			}

			if duplicated(counter) {
				return false
			}
		}
	}

	return true
}

// 重複判定
func duplicated(counter [10]int) bool {
	for i, count := range counter {
		if i == 0 {
			continue
		}

		if count >= 2 {
			return true
		}
	}
	return false
}

// 終了判定
func isFull(b *Board) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b[i][j] == 0 {
				return false
			}
		}
	}

	return true
}

func backtrack(b *Board) bool {
	// テストの時はスキップ(暫定)
	if flag.Lookup("test.v") == nil {
		time.Sleep(1 * time.Second)
		printBoard(*b)
	}

	if isFull(b) {
		return true
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b[i][j] == 0 {
				for candidate := 9; candidate >= 1; candidate-- {
					b[i][j] = candidate
					if isValid(b) {
						if backtrack(b) {
							return true
						}

						b[i][j] = 0
					} else {
						b[i][j] = 0
					}
				}

				return false
			}
		}
	}

	return false
}

func solve(input string) (Board, bool) {
	board := parseInput(input)

	printBoard(board)

	return board, backtrack(&board)
}

func main() {
	var input string
	fmt.Scan(&input)

	board, result := solve(input)

	if result {
		fmt.Println("Success!!")
		printBoard(board)
	} else {
		fmt.Println("can't solve...")
	}
}
