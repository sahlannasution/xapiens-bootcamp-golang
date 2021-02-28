package main

import "fmt"

func main() {
	// template header

	var keyword int
	var history []string

	var number1, number2 float64
	var operator string
	var status bool

	var i int
	fmt.Println("============================## Mini Calculator Apps ##============================")
	menu()
	fmt.Scanln(&keyword)
menus:
	for keyword > 0 {
		switch keyword {
		case 1:
			// var total float64
			status = true
			i = 0
			fmt.Printf("\n")
			fmt.Println("============================ ## CALCULATOR ## ============================")
		Calculator:
			for status != false {
				if i == 0 {
					fmt.Print("Please enter First number: ")
					fmt.Scanln(&number1)
					fmt.Print("Please enter Operator (+,-,/,*): ")
					fmt.Scanln(&operator)
					if operator == "#" {
						break Calculator
					}
					fmt.Print("Please enter Second number: ")
					fmt.Scanln(&number2)

				} else {
					fmt.Print("Please enter Operator (+,-,/,%,*): ")
					fmt.Scanln(&operator)
					if operator == "#" {
						break Calculator
					}
					fmt.Print("Please enter Second number: ")
					fmt.Scanln(&number2)
				}

				switch operator {
				case "+":
					history = append(history, add(number1, number2, &number1))
				case "-":
					history = append(history, subtracts(number1, number2, &number1))
				case "*":
					history = append(history, multplies(number1, number2, &number1))
				case "/":
					history = append(history, divides(number1, number2, &number1))
				default:
					fmt.Println("Invalid Operation")
					// break Calculator
					continue
				}
				i++
				// history = append(history, total)

				fmt.Printf("Result : %.1f\n", number1)
				// number1 = total
			}
		case 2:
			fmt.Printf("\n")
			fmt.Println("============================ ## HISTORY ## ============================")
			showHistories(history)
		case 3:
			fmt.Printf("\n")
			fmt.Println("Thankyou for using me! ^_^")
			fmt.Println("============================## END ##============================")
			break menus
		default:
			fmt.Println("Wrong keyword! Please check the Menu above! Your keyword :", keyword)
		}
		fmt.Printf("\n")
		menu()
		fmt.Scanln(&keyword)
	}

}

func menu() {
	fmt.Println("Menu :")
	fmt.Println("=== 1. Calculator")
	fmt.Println("=== 2. History")
	fmt.Println("=== 3. Exit")
	fmt.Println("### Make sure that your input is a number")
	fmt.Print("What are you looking for? : ")
}

func showHistories(history []string) {
	if len(history) == 0 {
		fmt.Println("There's no history here!")
	} else {
		for a := 0; a < len(history); a++ {
			fmt.Printf("%s\n", history[a])
		}
	}

}

func add(number1, number2 float64, total *float64) (history string) {
	*total = number1 + number2
	history = histories(number1, number2, total, " + ")
	return history
}

func subtracts(number1, number2 float64, total *float64) (history string) {
	*total = number1 - number2
	history = histories(number1, number2, total, " - ")
	return history
}

func multplies(number1, number2 float64, total *float64) (history string) {
	*total = number1 * number2
	history = histories(number1, number2, total, " * ")
	return history

}

func divides(number1, number2 float64, total *float64) (history string) {
	*total = number1 / number2
	history = histories(number1, number2, total, " / ")
	return history
}

func histories(number1, number2 float64, total *float64, operator string) (history string) {
	history = fmt.Sprintf("%.1f", number1) + operator + fmt.Sprintf("%.1f", number2)
	history = history + " = " + fmt.Sprintf("%.1f", *total)
	return
}
