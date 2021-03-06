package blackjackui

import (
	"../blackjack"
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

type ShellUI struct{}

func (b *ShellUI) PromptUser(msg string) string {
	fmt.Printf("%s> ", msg)
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()
}

func (b *ShellUI) Redraw(g blackjack.Game) {
	clearScreen()
	fmt.Println(g)
}

func (b *ShellUI) WinnerScreen() {
	fmt.Println("\nPlayer Wins!!!\n")
}

func (b *ShellUI) LoserScreen() {
	fmt.Println("\nDealer Wins.\n")
}

func (b *ShellUI) PushScreen() {
	fmt.Println("\nPush.\n")
}

func (b *ShellUI) InsuranceWin() {
	fmt.Println("\nDealer has Blackjack. Insurance Pays!\n")
}

func (b *ShellUI) QContinue() {
	fmt.Print("Press ENTER to continue...")
	fmt.Scanln()
}
