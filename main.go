package main

import "fmt"

func main() {
	var playerInput int
	player := "X"
	var gameBoard = []string{
		"-", "-", "-",
		"-", "-", "-",
		"-", "-", "-",
	}

	for {
		fmt.Println("Input a number between 1 - 9 to play: ")
		fmt.Scanln(&playerInput)
		if playerInput == 0 {
			break
		}
		addToBoard(player, playerInput, &gameBoard)
		if player == "X" {
			player = "O"
		} else {
			player = "X"
		}
		if isBoardFiled(&gameBoard) {
			break
		}

		printBoard(&gameBoard)
	}

	fmt.Println("==============================")
	fmt.Println("Final Board")
	fmt.Println("==============================")
	printBoard(gameBoard)
}

func printBoard(gameBoard []string) {
	fmt.Printf(" %s | %s | %s \n", gameBoard[0], gameBoard[1], gameBoard[2])
	fmt.Printf(" %s | %s | %s \n", gameBoard[3], gameBoard[4], gameBoard[5])
	fmt.Printf(" %s | %s | %s \n", gameBoard[6], gameBoard[7], gameBoard[8])
}

func isBoardFiled(gameBoard *[3]string) bool {
	var i int
	result := true
	for i = 0; i < 3; i++ {
		if gameBoard[i] == "-" {
			result = false
			break
		}
	}
	return result
}

func addToBoard(player string, playerInput int, gameBoard *[3]string) {
	gameBoard[playerInput] = player
	fmt.Println("Invalid number entered!!")

}

func checkForWinner(board []string) {
	winningCombos := [8][3]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}

	fmt.Print(winningCombos)
	for _, combo := range winningCombos {
		if (board[combo[0]] != "-") && (board[combo[0]] == board[combo[1]]) && (board[combo[0]] == board[combo[2]]) {
			// winner := board[combo[0]]
		}
	}
}
