package main

import (
	"fmt"
	"regexp"
)

type phoneBook struct {
	name        string
	phoneNumber string
	email       string
}

func main() {

	/* Declare variables */
	// Declare struct
	
	
	
	var phoneBookList = []phoneBook{}
	

	// Variables needed
	var keyword string
	var phone phoneBook
	var name string
	var number string
	var email string
	var index int
	// var status bool
	var search bool

	// var name, email, phoneNumber string

	/* End of declaring variables */

	fmt.Println("=========================== ## PhoneBook Apps ## ============================")
	menu()               // Call menu function
	fmt.Scanln(&keyword) // Choose Menu
	// fmt.Println("Masuk Main")

menus:
	for {
		switch keyword {
		case "1":
			fmt.Printf("\n")
			fmt.Println("============================ ## Add PhoneBook ## ============================")
			for {
				fmt.Print("Enter Name : ")
				fmt.Scanln(&phone.name)
				if len(phone.name) < 5 {
					fmt.Printf("\n")
					fmt.Println("Name of at least 5 Characters")
					fmt.Printf("\n")
				} else {
					break
				}
			}
			for {
				fmt.Printf("\n")
				fmt.Print("Enter Phone Number: ")
				fmt.Scanln(&phone.phoneNumber)
				if len(phone.phoneNumber) < 10 {
					fmt.Printf("\n")
					fmt.Println("Phone Number cannot be Less Than 10 Digits!")
					fmt.Println("Your input :" , phone.phoneNumber)
				} else if isNum(phone.phoneNumber) == false {
					fmt.Printf("\n")
					fmt.Println("Phone Number must be Numeric!")
					fmt.Println("Your input :" , phone.phoneNumber)
				} else {
					break
				}
			}

			for {
				fmt.Printf("\n")
				fmt.Print("Enter Email: ")
				fmt.Scanln(&phone.email)
				if isEmail(phone.email) == false {
					fmt.Println("Email Must be an Email!")
					fmt.Println("Your Email :" , phone.email)
				} else {
					break
				}
			}
			phoneBookList = append(phoneBookList, phone)
			fmt.Printf("\n")
			fmt.Printf("Name :%s\nPhone Number : %s\nEmail :%s\n", phone.name, phone.phoneNumber, phone.email)
			fmt.Println("Successfully added to PhoneBook!")
			fmt.Printf("\n")
			fmt.Println("========================== ## Add PhoneBook End ## ==========================")
		case "2":
			fmt.Printf("\n")
			fmt.Println("========================= ## Update Phone Number ## =========================")
			for {
				fmt.Print("Enter Name : ")
				fmt.Scanln(&name)
				if len(name) < 5 {
					fmt.Println("Name of at least 5 Characters")
					fmt.Printf("\n")
				} else {
					break
				}
			}
			search = false
			// search in Phone Book List
			for i := 0; i < len(phoneBookList); i++ {

				if phoneBookList[i].name == name {
					search = true
					index = i
					email = phoneBookList[i].email
					break
				}
			}

			if search == true { 
				// status = true
				for {
					fmt.Printf("\n")
					fmt.Print("Enter Phone Number: ")
					fmt.Scanln(&number)
					fmt.Println(len(number))
					if len(number) < 10 {
						fmt.Printf("\n")
						fmt.Println("Phone Number cannot be Less Than 10 Digits!")
						fmt.Println("Your input :" , number)
						
					}else if isNum(number) == false {
						fmt.Printf("\n")
						fmt.Println("Phone Number must be Numeric!")
						fmt.Println("Your input :" , number)
					} else {
						break
					}
				}

				phoneBookList[index].phoneNumber = number
				fmt.Printf("\n")
				fmt.Printf("Name :%s\nPhone Number : %s\nEmail :%s\n", name, number, email)
				fmt.Println("Successfully updated Phone Number!")
				fmt.Printf("\n")
			} else {
				fmt.Printf("\n")
				fmt.Println("Error 404 Not Found!")
				fmt.Printf("\n")
			}
			fmt.Println("======================= ## Update Phone Number End ## =======================")
		case "3":
			fmt.Printf("\n")
			fmt.Println("========================= ## Show Phone Book List ## ========================")
			showPhoneBook(phoneBookList)
			fmt.Println("======================= ## Show Phone Book List End ## ======================")
		case "4":
			fmt.Printf("\n")
			fmt.Println("Successfully escape from PhoneBook Apps!")
			fmt.Println("================================= ## END ## =================================")
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

func menu() {
	fmt.Println("============================================================")
	fmt.Println("## Welcome to PhoneBook Apps, please choose a Menu below! ##")
	fmt.Println("## **[1].Input PhoneBook                                  ##")
	fmt.Println("## **[2].Update Phone Number                              ##")
	fmt.Println("## **[3].Show PhoneBook List                              ##")
	fmt.Println("## **[4].Escape from Apps                                 ##")
	fmt.Println("============================================================")
	fmt.Print("What are you looking for? : ")
}

func showPhoneBook(phoneBookList []phoneBook) {
	if len(phoneBookList) == 0 {
		fmt.Println("There's no history here!")
	} else {
		for a := 0; a < len(phoneBookList); a++ {
			fmt.Printf("%d. %s\n", a+1, phoneBookList[a].name)
			fmt.Printf("   %s\n", phoneBookList[a].phoneNumber)
			fmt.Printf("   %s\n", phoneBookList[a].email)
		}
	}
}

func isNum(phoneNumber string) bool {
	isStringNumeric,_ := regexp.MatchString(`^[\d]+$`, phoneNumber)
	return isStringNumeric
}

func isEmail(email string) bool {
	isStringEmail,_ := regexp.MatchString("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$", email)
	return isStringEmail
}