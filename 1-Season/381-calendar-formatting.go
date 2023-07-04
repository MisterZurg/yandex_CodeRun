package main

import "fmt"

const (
	Monday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	var nDays int
	var weekday string

	fmt.Scan(&nDays, &weekday)

	var startFrom int
	switch weekday {
	case "Monday":
		startFrom = Monday
	case "Tuesday":
		startFrom = Tuesday
	case "Wednesday":
		startFrom = Wednesday
	case "Thursday":
		startFrom = Thursday
	case "Friday":
		startFrom = Friday
	case "Saturday":
		startFrom = Saturday
	case "Sunday":
		startFrom = Sunday
	}

	for i := 0; i < nDays+startFrom; i++ {
		if i > 0 && i%7 == 0 {
			fmt.Println()
		}

		if i < startFrom {
			fmt.Printf(".. ")
			continue
		} else if i-startFrom < 9 {
			fmt.Printf(".%d ", i-startFrom+1)
			continue
		}

		fmt.Printf("%d ", i-startFrom+1)
	}
}
