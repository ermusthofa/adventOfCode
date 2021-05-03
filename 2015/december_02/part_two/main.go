package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type box struct {
	width  int32
	height int32
	length int32
}

var (
	boxes       []box
	totalRibbon int32
)

func main() {

	defer measureTime("1st.go")()

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		s := strings.Split(fileScanner.Text(), "x")
		width, height, length := toInt(s)
		boxes = append(boxes, createBox(width, height, length))
	}

	for _, b := range boxes {
		shortestX, shortestY := findShortestRibbon(b)
		totalRibbon += (2 * shortestX) + (2 * shortestY) + (b.width * b.height * b.length)
	}

	fmt.Printf("Total amount of ribbon needed by the elves is %d feet\n", totalRibbon)
}

func createBox(width, height, length int32) box {
	return box{
		width:  width,
		height: height,
		length: length,
	}
}

func findShortestRibbon(b box) (int32, int32) {
	var d, shortestDistance []int32
	var ld int32
	var index int

	d = append(d, b.width)
	d = append(d, b.height)
	d = append(d, b.length)

	ld = d[0]
	for i, v := range d {
		if v > ld {
			ld = v
			index = i
		}
	}

	shortestDistance = removeFromSlice(d, index)
	return shortestDistance[0], shortestDistance[1]
}

func removeFromSlice(slice []int32, s int) []int32 {
	return append(slice[:s], slice[s+1:]...)
}

func toInt(s []string) (int32, int32, int32) {
	w, _ := strconv.Atoi(s[0])
	h, _ := strconv.Atoi(s[1])
	l, _ := strconv.Atoi(s[2])
	return int32(w), int32(h), int32(l)
}

func measureTime(funcName string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("Time taken by %s function is %v \n", funcName, time.Since(start))
	}
}
