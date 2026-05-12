/*
## Assignment 3 - Pet shelter

You are building a management system for a pet shelter. The shelter keeps track of the animals in its care and helps match them with new owners.

Your program should have two structs: a `Pet` and a `Shelter`.

A `Pet` has a name, a species, an age, and a status indicating whether they have been adopted.

A `Shelter` has a name and keeps a list of all the pets currently in its care.

To get started, populate your shelter with at least 5 creative and original pets of your own invention.

The user interacts with the shelter through a menu:
1. List available pets
2. Adopt a pet
3. Add a new pet
4. Find the oldest available pet
5. Quit

Behaviour:

- Listing available pets shows only pets that have not yet been adopted, with their full profile — name, species, age, and status.
- Adopting a pet takes a name as input, marks that pet as adopted, and prints a farewell message. If no pet with that name exists, or they are already adopted, print an appropriate message.
- Adding a new pet asks the user for a name, species, and age.
- Finding the oldest available pet prints the profile of the non-adopted pet with the highest age. If all pets have been adopted, print an appropriate message.

Methods to consider:

- A method on `Pet` to print its profile
- A method on `Pet` to mark it as adopted
- A method on `Shelter` to add a pet
- A method on `Shelter` to list available pets
- A method on `Shelter` to adopt a pet by name
- A method on `Shelter` to find the oldest available pet

Think carefully about which methods need a pointer receiver and which do not.

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func readLine(prompt string) string {
	fmt.Print(prompt)
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func readInt(prompt string) int {
	for {
		input := readLine(prompt)
		n, err := strconv.Atoi(input)
		if err == nil {
			return n
		}
		fmt.Println("Please enter a valid number.")
	}
}

type Pet struct {
	Name    string
	Species string
	Age     int
	Adopted bool
}

type Shelter struct {
	Name string
	Pets []Pet
}

func (p Pet) String() string {
	status := "available"
	if p.Adopted {
		status = "adopted"
	}
	return fmt.Sprintf("%s the %s | Age: %d | Status: %s",
		p.Name, p.Species, p.Age, status)
}

func (p *Pet) adopt() {
	p.Adopted = true
}

func (s *Shelter) addPet(pet Pet) {
	s.Pets = append(s.Pets, pet)
}

func (s *Shelter) listAvailable() {
	fmt.Printf("\n=== Available pets at %s ===\n", s.Name)
	found := false
	for _, pet := range s.Pets {
		if !pet.Adopted {
			fmt.Println(pet)
			found = true
		}
	}
	if !found {
		fmt.Println("No pets currently available.")
	}
	fmt.Println()
}

func (s *Shelter) adoptPet(name string) {
	for i, pet := range s.Pets {
		if pet.Name == name {
			if pet.Adopted {
				fmt.Printf("%s has already been adopted!\n\n", name)
				return
			}
			s.Pets[i].adopt()
			fmt.Printf("Congratulations! %s the %s has found a new home. Goodbye %s!\n\n",
				pet.Name, pet.Species, pet.Name)
			return
		}
	}
	fmt.Printf("No pet named %s found in the shelter.\n\n", name)
}

func (s *Shelter) oldestAvailable() {
	oldest := -1
	index := -1
	for i, pet := range s.Pets {
		if !pet.Adopted && pet.Age > oldest {
			oldest = pet.Age
			index = i
		}
	}
	if index == -1 {
		fmt.Println("All pets have been adopted!")
		return
	}
	fmt.Printf("\n=== Oldest available pet ===\n%s\n\n", s.Pets[index])
}

func main() {
	shelter := Shelter{
		Name: "Paws & Claws Sanctuary",
		Pets: []Pet{
			{Name: "Biscuit", Species: "three-legged corgi", Age: 5},
			{Name: "Tornado", Species: "anxious greyhound", Age: 3},
			{Name: "Professor Whiskers", Species: "judgmental persian cat", Age: 9},
			{Name: "Noodle", Species: "very long dachshund", Age: 2},
			{Name: "Gremlin", Species: "chaotic parrot", Age: 14},
		},
	}

	fmt.Printf("Welcome to %s!\n\n", shelter.Name)

	for {
		fmt.Println("1. List available pets")
		fmt.Println("2. Adopt a pet")
		fmt.Println("3. Add a new pet")
		fmt.Println("4. Find the oldest available pet")
		fmt.Println("5. Quit")

		choice := readInt("> ")
		fmt.Println()

		switch choice {
		case 1:
			shelter.listAvailable()

		case 2:
			name := readLine("Enter the name of the pet you want to adopt: ")
			fmt.Println()
			shelter.adoptPet(name)

		case 3:
			name := readLine("Pet name: ")
			species := readLine("Species: ")
			age := readInt("Age: ")
			fmt.Println()
			shelter.addPet(Pet{Name: name, Species: species, Age: age})
			fmt.Printf("%s the %s has been added to the shelter.\n\n", name, species)

		case 4:
			shelter.oldestAvailable()

		case 5:
			fmt.Println("Thanks for visiting. See you next time!")
			return

		default:
			fmt.Println("Invalid choice, please enter 1-5.\n")
		}
	}
}
