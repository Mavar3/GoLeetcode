package main

import (
	"strings"
)

func LongestSubstring(s string, k int) int {
	// Когда текст не подходит
	if len(s) < k {
		return 0
	}

	// Когда все буквы подходят под условие
	countMap := fillLittersCount(s)
	if isAllMoreOrEqK(k, countMap) {
		return len(s)
	}

	// Когда часть букв подходит под условие (откидываем, какие не подходят и продолжаем).
	// Вынести в метод?
	max := 0
	for _, el := range splitByExcludedLetters(s, k, countMap) {
		subArrayResult := LongestSubstring(el, k)
		if max < subArrayResult {
			max = subArrayResult
		}
	}

	return max
}

func fillLittersCount(str string) map[string][]int {
	littersCount := make(map[string][]int)
	for id, el := range str {
		char := string(el)
		littersCount[char] = append(littersCount[char], id)
	}

	return littersCount
}

func isAllMoreOrEqK(size int, result map[string][]int) bool {
	for _, el := range result {
		if len(el) < size {
			return false
		}
	}

	return true
}

func splitByExcludedLetters(str string, size int, result map[string][]int) []string {
	rune_str := []rune(str)
	for _, ids := range result {
		if len(ids) >= size {
			continue
		}
		for _, letterId := range ids {
			rune_str[letterId] = ' '
		}
	}
	return strings.Split(string(rune_str), " ")
}
