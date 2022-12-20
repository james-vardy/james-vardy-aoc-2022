package main

import (
	"bufio"
	"log"
	"os"
)

var commandList []string
var nestingLevel = 0
var tree []string

func main() {

	totalLines := 0
	parentDirectory := ""
	dirSizes := make(map[string]int)
	dirSizes[parentDirectory] = 0
	
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		commandList = append(commandList, scanner.Text())
		totalLines++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	parentDirectory, size = handleCommand(0, parentDirectory)

	for i := 0; i < len(tree); i++ {
		println(tree[i] + "\n")
	}

}

func handleCommand(lineNum int, parentDirectory string)(string, int) {

	if (commandList[lineNum][4:] == "$ cd") {
		var dir = commandList[lineNum][5:]
		if dir == ".." {
			nestingLevel--
		} else if (dir == "/") {
			nestingLevel = 0
		}  else {
			tempStr := ""
			for n := 0; n < nestingLevel; n++ {
				tempStr += "-"
			}
			tempStr += dir
			tree = append(tree, tempStr)
			nestingLevel++
		} 
	}
	


	return parentDirectory

}

