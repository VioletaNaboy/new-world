package main

import (
	"bufio"
	"fmt"
	"os"
)

type GameStatus struct {
	Location    string
	Items       []string
	IsBitten    bool
	IsExhausted bool
}

type Game struct {
	Status  GameStatus
	Scanner *bufio.Scanner
}

func main() {
	game := Game{
		Status: GameStatus{
			Location:   "cave",
			Items:      []string{"matches", "flashlight", "knife"},
			IsBitten:   false,
			IsExhausted: false,
		},
		Scanner: bufio.NewScanner(os.Stdin),
	}

	fmt.Println("Welcome to the game 'New World'.")
	fmt.Println("Your character wakes up at the entrance of a cave. He only remembers his name - Steven.")
	fmt.Println("Next to him is a backpack containing matches, a flashlight, and a knife.")
	fmt.Println("What will you do next?")

	for {
		switch game.Status.Location {
		case "cave":
			game.handleCave()
		case "forest":
			game.handleForest()
		case "camp":
			game.handleCamp()
		case "tent":
			game.handleTent()
		case "lake":
			game.handleLake()
		default:
			fmt.Println("Unknown location.")
			return
		}
	}
}

func (g *Game) handleCave() {
	fmt.Println("1. Enter the cave")
	fmt.Println("2. Follow the path to the forest")
	fmt.Print("> ")

	g.Scanner.Scan()
	choice := g.Scanner.Text()

	if choice == "1" {
		g.Status.Location = "cave"
		fmt.Println("It's too dark in the cave, you decide to exit.")
	} else if choice == "2" {
		g.Status.Location = "forest"
	} else {
		fmt.Println("Invalid choice, please try again.")
	}
}

func (g *Game) handleForest() {
	fmt.Println("You walk through the forest and encounter the body of a strange animal.")
	fmt.Println("1. Examine the body")
	fmt.Println("2. Continue walking")
	fmt.Println("3. Look around for water")
	fmt.Print("> ")

	g.Scanner.Scan()
	choice := g.Scanner.Text()

	if choice == "1" {
		fmt.Println("You find nothing of interest and decide to move on.")
		g.Status.Location = "camp"
	} else if choice == "2" {
		g.Status.Location = "camp"
	} else if choice == "3" {
		g.Status.Location = "lake"
	} else {
		fmt.Println("Invalid choice, please try again.")
	}
}

func (g *Game) handleCamp() {
	fmt.Println("You arrive at an abandoned camp. You are exhausted.")
	fmt.Println("1. Rest in the nearest tent")
	fmt.Println("2. Keep moving")
	fmt.Println("3. Search for supplies")
	fmt.Print("> ")

	g.Scanner.Scan()
	choice := g.Scanner.Text()

	if choice == "1" {
		g.Status.Location = "tent"
		g.Status.IsExhausted = true
	} else if choice == "2" {
		if g.Status.IsExhausted {
			fmt.Println("You are too exhausted and collapse unconscious. The game is over.")
			return
		}
		g.Status.Location = "forest"
	} else if choice == "3" {
		fmt.Println("You find some canned food and a bottle of water. You feel a bit more energized.")
		g.Status.IsExhausted = false
	} else {
		fmt.Println("Invalid choice, please try again.")
	}
}

func (g *Game) handleTent() {
	fmt.Println("In the nearest tent, you find a safe with a two-number combination lock.")
	fmt.Println("1. Try to open the safe")
	fmt.Println("2. Leave the tent")
	fmt.Print("> ")

	g.Scanner.Scan()
	choice := g.Scanner.Text()

	if choice == "1" {
		fmt.Println("Enter the code (two numbers separated by a space):")
		fmt.Print("> ")

		g.Scanner.Scan()
		code := g.Scanner.Text()
		if code == "17 17" {
			fmt.Println("The safe opens, but a large insect crawls out and bites your hand. You pass out.")
			g.Status.IsBitten = true
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

func (g *Game) handleLake() {
	fmt.Println("You find a small lake with clear water.")
	fmt.Println("1. Drink the water")
	fmt.Println("2. Rest by the lake")
	fmt.Println("3. Go back to the forest")
	fmt.Print("> ")

	g.Scanner.Scan()
	choice := g.Scanner.Text()

	if choice == "1" {
		fmt.Println("The water is refreshing, and you feel reinvigorated.")
		g.Status.IsExhausted = false
		g.Status.Location = "forest"
	} else if choice == "2" {
		fmt.Println("You rest by the lake and feel your strength returning.")
		g.Status.IsExhausted = false
	} else if choice == "3" {
		g.Status.Location = "forest"
	} else {
		fmt.Println("Invalid choice, please try again.")
	}
}
