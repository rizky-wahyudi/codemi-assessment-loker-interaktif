package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
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

func findEmptyLoker(loker []Loker) int {
	for i := 0; i < len(loker); i++ {
		if reflect.ValueOf(loker[i]).IsZero() {
			return i
		}
	}
	return -1
}

func isLokerEmpty(loker []Loker) bool {
	for i := 0; i < len(loker); i++ {
		if !reflect.ValueOf(loker[i]).IsZero() {
			return false
		}
	}
	return true
}

// Input Loker
func inputLoker(command []string, loker *[]Loker) (bool, int, error) {
	if len(command) == 3 {
		emptyLoker := findEmptyLoker(*loker)
		id, err := strconv.Atoi(command[2])
		if err != nil {
			return false, 0, err
		}
		if emptyLoker != -1 {
			_loker := Loker{num: emptyLoker + 1, itemType: command[1], itemId: id}
			(*loker)[emptyLoker] = _loker
			return true, emptyLoker + 1, nil
		} else {
			return false, 0, errors.New("Locker Full")
		}
	} else if len(command) > 3 {
		return false, 0, fmt.Errorf("Too much argument, expected 2, have %d", len(command)-1)
	} else {
		return false, 0, fmt.Errorf("Too few argument, expected 2, have %d", len(command)-1)
	}
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
				if !isLokerInitialized(loker) {
					fmt.Println("You haven't initialized the Locker yet")
				} else {
					if isLokerEmpty(loker) {
						fmt.Println("Loker Empty")
					} else {
						fmt.Printf("%-10s %-10s %-10s\n", "Loker No.", "ID Type", "ID Number")
						for i := 0; i < len(loker); i++ {
							if !reflect.ValueOf(loker[i]).IsZero() {
								fmt.Printf("%-10d %-10s %-10d\n", loker[i].num, loker[i].itemType, loker[i].itemId)
							}
						}
					}
				}
				break

			// Input
			case strings.EqualFold(commandSplit[0], "input"):
				if !isLokerInitialized(loker) {
					fmt.Println("You haven't initialized the Locker yet")
				} else {
					res, info, err := inputLoker(commandSplit, &loker)
					if res {
						fmt.Println("Id Card stored in loker", info)
					} else {
						fmt.Println("Error:", err)
					}
				}
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
