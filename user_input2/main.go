package main

import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	fmt.Print("Enter something: ")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	fmt.Println(line)
}