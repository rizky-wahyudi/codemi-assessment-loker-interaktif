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
				return false, errors.New("Cannot initialize Loker less than 1")
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

// Leave Loker
func leaveLoker(command []string, loker *[]Loker) (bool, int, error) {
	if len(command) == 2 {
		target, err := strconv.Atoi(command[1])
		if err != nil {
			return false, 0, err
		} else {
			if target >= 1 {
				if !reflect.ValueOf((*loker)[target-1]).IsZero() {
					(*loker)[target-1] = Loker{}
					return true, target, nil
				} else {
					return false, 0, errors.New("Loker already empty")
				}
			} else {
				return false, 0, errors.New("Loker number cannot less than 1")
			}

		}
	} else if len(command) > 2 {
		return false, 0, fmt.Errorf("Too much argument, expected 1, have %d", len(command)-1)
	} else {
		return false, 0, fmt.Errorf("Too few argument, expected 1, have %d", len(command)-1)
	}
}

func findById(loker []Loker, id int) int {
	for i := 0; i < len(loker); i++ {
		if loker[i].itemId == id && !reflect.ValueOf(loker[i]).IsZero() {
			return i
		}
	}
	return -1
}

// Find Loker
func findLoker(command []string, loker *[]Loker) (bool, int, error) {
	if len(command) == 2 {
		id, err := strconv.Atoi(command[1])
		if err != nil {
			return false, 0, err
		} else {
			res := findById(*loker, id)
			if res != -1 {
				return true, res + 1, nil
			} else {
				return false, 0, fmt.Errorf("Nothing match with ID %d", id)
			}

		}
	} else if len(command) > 2 {
		return false, 0, fmt.Errorf("Too much argument, expected 1, have %d", len(command)-1)
	} else {
		return false, 0, fmt.Errorf("Too few argument, expected 1, have %d", len(command)-1)
	}
}

func findByType(loker []Loker, idType string) []int {
	var idList []int
	for i := 0; i < len(loker); i++ {
		if strings.EqualFold(loker[i].itemType, idType) && !reflect.ValueOf(loker[i]).IsZero() {
			idList = append(idList, loker[i].itemId)
		}
	}
	return idList
}

// Search Loker
func searchLoker(command []string, loker *[]Loker) (bool, []int, error) {
	if len(command) == 2 {
		idType := command[1]
		res := findByType(*loker, idType)
		if !reflect.ValueOf(res).IsZero() {
			return true, res, nil
		} else {
			return false, nil, fmt.Errorf("Cannot found ID with type %s", idType)
		}
	} else if len(command) > 2 {
		return false, nil, fmt.Errorf("Too much argument, expected 1, have %d", len(command)-1)
	} else {
		return false, nil, fmt.Errorf("Too few argument, expected 1, have %d", len(command)-1)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var loker []Loker
	var command string
	var commandSplit []string

	fmt.Println("============================")
	fmt.Println("|    INTERACTIVE LOKER     |")
	fmt.Println("============================")

	for {
		fmt.Println("")
		fmt.Print(">> ")
		scanner.Scan()
		command = scanner.Text()

		if command != "" {

			commandSplit = strings.Split(command, " ")

			switch {

			//Init Loker
			case strings.EqualFold(commandSplit[0], "init"):
				res, err := initializeLoker(commandSplit, &loker)
				if res {
					fmt.Println(">>>> Loker initialized with", len(loker), "slot(s)")
				} else {
					fmt.Println("! Error:", err)
				}
				break

			// Status
			case strings.EqualFold(commandSplit[0], "status"):
				if !isLokerInitialized(loker) {
					fmt.Println("! You haven't initialized the Locker yet")
				} else {
					if isLokerEmpty(loker) {
						fmt.Println(">>>> Loker empty")
					} else {
						fmt.Println("|--------------------------------|")
						fmt.Printf("|%-10s %-10s %-10s|\n", "Loker No.", "ID Type", "ID Number")
						fmt.Println("|--------------------------------|")
						for i := 0; i < len(loker); i++ {
							if !reflect.ValueOf(loker[i]).IsZero() {
								fmt.Printf("|%-10d %-10s %-10d|\n", loker[i].num, loker[i].itemType, loker[i].itemId)
							}
						}
						fmt.Println("|--------------------------------|")
					}
				}
				break

			// Input
			case strings.EqualFold(commandSplit[0], "input"):
				if !isLokerInitialized(loker) {
					fmt.Println("! You haven't initialized the Locker yet")
				} else {
					res, info, err := inputLoker(commandSplit, &loker)
					if res {
						fmt.Println(">>>> Id Card stored in loker", info)
					} else {
						fmt.Println("! Error:", err)
					}
				}
				break

			// Leave
			case strings.EqualFold(commandSplit[0], "leave"):
				if !isLokerInitialized(loker) {
					fmt.Println("! You haven't initialized the Locker yet")
				} else {
					if isLokerEmpty(loker) {
						fmt.Println(">>>> Loker empty")
					} else {
						res, info, err := leaveLoker(commandSplit, &loker)
						if res {
							fmt.Println(">>>> Success empty Loker No.", info)
						} else {
							fmt.Println("! Error:", err)
						}
					}
				}
				break

			// Find
			case strings.EqualFold(commandSplit[0], "find"):
				if !isLokerInitialized(loker) {
					fmt.Println("! You haven't initialized the Locker yet")
				} else {
					if isLokerEmpty(loker) {
						fmt.Println(">>>> Loker empty")
					} else {
						res, info, err := findLoker(commandSplit, &loker)
						if res {
							fmt.Println(">>>> ID Number match in Loker No.", info)
						} else {
							fmt.Println("! Error:", err)
						}
					}
				}
				break

			// Search
			case strings.EqualFold(commandSplit[0], "search"):
				if !isLokerInitialized(loker) {
					fmt.Println("! You haven't initialized the Locker yet")
				} else {
					if isLokerEmpty(loker) {
						fmt.Println(">>>> Loker empty")
					} else {
						res, info, err := searchLoker(commandSplit, &loker)
						if res {
							fmt.Println(">>>> Found", len(info), "ID that match with the type: ", info)
						} else {
							fmt.Println("! Error:", err)
						}
					}
				}
				break

			case strings.EqualFold(commandSplit[0], "help"):
				fmt.Println("Command List")
				fmt.Printf("%-30s: %-30s\n", "Init <size>", "Initialize Loker size")
				fmt.Printf("%-30s: %-30s\n", "Status", "Display all non empty Loker")
				fmt.Printf("%-30s: %-30s\n", "Input <ID Type> <ID Number>", "Add ID Card to empty Loker")
				fmt.Printf("%-30s: %-30s\n", "Leave <Loker No.>", "Empty the Loker in specific No.")
				fmt.Printf("%-30s: %-30s\n", "Find <ID Number>", "Find Loker by ID Number")
				fmt.Printf("%-30s: %-30s\n", "Search <Id Type>", "Find every ID Number by ID Type")
				fmt.Printf("%-30s: %-30s\n", "Exit", "Exit Program")
				break

			// Exit
			case strings.EqualFold(commandSplit[0], "exit"):
				fmt.Println("Program exit...")
				os.Exit(0)

			// Default: No command found
			default:
				fmt.Println("No command found with name", commandSplit[0])
				fmt.Println("Try 'help' to see command list")
			}
		}
	}
}
