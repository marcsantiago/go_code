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

var (
	xColor = color.New(color.FgHiBlue, color.Bold).SprintFunc()
	oColor = color.New(color.FgHiRed, color.Bold).SprintFunc()
)

func checkBoard(bo []string, le string) (bool, error) {
	if le == x {
		le = fmt.Sprintf("%s", xColor(le))
	} else {
		le = fmt.Sprintf("%s", oColor(le))
	}

	win := (bo[0] == le && bo[1] == le && bo[2] == le) || (bo[3] == le && bo[4] == le && bo[5] == le) || (bo[6] == le && bo[7] == le && bo[8] == le) || (bo[0] == le && bo[3] == le && bo[6] == le) || (bo[1] == le && bo[4] == le && bo[7] == le) || (bo[2] == le && bo[5] == le && bo[8] == le) || (bo[0] == le && bo[4] == le && bo[8] == le) || (bo[2] == le && bo[4] == le && bo[6] == le)
	if win {
		return true, nil
	}

	empty := false
	for i := 0; i < 9; i++ {
		if bo[i] == " " {
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

func chooseRandomMoveFromList(items []int) int {
	t := time.Now()
	source := rand.NewSource(t.Unix())
	r := rand.New(source)
	n := r.Perm(len(items))[0]
	return items[n]
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

func createCopy(board []string) []string {
	copy := []string{}
	for _, n := range board {
		copy = append(copy, n)
	}
	return copy
}

func getMoveComputer(board []string) (int, error) {
	// check to see if computer can win
	for i := 0; i < 9; i++ {
		copy := createCopy(board)
		if copy[i] == " " {
			copy[i] = fmt.Sprintf("%s", oColor(o))
			win, _ := checkBoard(copy, o)
			if win {
				return i, nil
			}
		}
	}

	// block player is they can't win
	for i := 0; i < 9; i++ {
		copy := createCopy(board)
		if copy[i] == " " {
			copy[i] = fmt.Sprintf("%s", xColor(x))
			win, _ := checkBoard(copy, x)
			if win {
				return i, nil
			}
		}
	}

	// try for a corner
	corners := []int{0, 2, 6, 8}
	emptySpaces := []int{}
	for _, c := range corners {
		if board[c] == " " {
			emptySpaces = append(emptySpaces, c)
		}
	}
	if len(emptySpaces) > 0 {
		return chooseRandomMoveFromList(emptySpaces), nil
	}

	// try middle
	if board[4] == " " {
		return 4, nil
	}

	emptySpaces = []int{}
	for i := 0; i < 9; i++ {
		if board[i] == " " {
			emptySpaces = append(emptySpaces, i)
		}
	}
	if len(emptySpaces) > 0 {
		return chooseRandomMoveFromList(emptySpaces), nil
	}

	return 0, errors.New("No moves left")
}

func main() {
	var board []string
	fmt.Println("Welcome")
	fmt.Println("Pick mode")

	var mode int
	for {
		fmt.Println("Play two player mode or vesus the computer  (0, 1)")
		_, err := fmt.Scan(&mode)
		if err != nil && mode > 1 {
			fmt.Println("Invalid entry, please enter either a 0 or a 1")
			continue
		}
		break
	}

	quit := false
	for {
		if quit {
			return
		}
		turn := whoGoesfirst()
		if turn == 0 {
			fmt.Println("Player 1 goes first")
		} else {
			if mode == 1 {
				fmt.Println("Computer goes first")
			}
		}

		board = []string{" ", " ", " ", " ", " ", " ", " ", " ", " "}
		if mode == 1 && turn == 0 {
			drawBoard(board)
		}

		for {
			if mode == 0 {
				// player1
				if turn == 0 {
					fmt.Println("Player 1's turn")
					move := getMove(board)
					board[move] = fmt.Sprintf("%s", xColor(x))
					turn = 1
					drawBoard(board)
					fmt.Println("")
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
					fmt.Println("")
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
			} else {
				// player1
				if turn == 0 {
					fmt.Println("Player 1's turn")
					move := getMove(board)
					board[move] = fmt.Sprintf("%s", xColor(x))
					turn = 1
					drawBoard(board)
					fmt.Println("")
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
					// computer
					turn = 0
					fmt.Println("Computers 2's turn")
					move, err := getMoveComputer(board)
					if err != nil {
						drawBoard(board)
						fmt.Println("Game is a tie")
						break
					}

					board[move] = fmt.Sprintf("%s", oColor(o))
					fmt.Println("")
					win, err := checkBoard(board, o)
					if err != nil {
						fmt.Println("Game is a tie")
						break
					}
					drawBoard(board)
					if win {
						fmt.Println("Computer Wins")
						break
					}
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
