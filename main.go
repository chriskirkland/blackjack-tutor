package main

import (
	"./blackjack"
	"./blackjackui"
	"math/rand"
	"time"
)

func randomize() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	// seed RNG
	randomize()

	// setup UI
	ui := new(blackjackui.ShellUI)

	gameCount := 0
	gamesBeforeShuffle := 5
	var myGame blackjack.Game
	for {
		// shuffle up and deal
		if gameCount%gamesBeforeShuffle == 0 {
			myGame = blackjack.NewGame()
		}
		myGame.NewHand()

		// player decisions
		var input string
		for myGame.PlayerCanHit() && input != "s" {
			ui.Redraw(myGame)
			input := ui.PromptUser("What do you want to do (h/s)?")
			if input == "h" {
				// player gets and card and loop
				myGame.DealPlayer()
			} else if input == "s" {
				break
			} else {
				panic("Invalid input for player action!")
			}
		}

		// dealer does stuff
		myGame.GoDealer()
		ui.Redraw(myGame)

		// decide winners
		switch myGame.GetWinner() {
		case blackjack.StatePlayerWins:
			ui.WinnerScreen()
		case blackjack.StateDealerWins:
			ui.LoserScreen()
		case blackjack.StatePush:
			ui.PushScreen()
		}

		// continue...?
		ui.QContinue()
		gameCount++
	}
}
