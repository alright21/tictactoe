package main

import "fmt"

func printBoard(){

}

func checkWin(){

}

type player struct{
	symbol string
}

func main(){

	player1 := player{}
	player2 := player{}


	// wrap everything up in a function
	fmt.Println("Enter a character for player 1:")

	fmt.Scan(&player1.symbol)

	fmt.Println("Enter a character for player 2:")

	fmt.Scan(&player2.symbol)
	

}