package helper

import "fmt"

const (
	BREAK_CODE = iota
	CONTINUE_CODE
	OTHER_CODE
)

func ProcCommand(command string) int {
	switch command {
	case "q", "quit", "exit":
		fmt.Println("Hey Bye!")
		return BREAK_CODE
	case "":
		return CONTINUE_CODE
	case "m":
		showMonth()
		return CONTINUE_CODE
	case "a":
		showArea()
		return CONTINUE_CODE
	}
	return OTHER_CODE
}

func showMonth() {
	months := [12]string{
		"JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC",
	}

	fmt.Println("Month list")

	counter := 1
	for _, m := range months {
		fmt.Printf(" %2d ... %s", counter, m)
		if counter%3 == 0 {
			fmt.Println()
		}
		counter++
	}
	fmt.Println()
}

func showArea() {
	areas := map[string]string{
		"TYO": "Tokyo",
		"SEL": "Soul",
		"BKK": "Bangkok",
		"LAX": "Los Angeles",
		"SIN": "Singapore",
		"FRA": "Frankfurt",
	}

	fmt.Println("Area list")
	for c, n := range areas {
		fmt.Printf(" %s ... %s\n", c, n)
	}
	fmt.Println()
}
