package main

import (
	"fmt"
	"strconv"
	"strings"
)

const ans1 = 5
const ans2 = 1
const ans3 = 4
const ans4 = 5
const ans5 = 1

const prompt1 = "I enjoy watching arthouse films: "
const prompt2 = "I dislike sushi: "
const prompt3 = "I like playing board games: "
const prompt4 = "I would appreciate handmade crochet gifts: "
const prompt5 = "I am a morning person: "

const scoreThresholdMin = 5
const scoreThreshold1 = 65
const scoreThreshold2 = 85
const scoreThresholdMax = 100

// each degree away from desired answer will remove 5 points from the score
// example: desired answer = 5, given answer = 1 -> points to remove = 20 (4*5)
const pointsPerVariance = 5

func main() {
	printIntro()
	var in1 = getInput(prompt1)
	var in2 = getInput(prompt2)
	var in3 = getInput(prompt3)
	var in4 = getInput(prompt4)
	var in5 = getInput(prompt5)

	getTotalScore(in1, in2, in3, in4, in5)
}

func printIntro() {
	fmt.Printf("%s\n", strings.Repeat("*", 55))
	fmt.Printf("%32s\n", "Matchmaker")
	fmt.Println()
	fmt.Println("         Helping make connections since 2024")
	fmt.Printf("%s\n", strings.Repeat("*", 55))
	fmt.Println()
	fmt.Println("This program figures out if you and I are compatible.")
	fmt.Println("You will be given 5 prompts. To each prompt, answer 5")
	fmt.Println("if you strongly agree, 4 if you agree, 3 if you neither")
	fmt.Println("agree nor disagree, 2 if you disagree, and 1 if you")
	fmt.Println("strongly disagree.")
	fmt.Println()
}

func getInput(prompt string) int {
	var isValid = false
	var numInput = 0
	var strInput = ""

	for !isValid {
		fmt.Print(prompt)
		fmt.Scanln(&strInput)
		num, err := strconv.Atoi(strInput)
		if err != nil {
			fmt.Println("Please enter a number 1-5")

		} else if num < 1 || num > 5 {
			fmt.Println("Please enter 1-5.")

		} else {
			numInput = num
			isValid = true
		}
	}
	fmt.Println()
	return numInput
}

// adjusted method of determining final score, max score = 100%, min score = 5%
func getTotalScore(in1 int, in2 int, in3 int, in4 int, in5 int) {
	var totalScore int = 100 // max score
	score1 := (pointsPerVariance * abs(ans1-in1))
	score2 := (pointsPerVariance * abs(ans2-in2))
	score3 := (pointsPerVariance * abs(ans3-in3))
	score4 := (pointsPerVariance * abs(ans4-in4))
	score5 := (pointsPerVariance * abs(ans5-in5))

	scores := [5]int{score1, score2, score3, score4, score5}

	fmt.Printf("%s\n", strings.Repeat("*", 55))
	for i := 0; i < len(scores); i++ {
		totalScore -= scores[i]
		printPromptScore(i+1, scores[i])
	}

	fmt.Printf("\nTotal Score: %v%%\n", totalScore)
	printScoreSummary(totalScore)
}

func printPromptScore(promptNum int, score int) {
	fmt.Printf("Prompt %v Score: %v%%\n", promptNum, 100*(20-score)/20)
}

func printScoreSummary(total int) {
	if total == scoreThresholdMin {
		fmt.Println("Wow... we are total opposites.")
	} else if total == scoreThresholdMax {
		fmt.Println("Wow... are we the same person?")
	} else if total < scoreThreshold1 {
		fmt.Println("We don't seem to be very compatible.")
	} else if total < scoreThreshold2 {
		fmt.Println("We seem to have a bit in common... maybe we could be friends!")
	} else {
		fmt.Println("We have a lot in common! Are we best friends now?")
	}
	fmt.Println()
	fmt.Println("Thanks for using Matchmaker!")
	fmt.Printf("%s\n", strings.Repeat("*", 55))
	fmt.Println()
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
