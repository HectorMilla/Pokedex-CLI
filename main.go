package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	for i :=1; i >0; i++{
		fmt.Printf("Pokedex >")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan(){
			result := CleanInput(scanner.Text())[0]
			if _, ok := commands[result]; ok{
				commands[result].callback()
			} else {
				fmt.Print("Unknown command")
			}
		}
		
	}

}

