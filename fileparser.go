package main

//TODO:
// better feedback in the different cases

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	defaultInputFile  string = "input.txt"
	defaultOutputFile string = "output.txt"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, scanner.Err()
}

func filter(lines []string, parseFilter string) []string {
	var filteredLines []string

	for _, line := range lines {
		lineString := string(line)
		if strings.Contains(lineString, parseFilter) {
			lineSlice := strings.Fields(lineString)
			lineSliceWithoutCounter := lineSlice[1:]
			lineRune := []rune(strings.Join(lineSliceWithoutCounter, " "))
			filteredLines = append(filteredLines, string(lineRune))
		}
	}
	return filteredLines
}

func appendLines(lines []string, path string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func filterNewLines(outputLines []string, inputLines []string) []string {
	var newLines []string
	for _, line := range inputLines {
		if contains(outputLines, line) == false {
			newLines = append(newLines, line)
		}
	}
	return newLines
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

/**
Defining command line arguments
*/
var inputFlag = flag.String("input", defaultInputFile, "Define the input file to parse.")
var outputFlag = flag.String("output", defaultOutputFile, "Define the file to store the output.")
var filterFlag = flag.String("filter", "", "Define the filter to use for the input file parsing.")

func init() {
	// Defining short version for long flags
	flag.StringVar(inputFlag, "i", defaultInputFile, "see input")
	flag.StringVar(outputFlag, "o", defaultOutputFile, "see output")
	flag.StringVar(filterFlag, "f", "", "see filter")
}

func main() {
	flag.Parse()
	println("Input file: "+*inputFlag, "\nOutput file: "+*outputFlag, "\nFilter: "+*filterFlag)

	inputFile := *inputFlag
	outputFile := *outputFlag
	parseFilter := *filterFlag

	outputLines, err := readLines(outputFile)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	inputLines, err := readLines(inputFile)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	filteredLines := filter(inputLines, parseFilter)
	filteredNewLines := filterNewLines(outputLines, filteredLines)
	for _, line := range filteredNewLines {
		fmt.Println(line)
	}

	if err := appendLines(filteredNewLines, defaultOutputFile); err != nil {
		log.Fatalf("writeLines: %s", err)
	}
}
