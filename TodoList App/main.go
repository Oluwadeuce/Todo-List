package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Todo struct {
	Title string
	Done  bool
}

func main() {
	var todos []Todo
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n  To-DO LIST ")
		fmt.Println("1. Add To-Do")
		fmt.Println("2. View To-Dos")
		fmt.Println("3. Mark To-Dos as Done")
		fmt.Println("4. Delete To-Dos")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		input, _ := reader.ReadString('\n')
		choice, _ := strconv.Atoi(strings.TrimSpace(input))

		switch choice {
		case 1:
			fmt.Print("Enter To-Do title: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)

			todos = append(todos, Todo{Title: title})
			fmt.Println("To-do added!")

		case 2:
			if len(todos) == 0 {
				fmt.Println("No To-Dos added yet.")
				continue
			}

			fmt.Println("\nTo-Dos:")
			for i, todo := range todos {
				status := " "
				if todo.Done {
					status = "✓"
				}
				fmt.Printf("%d. %s %s\n", i+1, status, todo.Title)
			}

		case 3:
			fmt.Print("Enter To-Do number to mark as done: ")
			input, _ := reader.ReadString('\n')
			num, _ := strconv.Atoi(strings.TrimSpace(input))
			num-- 

			if num >= 0 && num <= len(todos) {
				todos[num-1].Done = true
				fmt.Println("Todo marked as done.")
			} else {
				fmt.Println("Invalid to-do number.")
			}

		case 4:
			fmt.Print("Enter To-Do number to delete: ")
			input, _ := reader.ReadString('\n')
			num, _ := strconv.Atoi(strings.TrimSpace(input))
			num-- 

			if num >= 0 && num <= len(todos) {
				todos = append(todos[:num], todos[num+1:]...)
				fmt.Println("To-do deleted.")
			} else {
				fmt.Println("Invalid To-Do number.")
			}

		case 5:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid option.")
		}
	}
}
