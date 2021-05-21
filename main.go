package main

import "fmt"

func printBoard(board [9]string){

	for i:=0;i<3;i++{

		fmt.Println("+-----+-----+-----+")
		fmt.Println("|     |     |     |")
		fmt.Printf("|  %s  |  %s  |  %s  |\n", board[0+i*3],board[1+i*3],board[2+i*3])
		fmt.Println("|     |     |     |")
	}
	fmt.Println("+-----+-----+-----+")

}

func checkWin(board [9] string) bool{

	//check rows
	for i:=0; i<3;  i++ {
		if board[0+i*3] != " " && board[0+i*3] == board[1+i*3] && board[1+i*3] == board[2+i*3]{
			return true
		}
	}

	//check columns
	for i:=0; i<3;  i++ {
		if board[i] != " " && board[i] == board[i+3] && board[i+3] == board[i+6]{
			return true
		}
	}

	//check diagonals
	if board[0] != " " && board[0] == board[4] && board[4] == board[8]{
		return true
	}

	if board[2] != " " && board[2] == board[4] && board[4] == board[6]{
		return true
	}

	return false
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

	const DIM = 9
	var board[DIM] string

	for i:=0;i<DIM;i++{
		board[i] = " "
	}

	activePlayer := &player1
	leftMoves := 9

	for !checkWin(board) && leftMoves>0{

		printBoard(board)
		fmt.Println("Choose where to put your mark 1-9:")

		var position int

		_, error := fmt.Scan(&position)

		if error != nil || (position < 1 || position > 9) {
			fmt.Println("You lost your turn!")
		}else{
			if board[position-1] == " "{
				board[position-1] = activePlayer.symbol
			}
		}
		// check who need to play next

		if activePlayer == &player1 {
			activePlayer = &player2
		}else {
			activePlayer = &player1
		}

		leftMoves-=1
	}

	printBoard(board)
	

}