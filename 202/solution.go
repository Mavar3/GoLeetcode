package main

import (
	"fmt"
	"strconv"
	"strings"
)

const resultValue = "1"

func IsHappy(n int) bool {
	visitedMap := make(map[string]int)
	for valueForCheck := fmt.Sprint(n); visitedMap[valueForCheck] == 0; valueForCheck = detectNewHappyNumber(strings.Split(valueForCheck, "")) {
		visitedMap[valueForCheck] += 1
	}

	return visitedMap[resultValue] > 0
}

func detectNewHappyNumber(str_numbers []string) string {
	result := 0
	for _, str_number := range str_numbers {
		int_value, _ := strconv.Atoi(str_number)
		result += int_value * int_value
	}
	return fmt.Sprint(result)
}
