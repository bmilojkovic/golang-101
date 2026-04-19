## Slices and maps

You are building a score tracking program for a pub quiz night. Three teams compete across multiple rounds, and you need to record their scores and report their performance.

Your program should store each team's round scores using a map where the key is the team name and the value is a slice of integers, one per round. Write a function `printResults` that takes the scores map and prints the following for each team:

* Their score in each round
* Their total score across all rounds
* Their best single round score
* Their average score per round

After printing the per-team breakdown, the function should also print the overall winner — the team with the highest total score.
In main, call `printResults` once with the initial four rounds of scores. Then add a bonus fifth round for all three teams and call `printResults` again to show the updated standings.

**Sample scores to use:**

|Team name |Round 1|Round 2|Round 3|Round 4|
|----------|-------|-------|-------|-------|
|Team Alpha| 8     | 5     | 9     | 6     |
|Team Beta | 4     | 7     | 7     | 10    |
|Team Gamma| 6     | 6     | 4     | 8     |
