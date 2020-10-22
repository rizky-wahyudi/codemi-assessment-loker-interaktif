package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Loker struct {
	num      int
	itemType string
	itemId   int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var command string
	var commandSplit []string
	for {

		scanner.Scan()
		command = scanner.Text()

		if command != "" {

			commandSplit = strings.Split(command, " ")

			switch {

			//Init Loker
			case strings.EqualFold(commandSplit[0], "init"):
				fmt.Println("init function")
				break

				// Status
			case strings.EqualFold(commandSplit[0], "status"):
				fmt.Println("status function")
				break

				// Input
			case strings.EqualFold(commandSplit[0], "input"):
				fmt.Println("input function")
				break

				// Leave
			case strings.EqualFold(commandSplit[0], "leave"):
				fmt.Println("leave function")
				break

				// Find
			case strings.EqualFold(commandSplit[0], "find"):
				fmt.Println("find function")
				break

				// Search
			case strings.EqualFold(commandSplit[0], "search"):
				fmt.Println("search function")
				break

				// Exit
			case strings.EqualFold(commandSplit[0], "exit"):
				os.Exit(0)

				// Default: No command found
			default:
				fmt.Println("No command found with name", commandSplit[0])
			}
		}
	}
}
