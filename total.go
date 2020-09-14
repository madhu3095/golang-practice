package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 && args[0] == "/help" {
		fmt.Println("Usage: total <Total meal Amount><tip Percentage>")
		fmt.Println("Example: total 20 10")
	} else {
		if len(args) != 2 {
			fmt.Println("you must enter 2 arguments! type /help for help")
		} else {

			mealTotal, _ := strconv.ParseFloat(args[0], 32)
			tipAmount, _ := strconv.ParseFloat(args[0], 32)
			fmt.Printf("\nyour meal total will be %.2f\n", calculateTotal(float32(mealTotal), float32(tipAmount)))
		}
	}

}
func calculateTotal(mealTotal float32, tipAmount float32) float32{
	totalPrice :=mealTotal + (mealTotal * (tipAmount /100))
	return totalPrice
}
