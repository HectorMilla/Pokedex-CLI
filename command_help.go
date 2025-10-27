package main

import (
	"fmt"
)
func commandHelp(){
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n\n")
	for _, value := range commands{
		fmt.Printf("%s: %s\n",value.name,value.description)
	}
}
