package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	startGame()
}

func startGame() {
	var playerInput string
	player := "X"
	var gameBoard = [][]string{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	}

	fmt.Println("Input a number in the format #row,#colmun to play. E.g 0,1 to mark row 1 column2. ")
	fmt.Println("Enter 'exit' to stop playing.")
	fmt.Println("==============================")
	for {
		fmt.Print("Player ", player, "'s move: ")
		fmt.Scanln(&playerInput)
		fmt.Println()
		if playerInput == "exit" {
			checkForWinner(gameBoard)
			break
		}

		err := addToBoard(player, playerInput, &gameBoard)
		if err != nil {
			fmt.Println(err)
			continue
		}

		switchPlayer(&player)

		hasWinner := checkForWinner(gameBoard)
		if hasWinner {
			break
		}

		if isBoardFilled(gameBoard) {
			break
		}

		printBoard(gameBoard)
	}

	fmt.Println("==============================")
	fmt.Println("Final Board")
	fmt.Println("==============================")
	printBoard(gameBoard)
}

func printBoard(gameBoard [][]string) {
	for index := range gameBoard {
		fmt.Printf(" %s | %s | %s \n", gameBoard[index][0], gameBoard[index][1], gameBoard[index][2])
	}
}

func isBoardFilled(gameBoard [][]string) bool {
	var i int
	result := true
	for i = 0; i < len(gameBoard); i++ {
		for index := range gameBoard[i] {
			if gameBoard[i][index] == "-" {
				result = false
				break
			}
		}
	}
	return result
}

func addToBoard(player string, playerInput string, gameBoard *[][]string) error {
	coordinates := strings.Split(playerInput, ",")

	if len(coordinates) != 2 {
		return errors.New("Invalid input format")
	}

	row, _ := strconv.Atoi(coordinates[0])
	col, _ := strconv.Atoi(coordinates[1])

	if row > 2 || col > 2 {
		return errors.New("Invalid coordinates entered")
	}

	if (*gameBoard)[row][col] != "-" {
		return errors.New("Move already made by player " + (*gameBoard)[row][col] + " 🥺")
	}

	(*gameBoard)[row][col] = player
	return nil
}

func checkForWinner(gameBoard [][]string) bool {
	if hasWinner, winner := findWinner(gameBoard); hasWinner {
		fmt.Println()
		fmt.Println("==============================")
		fmt.Println("GAME ENDED")
		fmt.Println()
		fmt.Println("Player " + winner + " Won This Round!! 💃🏽")
		fmt.Println()
		fmt.Println("==============================")
		fmt.Println()
		return true
	}
	return false
}

func findWinner(board [][]string) (bool, string) {
	rowWin := false
	colWin := false
	diagonalWin := false
	var winner string
	for index := range board {
		if (board[index][0] != "-") && (board[index][0] == board[index][1]) && (board[index][0] == board[index][2]) {
			rowWin = true
			winner = board[index][0]
		} else if (board[0][index] != "-") && (board[0][index] == board[1][index]) && (board[0][index] == board[2][index]) {
			colWin = true
			winner = board[0][index]
		}
	}

	if (board[0][0] != "-") && (board[0][0] == board[1][1]) && (board[0][0] == board[2][2]) {
		diagonalWin = true
		winner = board[0][0]
	} else if (board[0][2] != "-") && (board[0][2] == board[1][1]) && (board[0][2] == board[2][0]) {
		diagonalWin = true
		winner = board[0][2]
	}

	return (rowWin || colWin || diagonalWin), winner
}

func switchPlayer(player *string) {
	if *player == "X" {
		*player = "O"
	} else {
		*player = "X"
	}
}
