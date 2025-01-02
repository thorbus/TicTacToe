package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

type Game struct {
	board      [9]string
	player     string
	turnNumber int
}

func main() {
	var game Game
	game.player = "O"

	gameOver := false
	var winner string

	for !gameOver {
		PrintBoard(game.board)
		move := askForMove()
		err := game.play(move)
		if err != nil {
			fmt.Printf("\033[31m%s\033[0m\n", err) // Red text for errors
			continue
		}

		gameOver, winner = CheckForWinner(game.board, game.turnNumber)
	}

	PrintBoard(game.board)
	if winner == "" {
		fmt.Println("\033[33mIt's a draw!\033[0m") // Yellow text for a draw
	} else {
		fmt.Printf("\033[32mCongratulations! %s is the winner!\033[0m\n", winner) // Green text for the winner
	}
}

func CheckForWinner(b [9]string, n int) (bool, string) {
	// Horizontal test
	for i := 0; i < 9; i += 3 {
		if b[i] == b[i+1] && b[i+1] == b[i+2] && b[i] != "" {
			return true, b[i]
		}
	}

	// Vertical test
	for i := 0; i < 3; i++ {
		if b[i] == b[i+3] && b[i+3] == b[i+6] && b[i] != "" {
			return true, b[i]
		}
	}

	// Diagonal tests
	if b[0] == b[4] && b[4] == b[8] && b[0] != "" {
		return true, b[0]
	}
	if b[2] == b[4] && b[4] == b[6] && b[2] != "" {
		return true, b[2]
	}

	// Draw condition
	if n == 9 {
		return true, ""
	}
	return false, ""
}

func (game *Game) play(pos int) error {
	if pos < 1 || pos > 9 {
		return errors.New("invalid position, please choose between 1 and 9")
	}
	if game.board[pos-1] == "" {
		game.board[pos-1] = game.player
		game.SwitchPlayers()
		game.turnNumber++
		return nil
	}
	return errors.New("position already taken, try another")
}

func (game *Game) SwitchPlayers() {
	if game.player == "O" {
		game.player = "X"
	} else {
		game.player = "O"
	}
}

func askForMove() int {
	var move int
	fmt.Print("\033[36mEnter your move (1-9): \033[0m") // Cyan text for input prompt
	fmt.Scan(&move)
	return move
}

func PrintBoard(b [9]string) {
	ClearScreen()

	fmt.Println("\033[34mTic-Tac-Toe\033[0m") // Blue text for the title
	for i, v := range b {
		cell := v
		if cell == "" {
			cell = " "
		} else if cell == "O" {
			cell = "\033[35mO\033[0m" // Magenta "O"
		} else {
			cell = "\033[33mX\033[0m" // Yellow "X"
		}

		fmt.Printf(" %s ", cell)
		if (i+1)%3 == 0 {
			fmt.Println()
			if i < 8 {
				fmt.Println("---|---|---")
			}
		} else {
			fmt.Print("|")
		}
	}
	fmt.Println()
}

func ClearScreen() {
	c := exec.Command("clear") // For Unix/Linux/MacOS
	if os.PathSeparator == '\\' {
		c = exec.Command("cmd", "/c", "cls") // For Windows
	}
	c.Stdout = os.Stdout
	c.Run()
}
