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

func main() {
	printIntro()
	var in1 = getInput(prompt1)
	var in2 = getInput(prompt2)
	var in3 = getInput(prompt3)
	var in4 = getInput(prompt4)
	var in5 = getInput(prompt5)

	var score = getScore(in1, in2, in3, in4, in5)
	fmt.Println("score:", score)

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

func getScore(in1 int, in2 int, in3 int, in4 int, in5 int) int {
	var score int = 100 // max score
	score -= (2 * abs(ans1-in1))
	score -= (2 * abs(ans2-in2))
	score -= (2 * abs(ans3-in3))
	score -= (2 * abs(ans4-in4))
	score -= (2 * abs(ans5-in5))

	return score
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
