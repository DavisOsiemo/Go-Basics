package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randNumberGenerator() string {
	min := 1
	max := 500
	// generate random number and print on console
	// fmt.Println(rand.Intn(max - min) + min)
	randNumber := rand.Intn(max-min) + min
	randString := strconv.Itoa(randNumber)

	// A slice of random numbers
	formats := []string{
		randString,
	}
	return formats[rand.Intn(len(formats))]
}

func randNameGenerator() string {
	formats := []string{
		"John Doe",
		"Alex Ferguson",
		"Jose Mourinho",
		"Ngolo Kante",
	}
	return formats[rand.Intn(len(formats))]
}

func main() {
	students := []int{1, 2, 3, 4, 5, 6}
	rewardees := students[0:3]

	for _, rewardee := range rewardees {
		rewardeeStrn := strconv.Itoa(rewardee)

		fmt.Println(rewardeeStrn + " Hi there " + randNameGenerator() + " your raffle number is " + randNumberGenerator())
		// fmt.Print(randPoints)
	}
}
