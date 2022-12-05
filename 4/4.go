package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	type jobRange struct {
		upper int
		lower int
	}

	type pair struct {
		elfOne jobRange
		elfTwo jobRange
	}

	var elfPairs [10000]pair

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file word by word using scanner
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	elfPairId := 0
	for scanner.Scan() {
		// do something with a word
		ranges := strings.Split(scanner.Text(), ",")
		jobNumbersElfOne := strings.Split(ranges[0], "-")
		jobNumbersElfTwo := strings.Split(ranges[1], "-")
		elfPairs[elfPairId].elfOne.lower, err = strconv.Atoi(jobNumbersElfOne[0])
		elfPairs[elfPairId].elfOne.upper, err = strconv.Atoi(jobNumbersElfOne[1])
		elfPairs[elfPairId].elfTwo.lower, err = strconv.Atoi(jobNumbersElfTwo[0])
		elfPairs[elfPairId].elfTwo.upper, err = strconv.Atoi(jobNumbersElfTwo[1])

		elfPairId++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//check if pairs fully encapsulate
	encapsulate := 0
	for p := 0; p < elfPairId; p++ {
		if elfPairs[p].elfOne.lower <= elfPairs[p].elfTwo.lower && elfPairs[p].elfOne.upper >= elfPairs[p].elfTwo.upper {
			encapsulate++
		} else if elfPairs[p].elfTwo.lower <= elfPairs[p].elfOne.lower && elfPairs[p].elfTwo.upper >= elfPairs[p].elfOne.upper {
			encapsulate++
		}
	}

	// check if no overlap
	overlap := elfPairId
	for p := 0; p < elfPairId; p++ {
		if elfPairs[p].elfOne.upper < elfPairs[p].elfTwo.lower {
			overlap--
		} else if elfPairs[p].elfTwo.upper < elfPairs[p].elfOne.lower {
			overlap--
		}
	}

	fmt.Println(encapsulate)
	fmt.Println(overlap)

}
