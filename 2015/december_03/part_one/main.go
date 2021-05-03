package main

import (
	"fmt"
	"io/ioutil"
	"sync"
	"time"
)

var onlyOnce sync.Once

type coordinate struct {
	x int32
	y int32
}

type delivery struct {
	coordinate
	trackingLocation []coordinate
}

func main() {
	defer measureTime("1st.go")()

	clues, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	santa := newDelivery()

breakpoint:
	for _, c := range clues {
		switch string(c) {
		case "^":
			santa.startingPoint()
			santa.goNorth()
		case ">":
			santa.startingPoint()
			santa.goEast()
		case "v":
			santa.startingPoint()
			santa.goSouth()
		case "<":
			santa.startingPoint()
			santa.goWest()
		default:
			break breakpoint
		}
	}

	uniqueCoordinate := findUniqueCoordinate(santa.trackingLocation)
	fmt.Printf("Number of Houses that received a gift are %d\n", len(uniqueCoordinate))
}

func newDelivery() *delivery {
	return &delivery{
		coordinate: coordinate{
			x: 0,
			y: 0,
		},
		trackingLocation: []coordinate{
			{
				x: 0,
				y: 0,
			},
		},
	}
}

func findUniqueCoordinate(c []coordinate) []coordinate {
	keys := make(map[coordinate]bool)
	list := []coordinate{}

	for _, entry := range c {
		v := keys[entry]
		if !v {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func (d *delivery) startingPoint() {
	onlyOnce.Do(func() {
	})
}

func (d *delivery) goNorth() {
	d.coordinate.y++
	d.trackingLocation = append(d.trackingLocation, d.coordinate)
}

func (d *delivery) goSouth() {
	d.coordinate.y--
	d.trackingLocation = append(d.trackingLocation, d.coordinate)
}

func (d *delivery) goEast() {
	d.coordinate.x++
	d.trackingLocation = append(d.trackingLocation, d.coordinate)
}

func (d *delivery) goWest() {
	d.coordinate.x--
	d.trackingLocation = append(d.trackingLocation, d.coordinate)
}

func measureTime(appName string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("Time taken by %s application is %v \n", appName, time.Since(start))
	}
}
