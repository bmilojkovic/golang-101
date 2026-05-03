/*
Write a program that prints the numbers from 1 to 100, but with a twist:

If the number is divisible by 3, print Fizz next to the bumber
If the number is divisible by 5, print Buzz next to the number
If the number is divisible by both 3 and 5, print FizzBuzz next to the number
When the loop is done, print a summary line that shows how many Fizz, Buzz, and FizzBuzz numbers appeared. For example:

Fizz: 27, Buzz: 14, FizzBuzz: 6
*/

package main

import "fmt"

func main() {

	var fizz, buzz, fizzbuzz int = 0, 0, 0
	var n_0, n int = 1, 100

	/*
		>> Extra work: modify the program so it can work on any range of numbers, not only 1-100.
		fmt.Print("Upisi donju granicu: ")
		fmt.Scan(&n_0)
		fmt.Print("Upisi gornju granicu: ")
		fmt.Scan(&n)
	*/

	for i := n_0; i <= n; i++ {
		switch {
		case i%15 == 0:
			fmt.Printf("FizzBuzz %d\n", i)
			fizzbuzz += 1
		case i%3 == 0:
			fmt.Printf("Fizz %d\n", i)
			fizz += 1
		case i%5 == 0:
			fmt.Printf("Buzz %d\n", i)
			buzz += 1
		default:
			fmt.Println(i)
		}
	}

	fmt.Println()

	fmt.Printf("Broj Fizz brojeva: %d\n", fizz)
	fmt.Printf("Broj Buzz brojeva: %d\n", buzz)
	fmt.Printf("Broj FizzBuzz brojeva: %d\n\n", fizzbuzz)
}
