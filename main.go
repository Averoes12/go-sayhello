package main

import (
	"fmt"

	"github.com/Averoes12/go-sayhello/v2/myfunction"
)

func main() {

	var firstName, Lastname, Addrees string

	fmt.Println("Welcome to the hacktober fest project")
	fmt.Println("========================================")
	fmt.Print("input First name = ")
	fmt.Scanln(&firstName)
	fmt.Print("input LastName = ")
	fmt.Scanln(&Lastname)
	fmt.Print("input Address = ")
	fmt.Scanln(&Addrees)
	fmt.Println("")
	fmt.Println("Welcome")
	fmt.Println(myfunction.SayHello(firstName, Lastname))
	fmt.Println(myfunction.Addrees(Addrees))

	fmt.Print("Press Any key to Exit  ")
	fmt.Scanln()

}
