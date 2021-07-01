package main

import (
	"fmt"
	"strconv"
)

func CalculateFrameScore(frameVal string) int {
	frameScore := SumRolls(frameVal)
	return frameScore
}

func GetUserInput(frameNo int) string {
	var frameVal string
	fmt.Print("Enter Frame " + strconv.Itoa(frameNo) + ": ")
	fmt.Scanln(&frameVal)
	return frameVal
}

func SumRolls(input string) int {
	result := 0
	for i := 0; i < len(input); i++ {
		frame := string(input[i])
		switch frame {
		case "/":
			if result+10 <= 20 {
				result += 10
			}
		case "X":
			if result+10 <= 30 {
				result += 10
			}
		case "1", "2", "3", "4", "5", "6", "7", "8", "9", "0":
			slice, err := strconv.Atoi(frame)
			if err == nil {
				result += slice
			}
		case ",":
		default:
			fmt.Println("               Skipping unrecognized character: " + frame)
		}
	}
	return result
}

func printFinalScore(val int) {
	fmt.Println("")
	fmt.Println("----------------------")
	fmt.Println("[ Final Score ]: [ " + strconv.Itoa(val) + " ]")
	fmt.Println("----------------------")
}

func printGoodbye() {
	fmt.Println("")
	fmt.Println("Thanks for playing! ")
	fmt.Println("")
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
		input := GetUserInput(f)
		currentFrameScore := CalculateFrameScore(input)
		currentScore += currentFrameScore
		fmt.Println("           [" + strconv.Itoa(f) + "] Total Score: " + strconv.Itoa(currentScore))
	}
	printFinalScore(currentScore)
	printGoodbye()
}
