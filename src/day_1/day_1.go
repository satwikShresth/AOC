package day_1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"unicode"
)

type digit struct {
	numString string
	numValue  int
}

func (d *digit) size() int {
	return len(d.numString)
}

type lineValue struct {
	firstValue int
	lastValue  int
	firstFound bool
}

func (l *lineValue) setValue(val int) {
	if !l.firstFound {
		l.firstValue = val
		l.firstFound = true
	}
	l.lastValue = val
}

func (l *lineValue) getfinalValue() int {
	return (l.firstValue * 10) + l.lastValue
}

// one
// two
// three
// four
// five
// six
// seven
// eight
// nine

func Solution(dir string, filename string) {
	filename = filepath.Join(dir, "day_1", filename)
	file, err := os.Open(filename)

	if err != nil {
		log.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	question_1(bufio.NewScanner(file))
	file.Seek(0, 0)
	question_2(bufio.NewScanner(file))
}

func question_1(scanner *bufio.Scanner) {

	if err := scanner.Err(); err != nil {
		log.Println("Error reading from file:", err)
		return
	}

	total := 0
	idx := 1
	for scanner.Scan() {
		var val lineValue

		for _, runeValue := range scanner.Text() {

			if unicode.IsDigit(runeValue) {
				ret, _ := strconv.Atoi(string(runeValue))
				log.Println("digit: ", ret)
				val.setValue(ret)
			}

		}

		log.Printf("line: %d Value: %d", idx, val.getfinalValue())
		total += val.getfinalValue()
		idx++
	}

	fmt.Println("Puzzle 1")
	fmt.Println("Solution: ", total)
	fmt.Println("---------------------")
}

func question_2(scanner *bufio.Scanner) {

	if err := scanner.Err(); err != nil {
		log.Println("Error reading from file:", err)
		return
	}

	digitMap := map[string][]digit{
		"o": {
			digit{"one", 1},
		},
		"t": {
			digit{"two", 2},
			digit{"three", 3},
		},
		"f": {
			digit{"four", 4},
			digit{"five", 5},
		},
		"s": {
			digit{"six", 6},
			digit{"seven", 7},
		},
		"e": {
			digit{"eight", 8},
		},
		"n": {
			digit{"nine", 9},
		},
	}

	total := 0
	idx := 1
	for scanner.Scan() {
		var val lineValue

		for idx, runeValue := range scanner.Text() {

			if unicode.IsDigit(runeValue) {

				ret, _ := strconv.Atoi(string(runeValue))
				log.Println("digit: ", ret)
				val.setValue(ret)

			} else {

				value, exists := digitMap[string(runeValue)]

				if exists {
					for _, digit := range value {

						if digit.size()+idx <= len(scanner.Text()) && string(scanner.Text()[idx:digit.size()+idx]) == digit.numString {
							val.setValue(digit.numValue)
						}
					}

				}

			}

		}

		log.Printf("line: %d Value: %d", idx, val.getfinalValue())
		total += val.getfinalValue()
		idx++
	}

	fmt.Println("Puzzle 2")
	fmt.Println("Solution: ", total)
	fmt.Println("---------------------")
}
