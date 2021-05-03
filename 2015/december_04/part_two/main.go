package main

import (
	"fmt"
	"time"
)

func main() {
	defer measureTime("part_one")()

}

func measureTime(appName string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("Time taken by %s application is %v \n", appName, time.Since(start))
	}
}
