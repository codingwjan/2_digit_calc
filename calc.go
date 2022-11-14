package main

import "fmt"

func main() {

	//infinite loop
	for {

		//welcome message
		fmt.Println("Welcome to the calculator program!")

		//input first number
		fmt.Println("Please enter the first number:")
		var firstNumber float64
		fmt.Scanln(&firstNumber)
		//check if input is integer
		if firstNumber != float64(int(firstNumber)) {
			fmt.Println("Please enter an integer!")
			continue
		}

		//input second number
		fmt.Println("Please enter the second number:")
		var secondNumber float64
		fmt.Scanln(&secondNumber)
		if secondNumber != float64(int(secondNumber)) {
			fmt.Println("Please enter an integer!")
			continue
		}

		//input operation
		fmt.Println("Please enter the operation to perform (+, -, *, /):")
		var operation string
		fmt.Scanln(&operation)
		if operation != string(string(operation)) {
			fmt.Println("Please enter a string!")
			continue
		}

		//switch statement
		switch operation {
		case "+":
			fmt.Println("The result is:", firstNumber+secondNumber)
		case "-":
			fmt.Println("The result is:", firstNumber-secondNumber)
		case "*":
			fmt.Println("The result is:", firstNumber*secondNumber)
		case "/":
			fmt.Println("The result is:", firstNumber/secondNumber)
		default:
			fmt.Println("Invalid operation!")
		}

	}
}