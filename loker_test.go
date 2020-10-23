package main

import (
	"reflect"
	"testing"
)

var (
	loker []Loker
)

func TestInitializeLoker1(t *testing.T) {
	command := []string{"init", "10"}
	res, err := initializeLoker(command, &loker)
	if res {
		expectation := 10
		actual := len(loker)
		if actual != expectation {
			t.Errorf("Expected %v but got %v", expectation, actual)
		}
	} else {
		t.Logf("Error Message : %v", err)
		t.Fail()
	}
}

func TestInitializeLoker2(t *testing.T) {
	command := []string{"init", "0"}
	res, err := initializeLoker(command, &loker)
	if res {
		t.Errorf("Failed, The program should go to else condition")
	} else {
		expectation := "Cannot initialize Loker less than 1"
		actual := err.Error()
		if actual != expectation {
			t.Errorf("Expected '%v' but got '%v'", expectation, actual)
		}
	}
}

func TestInitializeLoker3(t *testing.T) {
	command := []string{"init", "-2"}
	res, err := initializeLoker(command, &loker)
	if res {
		t.Errorf("Failed, The program should go to else condition")
	} else {
		expectation := "Cannot initialize Loker less than 1"
		actual := err.Error()
		if actual != expectation {
			t.Errorf("Expected '%v' but got '%v'", expectation, actual)
		}
	}
}

func TestInitializeLoker4(t *testing.T) {
	command := []string{"init", "loker"}
	res, err := initializeLoker(command, &loker)
	if res {
		t.Errorf("Failed, The program should go to else condition")
	} else {
		expectation := "strconv.Atoi: parsing \"loker\": invalid syntax"
		actual := err.Error()
		if actual != expectation {
			t.Errorf("Expected '%v' but got '%v'", expectation, actual)
		}
	}
}

func TestInitializeLoker5(t *testing.T) {
	loker := make([]Loker, 2)
	command := []string{"init"}
	res, err := initializeLoker(command, &loker)
	if res {
		t.Errorf("Loker should not initialized")
	} else {
		expectation := "Too few argument, expected 1, have 0"
		actual := err.Error()
		if actual != expectation {
			t.Errorf("Expected %v but got %v", expectation, actual)
		}
	}
}

func TestInitializeLoker6(t *testing.T) {
	loker := make([]Loker, 2)
	command := []string{"init", "1", "2"}
	res, err := initializeLoker(command, &loker)
	if res {
		t.Errorf("Loker should not initialized")
	} else {
		expectation := "Too much argument, expected 1, have 2"
		actual := err.Error()
		if actual != expectation {
			t.Errorf("Expected %v but got %v", expectation, actual)
		}
	}
}

