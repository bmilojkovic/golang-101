# Assignment 2 - Eurovision score tracker

You are building a live score tracker for the Eurovision Song Contest. Countries compete by receiving scores from multiple judges, and your program should track those scores and report the standings at any time.

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

Judges can only award the following point values: `1, 2, 3, 4, 5, 6, 7, 8, 10, 12`. Your program should reject any score that is not in this list and ask the user to try again.

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
