/*
Guessing game tournament

Write a program that plays multiple rounds of the guessing game and keeps a score across rounds.

Each round works the same as before - a random number between 1 and 100, up to 7 guesses, too high/too low feedback.
However, the scoring system rewards efficiency: the fewer guesses a player uses, the more points they earn.
Scoring per round:

Guess it in 1 attempt: 10 points
Guess it in 2 attempts: 7 points
Guess it in 3 attempts: 5 points
Guess it in 4 attempts: 3 points
Guess it in 5-7 attempts: 1 point
Fail to guess it: 0 points

The game runs for 5 rounds. After each round, tell the player how many points they earned and their running total.
After all 5 rounds, print a final summary showing their total score, their best round (fewest guesses on a successful round), and a performance rating:

35-50 points: "Incredible"
20-34 points: "Solid"
10-19 points: "Room to improve"
0-9 points: "Better luck next time"

*/

package main

import (
	"fmt"
	"math/rand/v2"
)

func round() (int, bool) {

	var num int = rand.IntN(100)
	var guess int
	var guessed bool = false

	var guess_num int = 8

	for i := 1; i <= 7 && !guessed; i++ {
		fmt.Println("Guess a number (1-100): ")
		fmt.Scan(&guess)

		switch {
		case guess < num:
			fmt.Printf("Too low! Guesses used: %d/7\n", i)
		case guess > num:
			fmt.Printf("Too high! Guesses used: %d/7\n", i)
		default:
			fmt.Printf("Correct! You got it in %d guesses.\n", i)
			guessed = true
			guess_num = i
		}
	}
	if !guessed {
		fmt.Println("You lost! The secret number was ", num)
	}

	return guess_num, guessed
}

func main() {

	fmt.Println("Welcome to the Guessing Game Tournament! Good luck!")
	points := 0
	best_points := 0
	best_guess_num := 0
	best_idx := 0

	for i := 1; i <= 5; i++ {
		fmt.Println("Round ", i)

		attempts, found := round()
		points_round := 0

		if !found {
			points_round = 0
		} else {
			switch {
			case attempts == 1:
				points_round = 10
			case attempts == 2:
				points_round = 7
			case attempts == 3:
				points_round = 5
			case attempts == 4:
				points_round = 3
			default:
				points_round = 1
			}
		}

		if points_round > best_points {
			best_points = points_round
			best_guess_num = attempts
			best_idx = i
		}

		points += points_round

		fmt.Println("Points in this round: ", points_round)
		fmt.Println("Running total: ", points)
		fmt.Println()
	}

	fmt.Printf("\nSummary:\nTotal score: %d\nBest round: Round %d, %d points (%d guesses)\n", points, best_idx, best_points, best_guess_num)
	fmt.Print("Performance rating: ")

	if points > 35 {
		fmt.Println("Incredible!")
	} else if points > 20 {
		fmt.Println("Solid")
	} else if points > 10 {
		fmt.Println("Room to improve")
	} else {
		fmt.Println("Better luck next time!")
	}
}
