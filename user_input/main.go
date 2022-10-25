package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main(){
	fmt.Print("Enter Something: ")
	date, err := bufio.NewReader(os.Stdin).ReadString('\n')
	date = strings.TrimRight(date, "\n")
	if err != nil {
		fmt.Printf("Error is %v\n", err)
		return
	} 
	fmt.Println("Your input: -> ",date)
}