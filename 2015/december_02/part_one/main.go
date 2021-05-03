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
	boxes        []box
	totalWrapper int32
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
		boxes = append(boxes, createBox(int32(width), int32(height), int32(length)))
	}

	for _, b := range boxes {
		totalWrapper += calculateWrapper(b)
	}

	fmt.Printf("Total amount of wrapping paper needed by the elves is %d square feet\n", totalWrapper)
}

func createBox(width, height, length int32) box {
	return box{
		width:  width,
		height: height,
		length: length,
	}
}

func calculateWrapper(b box) int32 {
	var minWrapper []int32
	var wrapper int32 = 0

	minWrapper = append(minWrapper, calculateWXH(b.width, b.height))
	minWrapper = append(minWrapper, calculateWXL(b.width, b.length))
	minWrapper = append(minWrapper, calculateHXL(b.height, b.length))

	smallestWrapper := findSmallestArea(minWrapper)

	for _, mw := range minWrapper {
		wrapper += mw
	}

	return (2 * wrapper) + smallestWrapper
}

func findSmallestArea(w []int32) int32 {
	var t int32 = w[0]
	for _, v := range w {
		if v < t {
			t = v
		}
	}
	return t
}

func calculateWXH(width, height int32) int32 {
	return width * height
}

func calculateWXL(width, length int32) int32 {
	return width * length
}

func calculateHXL(height, length int32) int32 {
	return height * length
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
