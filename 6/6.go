package main

import (
	"fmt"
	"log"
	"os"
)

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func main() {

	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	uniqueStream := 14
	startOfPacketMarker := 0

	for i := range content {

		var uniqueTester []int
		var chars []int

		for j := 0; j < uniqueStream; j++ {
			chars = append(chars, int(content[i+j]))
		}

		uniqueTester = unique(chars)
		if !(len(uniqueTester) < len(chars)) {
			startOfPacketMarker = i + uniqueStream
			fmt.Println(startOfPacketMarker)
			break
		}

	}
}
