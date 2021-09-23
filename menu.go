package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Menu() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("A - Information du personnage")
	fmt.Println("B - Accéder à l'inventaire du personnage")
	fmt.Println("C - Accéder au marchand")
	fmt.Println("D - Quitter le jeu")
	fmt.Println("E - Suicide")
	fmt.Println("F - Tu es sûr ?")
	fmt.Println()

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\r\n", "", -1)

		switch text {
		case "A":
			P1.DisplayInfo()
		case "B":
			P1.DisplayInventory()
			P1.AccessInventory()
		case "C":
			Marchand.DisplayInvMarchand()
			P1.AccessInvMarchand()
		case "D":
			fmt.Println("Vous allez quitter le jeu")
			os.Exit(0)
		case "E":
			P1.Death()
		case "F":
			P1.PoisonPot()
		}
	}
}
