package main

import "fmt"

// constants used through the game
const DIM = 9

type board struct {

	grid [9]string
	player1 player
	player2 player

	activePlayer *player
	waitingPlayer *player

	winner *player

	movesLeft int
}


func NewBoard() *board {
	b := board{movesLeft: 9}
	for i:=0; i<DIM; i++ {
		b.grid[i] = " "
	}

	b.setUpPlayers()

	return &b
}
func (b *board) printBoard(){

	for i:=0;i<3;i++{

		fmt.Println("+-----+-----+-----+")
		fmt.Println("|     |     |     |")
		fmt.Printf("|  %s  |  %s  |  %s  |\n", b.grid[0+i*3],b.grid[1+i*3],b.grid[2+i*3])
		fmt.Println("|     |     |     |")
	}
	fmt.Println("+-----+-----+-----+")

}

func (b *board) setUpPlayers(){
	
	b.player1 = player{name: "Player 1"}
	b.player2 = player{name: "Player 2"}

	b.player1.getSymbol()
	b.player2.getSymbol()

	b.activePlayer, b.waitingPlayer = &b.player1, &b.player2
	
}

func (b *board) switchTurn() {
	b.activePlayer, b.waitingPlayer = b.waitingPlayer, b.activePlayer
}

func (b *board) playTurn(){

	validTurn := false

	for !validTurn {
		fmt.Printf("[%s] Choose where to put your mark 1-9:\n", b.activePlayer.name)

		var position int

		_, error := fmt.Scan(&position)

		// computer readable position
		position -=1

		if error != nil || (position < 0 || position > 8) {
			fmt.Printf("[%s] Choose a number between 1 and 9\n", b.activePlayer.name)
			continue
		}
		// check if the cell is already full
		if b.grid[position] != " " {
			fmt.Printf("[%s] Cell already full. Choose another one\n", b.activePlayer.name)
			continue
		}
		// regular play	
		b.grid[position] = b.activePlayer.symbol
		validTurn = true
		b.movesLeft-=1
	}
}

func (b *board) setWinner(symbol string){

	if symbol == b.player1.symbol{
		b.winner = &b.player1
	}else{
		b.winner = &b.player2
	}
}

func (b *board) checkWin() bool{

	//check rows
	for i:=0; i<3;  i++ {
		if b.grid[0+i*3] != " " && b.grid[0+i*3] == b.grid[1+i*3] && b.grid[1+i*3] == b.grid[2+i*3]{
			b.setWinner(b.grid[0+i*3])
			return true
		}
	}

	//check columns
	for i:=0; i<3;  i++ {
		if b.grid[i] != " " && b.grid[i] == b.grid[i+3] && b.grid[i+3] == b.grid[i+6]{
			b.setWinner(b.grid[i])
			return true
		}
	}

	//check diagonals
	if b.grid[0] != " " && b.grid[0] == b.grid[4] && b.grid[4] == b.grid[8]{
		b.setWinner(b.grid[0])
		return true
	}

	if b.grid[2] != " " && b.grid[2] == b.grid[4] && b.grid[4] == b.grid[6]{
		b.setWinner(b.grid[2])
		return true
	}

	return false
}

type player struct{
	name string
	symbol string
}

func (p *player) getSymbol(){

	symbolAllowed := false

	for !symbolAllowed {
		fmt.Printf("Enter a character symbol for %s:", p.name)

		fmt.Scan(&p.symbol)

		if(len(p.symbol)>1){
			p.symbol = ""
			fmt.Println("The length of the symbol must be 1")
		}else{
			symbolAllowed = true
		}
	}
}

func main(){

	
	tictactoe := NewBoard()

	for !tictactoe.checkWin() && tictactoe.movesLeft>0{

		tictactoe.printBoard()
		tictactoe.playTurn()
		tictactoe.switchTurn()
	}

	tictactoe.printBoard()
	// check if there is a winner
	if tictactoe.winner != nil {
		fmt.Printf("The winner is %s\n", tictactoe.winner.name)
	}else{
		fmt.Println("The game is a draw")
	}
	

}