package main

import "fmt"

func main() {
	/* Declare variables */
	var keyword int
	var history []string

	var number1, number2 float64
	var operator string
	// var status bool

	var i int
	/* End of declaring variables */

	fmt.Println("============================## Mini Calculator Apps ##============================")
	menu()               // Call menu function
	fmt.Scanln(&keyword) // Choose Menu
menus:
	for {
		switch keyword {
		case 1:
			fmt.Printf("\n")
			fmt.Println("============================ ## CALCULATOR ## ============================")
			i = 0 // initialize the state
		Calculator:
			for {
				if i == 0 { // First time when enter calculator menu do this
					fmt.Print("Please enter First number: ")
					fmt.Scanln(&number1)
					fmt.Print("Please enter Operator (+,-,/,*): ")
					fmt.Scanln(&operator)
					if operator == "#" { // if operator = # Go Back to Menu
						break Calculator
					}
					fmt.Print("Please enter Second number: ")
					fmt.Scanln(&number2)

				} else { // Second time and so on do this
					fmt.Print("Please enter Operator (+,-,/,%,*): ")
					fmt.Scanln(&operator)
					if operator == "#" { // if operator = # Go Back to Menu
						break Calculator
					}
					fmt.Print("Please enter Second number: ")
					fmt.Scanln(&number2)
				}

				// select the operator
				switch operator {
				// Why &number1? Because we want to change the value of number1 after done counting
				case "+":
					history = append(history, add(number1, number2, &number1)) // count and append to the slice
				case "-":
					history = append(history, subtracts(number1, number2, &number1)) // count and append to the slice
				case "*":
					history = append(history, multplies(number1, number2, &number1)) // count and append to the slice
				case "/":
					history = append(history, divides(number1, number2, &number1)) // count and append to the slice
				default:
					fmt.Println("Invalid Operation")
					continue // if none of the option above skip the process below and continue the loop
				}

				// Print Result of calculation
				fmt.Printf("Result : %.1f\n", number1)
				i++
			}
		case 2:
			fmt.Printf("\n")
			fmt.Println("============================ ## HISTORY ## ============================")
			showHistories(history) // show the history of calculator from slice history
		case 3:
			fmt.Printf("\n")
			fmt.Println("Thankyou for using me! ^_^")
			fmt.Println("============================## END ##============================")
			break menus
		default:
			// if none of the option above print this and Call the menu func back below
			fmt.Println("Wrong keyword! Please check the Menu above! Your keyword :", keyword)
		}
		fmt.Printf("\n")
		menu() // Call menu function
		fmt.Scanln(&keyword)
	}

}

/* Menu function */
func menu() {
	fmt.Println("Menu :")
	fmt.Println("=== 1. Calculator")
	fmt.Println("=== 2. History")
	fmt.Println("=== 3. Exit")
	fmt.Println("### Make sure that your input is a number")
	fmt.Print("What are you looking for? : ")
}

/* Function to show the histories of calculator*/
func showHistories(history []string) {
	if len(history) == 0 {
		fmt.Println("There's no history here!")
	} else {
		for a := 0; a < len(history); a++ {
			fmt.Printf("%s\n", history[a])
		}
	}
}

/* Function add operator */
func add(number1, number2 float64, total *float64) (history string) {
	*total = number1 + number2
	history = histories(number1, number2, total, " + ")
	return history
}

/* Function difference operator */
func subtracts(number1, number2 float64, total *float64) (history string) {
	*total = number1 - number2
	history = histories(number1, number2, total, " - ")
	return history
}

/* Function multiply operator */
func multplies(number1, number2 float64, total *float64) (history string) {
	*total = number1 * number2
	history = histories(number1, number2, total, " * ")
	return history

}

/* Function divide operator */
func divides(number1, number2 float64, total *float64) (history string) {
	*total = number1 / number2
	history = histories(number1, number2, total, " / ")
	return history
}

/* Function to concat the history of calculator */
func histories(number1, number2 float64, total *float64, operator string) (history string) {
	history = fmt.Sprintf("%.1f", number1) + operator + fmt.Sprintf("%.1f", number2)
	history = history + " = " + fmt.Sprintf("%.1f", *total)
	return
}
