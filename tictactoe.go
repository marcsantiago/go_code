package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	x        = "X"
	o        = "O"
	player1  = "player1"
	player2  = "player2"
	computer = "computer"
)

func checkBoard(b []string) (bool, error) {
	return false, nil
}

func drawBoard(b []string) {
	fmt.Printf("   | %s | %s | %s |\n", b[0], b[1], b[2])
	fmt.Printf("-------------------\n")
	fmt.Printf("   | %s | %s | %s |\n", b[3], b[4], b[5])
	fmt.Printf("-------------------\n")
	fmt.Printf("   | %s | %s | %s |\n", b[6], b[7], b[8])
	return
}

func whoGoesfirst() int {
	t := time.Now()
	source := rand.NewSource(t.Unix())
	r := rand.New(source)
	return r.Perm(2)[0]
}

func getMove(board []string) int {
	possibleInput := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for {
		fmt.Println("Please select a move (1-9)")
		var input int
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}
		good := false
		for _, n := range possibleInput {
			if input == n {
				good = true
				break
			}
		}
		if !good {
			fmt.Println("Invalid input")
			continue
		}
		input--
		if board[input] != " " {
			fmt.Println("Space is already taken, try again")
			continue
		}
		return input
	}
}

func main() {
	var board []string
	fmt.Println("Welcome")
	quit := false
	for {
		if quit {
			return
		}
		turn := whoGoesfirst()
		if turn == 0 {
			fmt.Println("Player 1 goes first")
		} else {
			fmt.Println("Player 2 goes first")
		}
		board = []string{" ", " ", " ", " ", " ", " ", " ", " ", " "}
		drawBoard(board)
		for {
			// player1
			if turn == 0 {
				fmt.Println("Player 1's turn")
				move := getMove(board)
				board[move] = x
				turn = 1
				drawBoard(board)
				win, err := checkBoard(board)
				if err != nil {
					fmt.Println("Game is a tie")
					break
				}
				if win {
					fmt.Println("Player 1 Wins")
					break
				}

			} else {
				// player2
				fmt.Println("Player 2's turn")
				move := getMove(board)
				board[move] = o
				turn = 0
				drawBoard(board)
				win, err := checkBoard(board)
				if err != nil {
					fmt.Println("Game is a tie")
					break
				}
				if win {
					fmt.Println("Player 2 Wins")
					break
				}
			}
		}

		for {
			fmt.Println("Would you like to play again? (yes, no)")
			var input string
			fmt.Scan(&input)
			if strings.Contains(strings.ToLower(input), "y") {
				break
			}
			quit = true
			break
		}

	}

}
