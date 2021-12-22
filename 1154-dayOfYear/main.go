package main

import "strconv"

func dayOfYear(date string) int {
	year, _ := strconv.Atoi(date[:4])
	month, _ := strconv.Atoi(date[5:7])
	day, _ := strconv.Atoi(date[8:10])
	if month < 2 {
		return day
	}
	months := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30}
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		months[1] = 29
	}
	for i := 0; i < month-1; i++ {
		day += months[i]
	}
	return day
}

func main() {

}
