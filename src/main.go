package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/satwikShresth/AOC/src/day_1"
)

func main() {
  var filename string

	inputDir := flag.String("i", "", "Directory for input data")
	solveExamples := flag.Bool("e", false, "Solve example problems if set")
	debug := flag.Bool("debug", false, "Enable debug mode for verbose logging")

	flag.Parse()

	if *debug {
		log.SetFlags(log.LstdFlags)
		log.SetOutput(os.Stdout)
		log.Println("Debug mode enabled")
	} else {
		log.SetFlags(0)
		log.SetOutput(io.Discard)
	}

  if *solveExamples{
    filename = "example.data"
  } else {
    filename="input.data"
  }

	if *inputDir == "" {
		fmt.Println("Input directory is required.")
		flag.Usage()
		os.Exit(1)
	}


	log.Println("Using input directory: ", *inputDir)

	fmt.Println("Advent of code 2023:-")
	fmt.Println("---------------------")
  fmt.Println("Day 1:")
	fmt.Println("-------")
  day_1.Solution(*inputDir,filename)
	fmt.Println("---------------------")
}
