package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100)

	for {
		var guess int
		fmt.Println("Enter a guess: ")
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Println("Inavlid guess: err:", err)
			continue
		}

		if guess > target {
			fmt.Println("Too high!")
			continue
		}

		if guess < target {
			fmt.Println("Too low!")
			continue
		}

		if guess == target {
			fmt.Println("You win!")
			break
		}

	}
}
