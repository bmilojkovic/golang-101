/* Assignment 2 - Eurovision score tracker

You are building a live score tracker for the Eurovision Song Contest.
Countries compete by receiving scores from multiple judges, and your program should track those scores and report the standings at any time.

Your program starts with the following countries already loaded:

| Country  | Initial scores |
|----------|----------------|
| Sweden   | 10, 7, 12      |
| Ukraine  | 12, 8, 10      |
| Portugal | 6, 12, 8       |
| Norway   | 8, 5, 7        |

The user can then interact with the program through a menu:

```
1. Add a country
2. Add a score for a country
3. Print standings
4. Quit
```

## Scoring rules

Judges can only award the following point values: `1, 2, 3, 4, 5, 6, 7, 8, 10, 12`.
Your program should reject any score that is not in this list and ask the user to try again.

## Standings

Standings should show for each country:

- Their scores from each judge
- Total points
- Average score
- Highest single score

After the per-country breakdown, print the current leader — the country with the highest total.

## Functions to consider

- A function that checks whether a score is valid
- A function that prints the full standings
- A function that prints a single country's results

You may find you need more functions than these, or that you want to combine or split them differently — that is completely fine.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var validScores = []int{1, 2, 3, 4, 5, 6, 7, 8, 10, 12}

func isValidScore(score int) bool {
	for _, v := range validScores {
		if score == v {
			return true
		}
	}
	return false
}

func printCountry(name string, scores []int) {
	fmt.Println(name + ":")
	total := 0
	best := 0
	for i, s := range scores {
		fmt.Println("  Judge", i+1, ":", s)
		total += s
		if s > best {
			best = s
		}
	}
	fmt.Println("  Total:", total)
	fmt.Printf("  Average: %.2f\n", float64(total)/float64(len(scores)))
	fmt.Println("  Best score:", best)
	fmt.Println()
}

func printStandings(scores map[string][]int) {
	leaderName := ""
	leaderScore := 0

	for name, s := range scores {
		printCountry(name, s)

		total := 0
		for _, v := range s {
			total += v
		}
		if total > leaderScore {
			leaderScore = total
			leaderName = name
		}
	}
	fmt.Println("Current leader:", leaderName, "with", leaderScore, "points")
}

func main() {
	scores := map[string][]int{
		"Sweden":   {10, 7, 12},
		"Ukraine":  {12, 8, 10},
		"Portugal": {6, 12, 8},
		"Norway":   {8, 5, 7},
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n1. Add a country")
		fmt.Println("2. Add a score for a country")
		fmt.Println("3. Print standings")
		fmt.Println("4. Quit")
		fmt.Print("\nChoice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Country name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			if _, exists := scores[name]; exists {
				fmt.Println("Country already exists.")
			} else {
				scores[name] = []int{}
				fmt.Println("Added", name)
			}

		case 2:
			fmt.Print("Country name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			if _, exists := scores[name]; !exists {
				fmt.Println("Country not found.")
				break
			}
			for {
				fmt.Print("Score to add: ")
				var s int
				fmt.Scan(&s)
				if isValidScore(s) {
					scores[name] = append(scores[name], s)
					fmt.Println("Score added.")
					break
				}
				fmt.Println("Invalid score. Must be one of: 1 2 3 4 5 6 7 8 10 12")
			}

		case 3:
			fmt.Println()
			printStandings(scores)

		case 4:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}
