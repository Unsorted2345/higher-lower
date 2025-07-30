package main

import (
	"fmt"
	"math/rand"

	"github.com/mitchellh/colorstring"
)

func main() {
	var (
		numberToGuess     int = rand.Intn(100)
		Guess             int
		Guesses           int    = 0
		PlayAgain         string = "yes"
		Games             int    = 0
		allTimeGuesses    int    = 0
		allTimeGuessesWon int    = 0
		roundsWon         int    = 0
		difficulty        int    = 0
		chances           int    = 0
		roundWon          bool
		RoundAvarageTrue  float64 = 0.0
	)

	colorstring.Println("[yellow]Welcome to higher lower!\n[white]A game that makes you adicted!")

	for PlayAgain == "yes" || PlayAgain == "Yes" || PlayAgain == "YES" || PlayAgain == "y" || PlayAgain == "Y" {
		Games++
		colorstring.Println("[blue]On what difficulty do you want to play on?")
		colorstring.Println("[white]1. [green]Easy: 10 guesses")
		colorstring.Println("[white]2. [yellow]Medium: 7 guesses")
		colorstring.Println("[white]3. [red]Hard: 4 guesses")
		colorstring.Println("[white]4. [green]I[yellow]n[red]s[blue]a[green]n[yellow]e: [red]1 guess")
		for difficulty < 1 || difficulty > 4 {
			fmt.Scan(&difficulty)
			switch difficulty {
			case 1:
				chances = 10
			case 2:
				chances = 7
			case 3:
				chances = 4
			case 4:
				chances = 1
			default:
				colorstring.Println("[red]Invalid difficulty. Please try again.")
				continue
			}
		}

		Guesses = 0
		numberToGuess = rand.Intn(100)
		colorstring.Print("[blue]Guess a number between 0 and 100: ")
		fmt.Scan(&Guess)
		Guesses++
		allTimeGuesses++

		for Guess != numberToGuess && Guesses < chances {
			if Guess < numberToGuess {
				colorstring.Println("[red]Too low")
			} else if Guess > numberToGuess {
				colorstring.Println("[red]Too high")
			}
			fmt.Scan(&Guess)
			Guesses++
			allTimeGuesses++
			if Guess == numberToGuess {
				colorstring.Printf("[white]You guessed the number [green]%v [white]after %v guesses!\n", numberToGuess, Guesses)
				roundWon = true
				roundsWon++
			}
		}

		if roundWon == true {
			allTimeGuessesWon += Guesses

			RoundAvarageTrue = float64(allTimeGuessesWon) / float64(roundsWon)
		} else {
			colorstring.Printf("[white]the number was [green]%v\n", numberToGuess)
		}

		colorstring.Printf("[white]You played [green]%v [white]games. [green]%v [white]of these where won with an average of [green]%v [white]guesses per game.\n", Games, roundsWon, RoundAvarageTrue)
		colorstring.Print("[white]Would you like to play again? (yes/no): ")
		fmt.Scan(&PlayAgain)

		Guesses = 0
		difficulty = 0
		roundWon = false
	}
}
