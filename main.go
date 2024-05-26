package main

import (
	"bufio"
	"fmt"
	"os"
)

type GameStatus struct {
	Location   string
	Items      []string
	IsBitten   bool
	IsExhausted bool
}

func main() {
	status := GameStatus{
		Location:   "cave",
		Items:      []string{"matches", "flashlight", "knife"},
		IsBitten:   false,
		IsExhausted: false,
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the game 'New World'.")
	fmt.Println("Your character wakes up at the entrance of a cave. He only remembers his name - Steven.")
	fmt.Println("Next to him is a backpack containing matches, a flashlight, and a knife.")
	fmt.Println("What will you do next?")

	for {
		switch status.Location {
		case "cave":
			handleCave(&status, scanner)
		case "forest":
			handleForest(&status, scanner)
		case "camp":
			handleCamp(&status, scanner)
		case "tent":
			handleTent(&status, scanner)
		case "lake":
			handleLake(&status, scanner)
		default:
			fmt.Println("Unknown location.")
			return
		}
	}
}

func handleCave(status *GameStatus, scanner *bufio.Scanner) {
	fmt.Println("1. Enter the cave")
	fmt.Println("2. Follow the path to the forest")
	fmt.Print("> ")

	scanner.Scan()
	choice := scanner.Text()

	if choice == "1" {
		status.Location = "cave"
		fmt.Println("It's too dark in the cave, you decide to exit.")
	} else if choice == "2" {
		status.Location = "forest"
	} else {
		fmt.Println("Invalid choice, please try again.")
	}
}

func handleForest(status *GameStatus, scanner *bufio.Scanner) {
	fmt.Println("You walk through the forest and encounter the body of a strange animal.")
	fmt.Println("1. Examine the body")
	fmt.Println("2. Continue walking")
	fmt.Println("3. Look around for water")
	fmt.Print("> ")

	scanner.Scan()
	choice := scanner.Text()

	if choice == "1" {
		fmt.Println("You find nothing of interest and decide to move on.")
		status.Location = "camp"
	} else if choice == "2" {
		status.Location = "camp"
	} else if choice == "3" {
		status.Location = "lake"
	} else {
		fmt.Println("Invalid choice, please try again.")
	}
}

func handleCamp(status *GameStatus, scanner *bufio.Scanner) {
	fmt.Println("You arrive at an abandoned camp. You are exhausted.")
	fmt.Println("1. Rest in the nearest tent")
	fmt.Println("2. Keep moving")
	fmt.Println("3. Search for supplies")
	fmt.Print("> ")

	scanner.Scan()
	choice := scanner.Text()

	if choice == "1" {
		status.Location = "tent"
		status.IsExhausted = true
	} else if choice == "2" {
		if status.IsExhausted {
			fmt.Println("You are too exhausted and collapse unconscious. The game is over.")
			return
		}
		status.Location = "forest"
	} else if choice == "3" {
		fmt.Println("You find some canned food and a bottle of water. You feel a bit more energized.")
		status.IsExhausted = false
	} else {
		fmt.Println("Invalid choice, please try again.")
	}
}

func handleTent(status *GameStatus, scanner *bufio.Scanner) {
	fmt.Println("In the nearest tent, you find a safe with a two-number combination lock.")
	fmt.Println("1. Try to open the safe")
	fmt.Println("2. Leave the tent")
	fmt.Print("> ")

	scanner.Scan()
	choice := scanner.Text()

	if choice == "1" {
		fmt.Println("Enter the code (two numbers separated by a space):")
		fmt.Print("> ")

		scanner.Scan()
		code := scanner.Text()
		if code == "17 17" {
			fmt.Println("The safe opens, but a large insect crawls out and bites your hand. You pass out.")
			status.IsBitten = true
			fmt.Println("The game is over.")
			return
		} else {
			fmt.Println("Incorrect code. The safe does not open.")
		}
	} else if choice == "2" {
		fmt.Println("You leave the tent and decide to rest outside.")
		fmt.Println("The game is over.")
		return
	} else {
		fmt.Println("Invalid choice, please try again.")
	}
}

func handleLake(status *GameStatus, scanner *bufio.Scanner) {
	fmt.Println("You find a small lake with clear water.")
	fmt.Println("1. Drink the water")
	fmt.Println("2. Rest by the lake")
	fmt.Println("3. Go back to the forest")
	fmt.Print("> ")

	scanner.Scan()
	choice := scanner.Text()

	if choice == "1" {
		fmt.Println("The water is refreshing, and you feel reinvigorated.")
		status.IsExhausted = false
		status.Location = "forest"
	} else if choice == "2" {
		fmt.Println("You rest by the lake and feel your strength returning.")
		status.IsExhausted = false
	} else if choice == "3" {
		status.Location = "forest"
	} else {
		fmt.Println("Invalid choice, please try again.")
	}
}
