package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: program <filename>")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var rangesToCheck []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rangesToCheck = strings.Split(scanner.Text(), ",")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var invalidIDs []int
	for _, rangeToCheck := range rangesToCheck {
		splitRange := strings.Split(rangeToCheck, "-")
		start, err := strconv.Atoi(splitRange[0])
		if err != nil {
			log.Fatal(err)
		}

		end, err := strconv.Atoi(splitRange[1])
		if err != nil {
			log.Fatal(err)
		}

		invalidIDs = append(invalidIDs, getInvalidIDsInRange(start, end)...)	
	}

	var sum int
	for _, id := range invalidIDs {
		sum += id	
	}
	fmt.Println(sum)
}

func getInvalidIDsInRange(start, end int) []int {
	result := []int{}
	for i := start; i < end; i++ {
		if checkRepeatedNum(i) {
			result = append(result, i)
		}
	}
	return result
}

// i didn't like this prompt so i heavily used claude to solve this one
func checkRepeatedNum(num int) bool {
	// Count digits
	temp := num
	digits := 0
	for temp > 0 {
		digits++
		temp /= 10
	}
	
	// Try all pattern lengths that could repeat at least twice
	for patternLen := 1; patternLen <= digits/2; patternLen++ {
		if digits%patternLen != 0 {
			continue
		}
		
		divisor := pow10(patternLen)
		pattern := num % divisor
		
		// Check if entire number is this pattern repeated
		temp := num
		valid := true
		for temp > 0 {
			if temp%divisor != pattern {
				valid = false
				break
			}
			temp /= divisor
		}
		
		if valid {
			return true
		}
	}
	
	return false
}

func pow10(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= 10
	}
	return result
}
