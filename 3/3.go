package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Println((PartTwo()))

}

func PartOne() int {

	type backpack struct {
		contents     string
		firstPocket  string
		secondPocket string
		commonItem   byte
		priority     int
	}

	var backpackList [1000]backpack

	// open file
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	numBackpacks := 0
	for scanner.Scan() {
		// do something with a line
		backpackList[numBackpacks].contents = scanner.Text()
		backpackList[numBackpacks].firstPocket = (backpackList[numBackpacks].contents)[:(len(backpackList[numBackpacks].contents) / 2)]
		backpackList[numBackpacks].secondPocket = (backpackList[numBackpacks].contents)[(len(backpackList[numBackpacks].contents) / 2):]
		numBackpacks++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < numBackpacks; i++ {

		backpackList[i].commonItem = 0
		for j := 0; j < len(backpackList[i].firstPocket); j++ {
			for k := 0; k < len(backpackList[i].secondPocket); k++ {

				if backpackList[i].firstPocket[j] == backpackList[i].secondPocket[k] {
					backpackList[i].commonItem = backpackList[i].firstPocket[j]
				}

			}
		}
	}

	for i := 0; i < numBackpacks; i++ {

		com := backpackList[i].commonItem
		if com > 96 { // lower case
			com -= 96
		} else { // upper
			com -= 38
		}

		backpackList[i].priority = int(com)

	}

	//sum up priorities
	priorityTotal := 0
	for i := 0; i < numBackpacks; i++ {
		priorityTotal += backpackList[i].priority
	}

	return priorityTotal

}

func PartTwo() int {

	type backpack struct {
		contents   string
		commonItem byte
		priority   int
	}

	var backpackList [1000]backpack

	// open file
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	numBackpacks := 0
	for scanner.Scan() {
		// do something with a line
		backpackList[numBackpacks].contents = scanner.Text()
		numBackpacks++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	total := 0
	for i := 0; i < numBackpacks/3; i++ {

		a := backpackList[i*3].contents
		b := backpackList[i*3+1].contents
		c := backpackList[i*3+2].contents

		var common [100]byte
		charPos := 0
		for j := 0; j < len(a); j++ {
			for k := 0; k < len(b); k++ {
				if a[j] == b[k] {
					common[charPos] = a[j]
					charPos++
				}
			}
		}

		var commonAll byte
		for j := 0; j < len(common); j++ {
			for k := 0; k < len(c); k++ {
				if common[j] == c[k] {
					commonAll = common[j]
				}
			}
		}

		if commonAll > 96 { // lower case
			commonAll -= 96
		} else { // upper
			commonAll -= 38
		}

		total += int(commonAll)

	}

	return total

}
