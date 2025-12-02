package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	numZerosHit := 0
	numZeroPassedTotal := 0

	currentVal := 50
	fmt.Printf("The dial starts by pointing at %d\n", currentVal)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		direction := string(line[0])
		amount, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}

		numZeroPasses := 0
		switch direction {
		case "L":
			currentVal -= (amount % 100)
			if currentVal < 0 {
				currentVal = 100 + currentVal
				if (currentVal + (amount % 100)) != 100 {
					numZeroPasses += 1
				}
			}
		case "R":
			currentVal += (amount % 100)
			if currentVal == 100 {
				currentVal = 0 // this 0 covered below with currentVal == 0 check
			} else if currentVal > 100 {
				currentVal = currentVal - 100
				if (currentVal - (amount % 100) != 0) {
					numZeroPasses += 1
				}
			}
		}

		numZeroPasses += int(amount/100)

		fmt.Printf("The dial is rotated %s to point at %d; passed zero %d times\n", line, currentVal, numZeroPasses)

		numZeroPassedTotal += numZeroPasses
		if currentVal == 0 {
			numZerosHit += 1
		}
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nNumber of Zeros Values found: %d\n", numZerosHit + numZeroPassedTotal)
}

