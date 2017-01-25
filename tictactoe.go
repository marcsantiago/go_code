package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/fatih/color"
)

const (
	x        = "X"
	o        = "O"
	player1  = "player1"
	player2  = "player2"
	computer = "computer"
)

func checkBoard(bo []string, le string) (bool, error) {
	win := ((bo[7] == le && bo[8] == le && bo[9] == le) || // across the top
		(bo[4] == le && bo[5] == le && bo[6] == le) || // across the middle
		(bo[1] == le && bo[2] == le && bo[3] == le) || // across the bottom
		(bo[7] == le && bo[4] == le && bo[1] == le) || // down the left side
		(bo[8] == le && bo[5] == le && bo[2] == le) || // down the middle
		(bo[9] == le && bo[6] == le && bo[3] == le) || // down the right side
		(bo[7] == le && bo[5] == le && bo[3] == le) || // diagonal
		(bo[9] == le && bo[5] == le && bo[1] == le)) // diagonal

	if win {
		return true, nil
	}

	empty := false
	for i := 0; i < 10; i++ {
		if bo[0] == " " {
			empty = true
			break
		}
	}

	if !empty {
		return false, errors.New("Board is full")
	}

	return false, nil
}

func drawBoard(bo []string) {
	fmt.Printf("   | %s | %s | %s |\n", bo[0], bo[1], bo[2])
	fmt.Printf("-------------------\n")
	fmt.Printf("   | %s | %s | %s |\n", bo[3], bo[4], bo[5])
	fmt.Printf("-------------------\n")
	fmt.Printf("   | %s | %s | %s |\n", bo[6], bo[7], bo[8])
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
	xColor := color.New(color.FgHiBlue, color.Bold).SprintFunc()
	oColor := color.New(color.FgHiRed, color.Bold).SprintFunc()
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
				board[move] = fmt.Sprintf("%s", xColor(x))
				turn = 1
				drawBoard(board)
				win, err := checkBoard(board, x)
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
				board[move] = fmt.Sprintf("%s", oColor(o))
				turn = 0
				drawBoard(board)
				win, err := checkBoard(board, o)
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
