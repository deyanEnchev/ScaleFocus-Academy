package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	m := groupSlices(citiesAndPrices())

	// fmt.Println(m)
	for k := range m {
		fmt.Printf("%s: ", k)
		for i, v := range m[k] {
			if i == 0 {
				fmt.Print(v)
			} else {
				fmt.Print(",", v)
			}
			
		}
		fmt.Println()
	}
}

func citiesAndPrices() ([]string, []int) {
	rand.Seed(time.Now().UnixMilli())
	cityChoices := []string{"Berlin", "Moscow", "Chicago", "Tokyo", "London"}
	dataPointCount := 10

	cities := make([]string, dataPointCount)
	for i := range cities {
		cities[i] = cityChoices[rand.Intn(len(cityChoices))]
	}

	prices := make([]int, dataPointCount)

	for i := range prices {
		prices[i] = rand.Intn(100)
	}

	return cities, prices
}

func groupSlices(keySlice []string, valueSlice []int) map[string][]int {

	m := make(map[string][]int)

	for i := 0; i < len(keySlice); i++ {
		m[keySlice[i]] = append(m[keySlice[i]], valueSlice[i])

	}

	return m
}
