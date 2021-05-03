package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	defer measureTime("part_one")()
	c := make(chan string)
	done := make(chan bool)

	go func() {
		str := []string{"0", "00", "000", "0000", "00000"}
		for _, s := range str {
			c <- s
		}
	}()

	select {
	case v := <-c:
		fmt.Println(v)
		if isStartedWithFiveZero(v) {
			done <- true
		}
	case d := <-done:
		if d {
			os.Exit(0)
		}
	default:
		fmt.Println("no response received!")
	}

}

func isStartedWithFiveZero(s string) bool {
	return strings.Contains(s, "00000")
}

func measureTime(appName string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("Time taken by %s application is %v \n", appName, time.Since(start))
	}
}
