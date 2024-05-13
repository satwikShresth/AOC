package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Day_2 struct{}

type Game struct {
	id    int
	red   int
	blue  int
	green int
  possible bool
}
func (g *Game) print() {
  println("--------------------")
  println("id: ",g.id)
  println("----")
  println("Red: ",g.red)
  println("Blue: ",g.green)
  println("Green: ",g.blue)
  println("--------------------")
}

func (g *Game) power() int {
	return g.red * g.green * g.blue
}

func (d Day_2) Solution(dir string, filename string) {

	filename = filepath.Join(dir, "day_2", filename)
	file, err := os.Open(filename)

	if err != nil {
		log.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	d.question_1(bufio.NewScanner(file))
}

func (d Day_2) question_1(scanner *bufio.Scanner) {
	if err := scanner.Err(); err != nil {
		log.Println("Error reading from file:", err)
		return
	}

  totalPower := 0
  total := 0
	for scanner.Scan() {
    // println(scanner.Text())
		var id int
		line := strings.Split(scanner.Text(), ":")
		fmt.Sscanf(line[0], "Game %d", &id)

		str1 := strings.Split(line[1], ";")
    game := Game{
      id:id,
      possible: true,
    }

    // println("str1: ",str1)
		for _, iStr1 := range str1 {
			str2 := strings.Split(iStr1, ",")

      // println("str2: ",str2)
			for _, iStr2 := range str2 {
				var num int
				var colorVal string

				fmt.Sscanf(iStr2, "%d %s", &num, &colorVal)
        // println("iStr2: ",iStr2)

				switch colorVal {
				case "red":
          if num > game.red {game.red = num}
          if num> 12 { game.possible = false }
				case "green":
          if num > game.green {game.green = num}
          if num > 13 { game.possible = false }
				case "blue":
          if num > game.blue {game.blue = num}
          if num > 14 { game.possible = false }
				}
			}
		}

		if game.possible {
			total += game.id
		}
    // game.print()
    totalPower += game.power()
	}

	fmt.Println("Puzzle 1")
	fmt.Println("Solution: ", total)
	fmt.Println("---------")
	fmt.Println("Puzzle 2")
	fmt.Println("Solution: ", totalPower)
	fmt.Println("---------")
	return
}
