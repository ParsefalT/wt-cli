package main

import (
	"fmt"
	"test/some"
)

type stringMap = map[int]string

func main() {
	some.Success("__Tasks__")
	tasks := stringMap{}
Menu:
	for {
		num := getMenu()
		switch num { 
		case 1:
			strline(45, "*")
			some.Info(tasks)
			strline(45, "*")
		case 2:
			addTask(tasks)
			some.Success("task is append")
		case 3:
			deleteTask(tasks)
			some.Delete("task is delete")
		case 4:
			some.Exit("You are exit")
			break Menu
		}
	}
}

func getMenu() int {
	var num int
	fmt.Println("1. get tasks")
	fmt.Println("2. add task")
	fmt.Println("3. del task")
	fmt.Println("4. exit")
	fmt.Scan(&num)
	return num
}

func addTask(task stringMap) {
	var newKey int
	var newString string
	fmt.Println("type key")
	fmt.Scan(&newKey)
	fmt.Println("type task")
	fmt.Scan(&newString)
	task[newKey] = newString
}

func deleteTask(task stringMap) {
	var val int
	fmt.Println("type number to delete")
	fmt.Scan(&val)
	delete(task, val)
}

func strline(val int, char string) {
	chars := ""
	for i := 0; i < val; i++ {
		chars += char
	}
	println(chars)
}
