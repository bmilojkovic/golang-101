/*
## Pointers and structs

You are building a simple player management system for a game. A player has a name, a health value (0–100), a score, and a level.
Write a program that models a single player and simulates a short sequence of events during a game session.

Your Player should support the following behaviours:

- Taking damage — health cannot go below 0
- Healing — health cannot go above 100
- Gaining score points
- Levelling up — increases the player's level by 1 and fully restores their health to 100
- Describing themselves — returns a formatted string showing all of their current stats

In main, create a player and run them through the following sequence of events:

- Print their starting stats
- They take 40 damage and gain 150 points
- Print their stats
- They heal 20 and gain 200 points
- Print their stats
- They take 999 damage (the game doesn't have to end here - we are just testing the Player struct)
- Print their stats
- They level up
- Print their final stats
*/
package main

import "fmt"

type Player struct {
	Name   string
	Health int
	Score  int
	Level  int
}

func (p *Player) takeDamage(dmg int) {
	p.Health -= dmg
	if p.Health < 0 {
		p.Health = 0
	}
}

func (p *Player) heal(extraHealth int) {
	p.Health = min(100, p.Health+extraHealth)
}

func (p *Player) increaseScore(increase int) {
	p.Score += increase
}

func (p *Player) levelUp() {
	p.Level++

	p.Health = 100
}

func (p Player) describe() string {
	return fmt.Sprintf("%s (lvl %d) HP: %d Score: %d", p.Name, p.Level, p.Health, p.Score)
}

func main() {
	player := Player{
		Name:   "Conan",
		Health: 100,
		Score:  5,
		Level:  9,
	}

	fmt.Printf("%s\n", player.describe())
	player.takeDamage(40)
	player.increaseScore(150)
	fmt.Printf("%s\n", player.describe())
	player.heal(20)
	player.increaseScore(200)
	fmt.Printf("%s\n", player.describe())
	player.takeDamage(999)
	fmt.Printf("%s\n", player.describe())
	player.levelUp()
	fmt.Printf("%s\n", player.describe())
}
