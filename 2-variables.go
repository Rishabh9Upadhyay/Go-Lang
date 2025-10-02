package main

import "fmt"

func main() {
	var shortGolang = "Watch Go crash course"
	var fullGolang = "Watch Nana's Golang Full Course"
	rewardDessert := "Reward myself with a donut"

	fmt.Println("###### Welcome to our Todolist App! ######")

	fmt.Println("List of my Todos")
	fmt.Println(shortGolang)
	fmt.Println(fullGolang)
	fmt.Println(rewardDessert)

	fmt.Println()
	fmt.Println("Tutorials")
	fmt.Println(shortGolang)
	fmt.Println(fullGolang)

	fmt.Println()
	fmt.Println("Rewards")
	fmt.Println(rewardDessert)

	fmt.Println()
	fmt.Println("My Project")
	fmt.Println(fullGolang)

	var num1 = 100
	var bool1 = true
	fmt.Println(num1)
	fmt.Println(bool1)

	if bool1 {
		fmt.Println("Rishabh Upadhyay")
	} else {
		fmt.Println("Risabh fix issue")
	}
}
