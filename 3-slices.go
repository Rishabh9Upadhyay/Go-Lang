package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	var shortGolang = "Watch Go crash course"
	var fullGolang = "Watch Nana's Golang Full Course"
	var rewardDessert = "Reward myself with a donut"

	var taskItems = []string{shortGolang, fullGolang, rewardDessert}

	fmt.Println("###### Welcome to our Todolist App! ######")

	fmt.Println("List of my Todos")
	taskItems = append(taskItems, "Rishabh Upadhyay")
	fmt.Println("Tasks:", taskItems)

	slice := []int{1, 2, 3}
	fmt.Println(slice)
	slice = append(slice, 1, 2)
	fmt.Println(slice)
	slices.Reverse(slice)
	fmt.Println(slice)
	sort.Slice(slice, func(i, j int) bool {
		return i > j
	})
	fmt.Println(slice)

	array := [10]int{}
	fmt.Println(array)
	slice1 := array[:0] // slice backed by array, starting empty
	slice1 = append(slice1, 1, 2, 3)
	fmt.Println(array)
	fmt.Println(slice1)

	fmt.Println(array[0])
}
