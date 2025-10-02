package main

import "fmt"

func main() {
	fmt.Println("###### Welcome to our Todolist App! ######")

	var shortGolang = "Watch Go crash course"
	var fullGolang = "Watch Nana's Golang Full Course"
	var rewardDessert = "Reward myself with a donut"
	var taskItems = []string{shortGolang, fullGolang, rewardDessert}

	printTasks(taskItems)
	fmt.Println()

	taskItems = addTask(taskItems, "Go for a run")
	taskItems = addTask(taskItems, "Practice coding in Go")

	fmt.Println("Updated list")
	printTasks(taskItems)

	result := sum1(1, 2, 3)
	fmt.Println((result))
}

func printTasks(taskItems []string) {
	fmt.Println("List of my Todos")
	for index, task := range taskItems {
		fmt.Printf("%d: %s\n", index+1, task)
	}
}

func addTask(taskItems []string, newTask string) []string {
	var updatedTaskItems = append(taskItems, newTask)
	return updatedTaskItems
}

func sum1(num1 int, num2 int, num3 int) int {
	return num1 + num2 + num3
}
