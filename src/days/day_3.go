package days

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"
)

type number struct {
	val        int
	containVal bool
	width      int
	height     int
	length     int
}

func (n *number) add(x int) {
	if !n.containVal {
		n.containVal = true
	}
	n.val = (n.val*10 + x)
	n.length++
}

func (n *number) checkBorder(arr *[]string, xLoc int, yLoc int) bool {
  xLoc++
	checkUp := false
	checkDown := false
	symbols := "123456789."

	if yLoc-1 >= 0 {
		checkUp = true
	}

	if yLoc+1 <= n.height {
		checkDown = true
	}

	xStart := xLoc - n.length
	if xLoc <= n.width {
		right := string((*arr)[yLoc][xLoc])
		if !strings.ContainsAny(right, symbols) {
			return true
		}
		n.length++
	} //checkLeft

	if xStart-1 >= 0 {
		left := string((*arr)[yLoc][xStart-1])
		if !strings.ContainsAny(left, symbols) {
			return true
		}
		n.length++
		xStart--
	} //checkright

	for i := 0; i < n.length; i++ {
		x := xStart + i
		if checkUp { // up
			up := (*arr)[yLoc-1][x]
			if !strings.ContainsAny(string(up), symbols) {
				return true
			}
		}

		if checkDown { //down
			down := (*arr)[yLoc+1][x]
			if !strings.ContainsAny(string(down), symbols) {
				return true
			}
		}
	}

	return false
}

func (n *number) reset() {
	n.val = 0
	n.containVal = false
	n.length = 0
}

type Day_3 struct{}

func (d Day_3) Solution(dir string, filename string) {

	filename = filepath.Join(dir, "day_3", filename)
	file, err := os.Open(filename)

	if err != nil {
		log.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	d.question_1(bufio.NewScanner(file))
}

func (d Day_3) question_1(scanner *bufio.Scanner) {
	if err := scanner.Err(); err != nil {
		log.Println("Error reading from file:", err)
		return
	}

	fileData := []string{}

	for scanner.Scan() {
		fileData = append(fileData, string(scanner.Text()))
	}

	total := 0
	for idx1, rowdata := range fileData {
		num := number{val: 0, containVal: false, width: len(rowdata) - 1, height: len(fileData) - 1, length: 0}
		for idx2, eachChar := range rowdata {
			if unicode.IsDigit(eachChar) {
				ret, _ := strconv.Atoi(string(eachChar))
				num.add(ret)
			} else if num.containVal {
				if num.checkBorder(&fileData, idx2-1, idx1) {
					total += num.val
				}
				num.reset()
			}
		}
		if num.containVal {
			if num.checkBorder(&fileData,num.width, idx1) {
				total += num.val
			}
			num.reset()
		}

	}

	println("Puzzle 1")
	println("Solution: ", total)
	println("---------")
}
