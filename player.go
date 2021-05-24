package main

import "fmt"

type player struct{
	name string
	symbol string
}

func (p *player) GetSymbol(){

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