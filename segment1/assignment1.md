## Assignment 1 - The guessing game

Write a program that plays a number guessing game with the user.

The program should pick a random whole number between 1 and 100 as the secret number. The user gets 7 attempts to guess it.
On each turn, read a number from the user. Tell them whether their guess was too high, too low, or correct. If they guess correctly, congratulate them, tell them how many attempts it took, and end the game. If they use all 7 attempts without guessing correctly, tell them they lost and reveal the secret number.

Example output:
```
Guess a number (1-100): 70
Too high! Guesses used: 1/7
Guess a number (1-100): 10
Too low! Guesses used: 2/7
Guess a number (1-100): 22
Correct! You got it in 3 guesses.
```