func TestIsLokerInitialized1(t *testing.T) {
	loker = nil
	actual := isLokerInitialized(loker)
	expectation := false
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestIsLokerInitialized2(t *testing.T) {
	loker = []Loker{Loker{1, "SIM", 101010}}
	actual := isLokerInitialized(loker)
	expectation := true
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestFindEmptyLoker1(t *testing.T) {
	loker = []Loker{Loker{1, "SIM", 101010}, Loker{}}
	actual := findEmptyLoker(loker)
	expectation := 1
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestFindEmptyLoker2(t *testing.T) {
	loker = []Loker{Loker{1, "SIM", 101010}, Loker{2, "SIM", 112211}, Loker{}, Loker{3, "SIM", 111111}}
	actual := findEmptyLoker(loker)
	expectation := 2
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestFindEmptyLoker3(t *testing.T) {
	loker = []Loker{Loker{1, "SIM", 101010}, Loker{2, "SIM", 112211}, Loker{3, "SIM", 111111}}
	actual := findEmptyLoker(loker)
	expectation := -1
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestIsLokerEmpty1(t *testing.T) {
	loker = []Loker{Loker{1, "SIM", 101010}, Loker{}}
	actual := isLokerEmpty(loker)
	expectation := false
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestIsLokerEmpty2(t *testing.T) {
	loker = []Loker{}
	actual := isLokerEmpty(loker)
	expectation := true
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestIsLokerEmpty3(t *testing.T) {
	loker = []Loker{Loker{1, "SIM", 101010}, Loker{2, "SIM", 112211}, Loker{3, "SIM", 111111}}
	actual := isLokerEmpty(loker)
	expectation := false
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestInputLoker1(t *testing.T) {
	loker := make([]Loker, 10)
	command := []string{"input", "SIM", "111111"}
	res, info, err := inputLoker(command, &loker)
	if res {
		expectation := Loker{1, "SIM", 111111}
		actual := loker[0]
		if actual != expectation {
			t.Errorf("Expected %v but got %v", expectation, actual)
		}

		expectation2 := 1
		actual2 := info
		if actual2 != expectation2 {
			t.Errorf("Expected %v but got %v", expectation, actual)
		}
	} else {
		t.Errorf("Error Message : %v", err)
	}
}

func TestInputLoker2(t *testing.T) {
	loker := make([]Loker, 1)
	loker[0] = Loker{1, "SIM", 121212}
	command := []string{"input", "SIM", "111111"}
	res, info, err := inputLoker(command, &loker)
	if res {
		t.Errorf("Loker should full")
		_ = info
	} else {
		expectation := "Locker Full"
		actual := err.Error()
		if actual != expectation {
			t.Errorf("Expected %v but got %v", expectation, actual)
		}
	}
}

func TestInputLoker3(t *testing.T) {
	loker := make([]Loker, 2)
	command := []string{"input", "SIM"}
	res, info, err := inputLoker(command, &loker)
	if res {
		t.Errorf("Loker should full")
		_ = info
	} else {
		expectation := "Too few argument, expected 2, have 1"
		actual := err.Error()
		if actual != expectation {
			t.Errorf("Expected %v but got %v", expectation, actual)
		}
	}
}

func TestInputLoker4(t *testing.T) {
	loker := make([]Loker, 2)
	command := []string{"input", "SIM", "121212", "122222"}
	res, info, err := inputLoker(command, &loker)
	if res {
		t.Errorf("Loker should full")
		_ = info
	} else {
		expectation := "Too much argument, expected 2, have 3"
		actual := err.Error()
		if actual != expectation {
			t.Errorf("Expected %v but got %v", expectation, actual)
		}
	}
}

func TestLeaveLoker1(t *testing.T) {
	loker := make([]Loker, 2)
	loker[0] = Loker{1, "SIM", 121212}
	command := []string{"leave", "1"}
	res, info, err := leaveLoker(command, &loker)
	if res {
		expectation := 1
		actual := info
		if actual != expectation {
			t.Errorf("Expected %v but got %v", expectation, actual)
		}
	} else {
		t.Errorf("Error Message : %v", err)
	}
}

func TestLeaveLoker2(t *testing.T) {
	loker := make([]Loker, 2)
	command := []string{"leave", "2"}
	res, info, err := leaveLoker(command, &loker)
	if res {
		t.Errorf("Should not leave loker")
		_ = info
	} else {
		expectation := "Loker already empty"
		actual := err.Error()
		if actual != expectation {
			t.Errorf("Expected %v but got %v", expectation, actual)
		}
	}
}

func TestLeaveLoker3(t *testing.T) {
	loker := make([]Loker, 2)
	command := []string{"leave", "0"}
	res, info, err := leaveLoker(command, &loker)
	if res {
		t.Errorf("Should not leave loker")
		_ = info
	} else {
		expectation := "Loker number cannot less than 1"
		actual := err.Error()
		if actual != expectation {
			t.Errorf("Expected %v but got %v", expectation, actual)
		}
	}
}

func TestLeaveLoker4(t *testing.T) {
	loker := make([]Loker, 2)
	command := []string{"leave"}
	res, info, err := leaveLoker(command, &loker)
	if res {
		t.Errorf("Should not leave loker")
		_ = info
	} else {
		expectation := "Too few argument, expected 1, have 0"
		actual := err.Error()
		if actual != expectation {
			t.Errorf("Expected %v but got %v", expectation, actual)
		}
	}
}

func TestLeaveLoker5(t *testing.T) {
	loker := make([]Loker, 2)
	command := []string{"leave", "0", "1"}
	res, info, err := leaveLoker(command, &loker)
	if res {
		t.Errorf("Should not leave loker")
		_ = info
	} else {
		expectation := "Too much argument, expected 1, have 2"
		actual := err.Error()
		if actual != expectation {
			t.Errorf("Expected %v but got %v", expectation, actual)
		}
	}
}

func TestLeaveLoker6(t *testing.T) {
	loker := make([]Loker, 2)
	command := []string{"leave", "loker"}
	res, info, err := leaveLoker(command, &loker)
	if res {
		t.Errorf("Should not leave loker")
		_ = info
	} else {
		expectation := "strconv.Atoi: parsing \"loker\": invalid syntax"
		actual := err.Error()
		if actual != expectation {
			t.Errorf("Expected '%v' but got '%v'", expectation, actual)
		}
	}
}

func TestFindById1(t *testing.T) {
	loker = nil
	actual := findById(loker, 101010)
	expectation := -1
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestFindById2(t *testing.T) {
	loker = []Loker{Loker{1, "SIM", 101010}}
	actual := findById(loker, 101010)
	expectation := 0
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestFindById3(t *testing.T) {
	loker = []Loker{Loker{1, "SIM", 101010}, Loker{2, "KTP", 121212}, Loker{3, "SIM", 111111}}
	actual := findById(loker, 121212)
	expectation := 1
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestFindByType1(t *testing.T) {
	loker = nil
	actual := findByType(loker, "SIM")
	expectation := []int{}
	if reflect.DeepEqual(actual, expectation) {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestFindByType2(t *testing.T) {
	loker = []Loker{Loker{1, "SIM", 101010}}
	actual := findByType(loker, "SIM")
	expectation := []int{101010}
	if !reflect.DeepEqual(actual, expectation) {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestFindByType3(t *testing.T) {
	loker = []Loker{Loker{1, "SIM", 101010}, Loker{2, "KTP", 121212}, Loker{3, "SIM", 111111}}
	actual := findByType(loker, "SIM")
	expectation := []int{101010, 111111}
	if !reflect.DeepEqual(actual, expectation) {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}
