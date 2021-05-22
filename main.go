package main

import "fmt"

// constants used through the game
const DIM = 9

func main(){

	
	game := NewTicTacToe()

	for !game.checkWin() && game.movesLeft>0{

		game.printBoard()
		game.playTurn()
		game.switchTurn()
	}

	game.printBoard()
	// check if there is a winner
	if game.winner != nil {
		fmt.Printf("The winner is %s\n", game.winner.name)
	}else{
		fmt.Println("The game is a draw")
	}
	

}