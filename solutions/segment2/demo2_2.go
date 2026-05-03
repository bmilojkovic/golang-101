/*
You are building a score tracking program for a pub quiz night.
Three teams compete across multiple rounds, and you need to record their scores and report their performance.

Your program should store each team's round scores using a map where the key is the team name and the value is
a slice of integers, one per round. Write a function `printResults` that takes the scores map and prints the following for each team:

* Their score in each round
* Their total score across all rounds
* Their best single round score
* Their average score per round

After printing the per-team breakdown, the function should also print the overall winner — the team with the highest total score.
In main, call `printResults` once with the initial four rounds of scores.
Then add a bonus fifth round for all three teams and call `printResults` again to show the updated standings.

**Sample scores to use:**

|Team name |Round 1|Round 2|Round 3|Round 4|
|----------|-------|-------|-------|-------|
|Team Alpha| 8     | 5     | 9     | 6     |
|Team Beta | 4     | 7     | 7     | 10    |
|Team Gamma| 6     | 6     | 4     | 8     |

*/

package main

import "fmt"

func printResults(scores map[string][]int) {
	winnerName := ""
	winnerScore := 0

	for name, score := range scores {
		fmt.Println(name, ": ")
		// score in each round
		total := 0
		best := 0
		for i, val := range score {
			fmt.Println("Round ", i+1, ": ", val)
			total += val
			if val > best {
				best = val

			}
		}

		if total > winnerScore {
			winnerScore = total
			winnerName = name
		}

		// total score actoss all rounds
		fmt.Println("Total score: ", total)

		// best single round store
		fmt.Println("Best single round score: ", best)

		// avg score per round
		fmt.Println("Average score per round: ", float64(total)/float64(len(score)))

		fmt.Println()
	}
	fmt.Println("Best team is ", winnerName, ", with score: ", winnerScore)
}

func main() {
	t1_scores := []int{8, 5, 9, 6}
	t2_scores := []int{4, 7, 7, 10}
	t3_scores := []int{6, 6, 4, 8}

	scores := map[string][]int{
		"Team Alpha": t1_scores,
		"Team Beta":  t2_scores,
		"Team Gamma": t3_scores,
	}

	printResults(scores)

	fmt.Println()
	fmt.Println("-------- ROUND 5 ---------")
	fmt.Println()

	scores["Team Alpha"] = append(scores["Team Alpha"], 2)
	scores["Team Beta"] = append(scores["Team Beta"], 3)
	scores["Team Gamma"] = append(scores["Team Gamma"], 10)
	printResults(scores)

}
