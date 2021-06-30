package main

import (
	"fmt"
	"strconv"
)

func getFrameScore(frameNo int) int {
	// TODO - validate frameVal
	var frameVal string
	fmt.Print("Enter Your Frame [" + strconv.Itoa(frameNo) + "]: ")
	fmt.Scanln(&frameVal)
	frameScore := score(frameVal)
	return frameScore
}

func score(input ...string) int {
	// TODO - scoring logic
	result := 5
	return result
}

func printWelcome() {
	fmt.Println("")
	fmt.Println("---------------------- ")
	fmt.Println("Welcome to GO-bowling! ")
	fmt.Println("---------------------- ")
	fmt.Println("")
	fmt.Println("I'll help keep track of your bowling game.")
	fmt.Println("")
}

func main() {
	score := 0
	printWelcome()
	for f := 1; f <= 10; f++ {
		currentFrameScore := getFrameScore(f)
		score += currentFrameScore
		fmt.Println("Frame " + strconv.Itoa(f) + ": " + strconv.Itoa(score))
	}
}
