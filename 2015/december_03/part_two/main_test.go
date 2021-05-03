package main

import (
	"io/ioutil"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "1st_sample_input.txt",
		},
		{
			name: "2nd_sample_input.txt",
		},
		{
			name: "3rd_sample_input.txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			clues, err := ioutil.ReadFile(tt.name)
			if err != nil {
				panic(err)
			}

			santa := newDelivery()
			roboSanta := newDelivery()

			for i, c := range clues {
				if modZero(i) {
					roboSanta.route(c)
				} else {
					santa.route(c)
				}
			}

			aggregatedCoordinate := findAggregatedCoordinate(santa.trackingLocation, roboSanta.trackingLocation)
			uniqueCoordinate := findUniqueCoordinate(aggregatedCoordinate)
			t.Logf("Number of Houses from sample %s that received a gift are %d\n", tt.name, len(uniqueCoordinate))
		})
	}
}
