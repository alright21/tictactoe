package main

import "fmt"

// Player represents a player in Tictactoe game
type Player struct{
	name string
	symbol string
}
// GetSymbol get the symbol for the tictactoe player
func (p *Player) GetSymbol(){

	symbolAllowed := false

	for !symbolAllowed {
		fmt.Printf("Enter a character symbol for %s:", p.name)

		fmt.Scanln(&p.symbol)

		if(len(p.symbol)>1){
			p.symbol = ""
			fmt.Println("The length of the symbol must be 1")
		}else{
			symbolAllowed = true
		}
	}
}