package gobowling

import (
	"fmt"
	"strconv"
)

func getFrameScore(frameNo int) int {
	// TODO - validate frameVal, repeat user query if bad
	var frameVal string
	fmt.Print("Enter Your Frame [" + strconv.Itoa(frameNo) + "]: ")
	fmt.Scanln(&frameVal)
	frameScore := Score(frameVal)
	return frameScore
}

func Score(input ...string) int {
	// TODO - scoring logic
	fmt.Print(input, " ")
	for _, frame := range input {
		// TODO - parse frame input and score
		fmt.Println(frame)
	}
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
	currentScore := 0
	printWelcome()
	for f := 1; f <= 10; f++ {
		currentFrameScore := getFrameScore(f)
		currentScore += currentFrameScore
		fmt.Println("Frame " + strconv.Itoa(f) + ": " + strconv.Itoa(currentScore))
	}
}
