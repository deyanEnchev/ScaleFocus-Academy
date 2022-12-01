package main

import (
	"fmt"
	"log"
	"sort"
	"time"
)

func main() {
	input := []string{"Sep-14-2008", "Mar-18-2022", "Dec-03-2021", "Jan-16-1920", "Jan-26-2240"}

	format := "Jan-02-2006"

	parsedDates, err := sortDates(format, input...)

	if err != nil {
		log.Fatal("Wrong date format provided")
	}

	fmt.Println(parsedDates)

}

type dates []time.Time

func (s dates) Less(i, j int) bool { return s[i].Before(s[j]) }
func (s dates) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s dates) Len() int           { return len(s) }

func sortDates(format string, dateSlice ...string) ([]string, error) {
	//Should sort dates in ascending order (oldest to recent)
	//This func should parse and convert each string to a time.Time

	var parsedDates dates = make([]time.Time, 0, len(dateSlice))

	for i := 0; i < len(dateSlice); i++ {
		temp, err := time.Parse(format, dateSlice[i])
		if err != nil {
			return []string{}, err
		}
		parsedDates = append(parsedDates, temp)
	}

	sort.Sort(parsedDates)

	outputs := make([]string, 0, len(dateSlice))

	for _, output := range parsedDates {
		outputs = append(outputs, output.Format(format))
	}

	return outputs, nil
}
