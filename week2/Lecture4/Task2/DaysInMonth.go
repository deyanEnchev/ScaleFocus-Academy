package main

import "fmt"

func main() {

	var month, year int = 2, 2004

	days, isLeapYear := daysInMonth(month, year)

	fmt.Printf("The days in the %d month are %d.\n", month, days)

	if isLeapYear {
		fmt.Printf("The year %d is a leap year.", year)
	} else {
		fmt.Printf("The year %d is not a leap year.", year)
	}

}

func daysInMonth(month int, year int) (int, bool) {

	var isLeapYear bool

	if year%400 == 0 && year%4 == 0 {
		isLeapYear = true
	}

	if year%4 == 0 && year%100 != 0 {
		isLeapYear = true
	}

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31, isLeapYear
	case 4, 6, 9, 11:
		return 30, isLeapYear
	case 2:

		if isLeapYear {
			return 29, true
		}

		return 28, false
	default:
		fmt.Println("Wrong month number! Choose a month from 1-12")
	}
	return 0, false
}
