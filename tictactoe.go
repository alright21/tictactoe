package main

import "fmt"

type tictactoe struct {

	grid [9]string
	player1 player
	player2 player

	activePlayer *player
	waitingPlayer *player

	winner *player

	movesLeft int
}


func NewTicTacToe() *tictactoe {
	t := tictactoe{movesLeft: 9}
	for i:=0; i<9; i++ {
		t.grid[i] = " "
	}

	t.setUpPlayers()

	return &t
}
func (t *tictactoe) printBoard(){

	for i:=0;i<3;i++{

		fmt.Println("+-----+-----+-----+")
		fmt.Println("|     |     |     |")
		fmt.Printf("|  %s  |  %s  |  %s  |\n", t.grid[0+i*3],t.grid[1+i*3],t.grid[2+i*3])
		fmt.Println("|     |     |     |")
	}
	fmt.Println("+-----+-----+-----+")

}

func (t *tictactoe) setUpPlayers(){
	
	t.player1 = player{name: "Player 1", symbol: "X"}
	t.player2 = player{name: "Player 2", symbol: "O"}

	t.player1.GetSymbol()
	t.player2.GetSymbol()

	t.activePlayer, t.waitingPlayer = &t.player1, &t.player2
	
}

func (t *tictactoe) switchTurn() {
	t.activePlayer, t.waitingPlayer = t.waitingPlayer, t.activePlayer
}

func (t *tictactoe) playTurn(){

	validTurn := false

	for !validTurn {
		fmt.Printf("[%s] Choose where to put your mark 1-9:\n", t.activePlayer.name)

		var position int

		_, error := fmt.Scan(&position)

		// computer readable position
		position -=1

		if error != nil || (position < 0 || position > 8) {
			fmt.Printf("[%s] Choose a number between 1 and 9\n", t.activePlayer.name)
			continue
		}
		// check if the cell is already full
		if t.grid[position] != " " {
			fmt.Printf("[%s] Cell already full. Choose another one\n", t.activePlayer.name)
			continue
		}
		// regular play	
		t.grid[position] = t.activePlayer.symbol
		validTurn = true
		t.movesLeft-=1
	}
}

func (t *tictactoe) setWinner(symbol string){

	if symbol == t.player1.symbol{
		t.winner = &t.player1
	}else{
		t.winner = &t.player2
	}
}

func (t *tictactoe) checkWin() bool{

	//check rows
	for i:=0; i<3;  i++ {
		if t.grid[0+i*3] != " " && t.grid[0+i*3] == t.grid[1+i*3] && t.grid[1+i*3] == t.grid[2+i*3]{
			t.setWinner(t.grid[0+i*3])
			return true
		}
	}

	//check columns
	for i:=0; i<3;  i++ {
		if t.grid[i] != " " && t.grid[i] == t.grid[i+3] && t.grid[i+3] == t.grid[i+6]{
			t.setWinner(t.grid[i])
			return true
		}
	}

	//check diagonals
	if t.grid[0] != " " && t.grid[0] == t.grid[4] && t.grid[4] == t.grid[8]{
		t.setWinner(t.grid[0])
		return true
	}

	if t.grid[2] != " " && t.grid[2] == t.grid[4] && t.grid[4] == t.grid[6]{
		t.setWinner(t.grid[2])
		return true
	}

	return false
}