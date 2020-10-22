package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Loker struct {
	num      int
	itemType string
	itemId   int
}

// Initialize Loker
func initializeLoker(command []string, loker *[]Loker) (bool, error) {
	if len(command) == 2 {
		size, err := strconv.Atoi(command[1])
		if err != nil {
			return false, err
		} else {
			if size >= 1 {
				*loker = make([]Loker, size)
				return true, nil
			} else {
				return false, errors.New("Cannot initialize loker less than 1")
			}
		}
	} else if len(command) > 2 {
		return false, fmt.Errorf("Too much argument, expected 1, have %d", len(command)-1)
	} else {
		return false, fmt.Errorf("Too few argument, expected 1, have %d", len(command)-1)
	}
}

// Check if loker initialized or not
func isLokerInitialized(loker []Loker) bool {
	if len(loker) > 0 {
		return true
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var loker []Loker
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
				res, err := initializeLoker(commandSplit, &loker)
				if res {
					fmt.Println("Loker initialized with", len(loker), "slot(s)")
				} else {
					fmt.Println("Error:", err)
				}
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
