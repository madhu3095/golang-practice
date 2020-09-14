package main

import (
	"runtime"
	"fmt"
	"os"
	"bufio"
)

func main() {
	//this is where stuff happens
	reader :=bufio.NewReader(os.Stdin)
	fmt.Println("what is your name?")
	text, _ := reader.ReadString('\n')
	fmt.Printf("hello %v", text)
	fmt.Printf(
		"we are using go %v running in %v\n",
		runtime.Version(),
		runtime.GOOS,
	)
}
