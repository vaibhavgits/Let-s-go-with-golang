package main

import "fmt"

func main() {
	var username string = "vaibhav"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	//uint can only take value from 0-255
	// var smallVal2 uint8 = 256
	// fmt.Println(smallVal2)
	// fmt.Printf("Variable is of type: %T \n", smallVal2)

	//implicit way, lexer helps in this. You need not to decalre the variable type
	var website = "learncodeonline.in"
	fmt.Println(website)

	//now if you put website=3, it will genenrate an error.

	//no var style, walrus operator(:=)
	numberOfUsers := 300000
	fmt.Println(numberOfUsers)

	//const
	//Public variable
	const LoginToken string = "jbsafkjbjsfs"
	fmt.Println(LoginToken)

}
