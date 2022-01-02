package main

func dayOfTheWeek(day int, month int, year int) string {
	w := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	m := []int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}
	if (year%4 == 0 && month < 3) || year == 2100 {
		day--
	}
	day += (year-1971)*365 + m[month-1] + (year-1968)/4
	return w[(day+4)%7]
}

func main() {

}
