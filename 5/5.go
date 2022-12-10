package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func reverseArray(arr []byte) []byte {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	defer f.Close()

	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	type stack struct {
		id    int
		state []byte
	}

	type instruction struct {
		amount int
		from   int
		to     int
	}

	// find row length
	var stacksIndex []byte
	for c := range content { //scan to find a one, returns highest stack
		if content[c] == '1' {
			for d := 0; unicode.IsDigit(rune(content[c+d])); d += 4 {
				stacksIndex = append(stacksIndex, (content[c+d]))
			}
			break
		}
	}

	var amountOfStacks int = int((stacksIndex[len(stacksIndex)-1] - 48))
	charsPerLine := amountOfStacks*4 + 1

	var stacksList []stack

	for stackNum := 0; stackNum < amountOfStacks; stackNum++ {
		var tempStack stack
		tempStack.id = stackNum + 1
		for i := stackNum*4 + 1; i < (amountOfStacks-1)*charsPerLine; i += (charsPerLine) {
			if content[i] != 32 {
				tempStack.state = append(tempStack.state, content[i])
			}
		}
		stacksList = append(stacksList, tempStack)
	}

	var instructionList []instruction
	for scanner.Scan() {
		if scanner.Text() == "move" {
			var newInstruction instruction
			scanner.Scan()
			newInstruction.amount, err = strconv.Atoi((scanner.Text()))
			scanner.Scan()
			scanner.Scan()
			newInstruction.from, err = strconv.Atoi((scanner.Text()))
			scanner.Scan()
			scanner.Scan()
			newInstruction.to, err = strconv.Atoi((scanner.Text()))
			instructionList = append(instructionList, newInstruction)
		}
	}

	for s := range stacksList {
		stacksList[s].state = reverseArray(stacksList[s].state)
	}

	fmt.Println(stacksList)
	for i := range instructionList {
		n := instructionList[i].amount
		for m := n - 1; m >= 0; m-- {
			stacksList[instructionList[i].to-1].state = append(stacksList[instructionList[i].to-1].state, (stacksList[instructionList[i].from-1].state[len(stacksList[instructionList[i].from-1].state)-1-m]))
		}
		stacksList[instructionList[i].from-1].state = stacksList[instructionList[i].from-1].state[:(len(stacksList[instructionList[i].from-1].state) - n)] // Pop
		fmt.Println(stacksList)
	}

	for s := range stacksList {
		fmt.Printf("%c", stacksList[s].state[len(stacksList[s].state)-1])
	}
}
