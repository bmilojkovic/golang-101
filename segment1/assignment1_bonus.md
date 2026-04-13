## Assignment 1 (advanced) — Guessing game tournament
Write a program that plays multiple rounds of the guessing game and keeps a score across rounds.

Each round works the same as before - a random number between 1 and 100, up to 7 guesses, too high/too low feedback. However, the scoring system rewards efficiency: the fewer guesses a player uses, the more points they earn.
Scoring per round:

Guess it in 1 attempt: 10 points
Guess it in 2 attempts: 7 points
Guess it in 3 attempts: 5 points
Guess it in 4 attempts: 3 points
Guess it in 5-7 attempts: 1 point
Fail to guess it: 0 points

The game runs for 5 rounds. After each round, tell the player how many points they earned and their running total. After all 5 rounds, print a final summary showing their total score, their best round (fewest guesses on a successful round), and a performance rating:

35-50 points: "Incredible"
20-34 points: "Solid"
10-19 points: "Room to improve"
0-9 points: "Better luck next time"
