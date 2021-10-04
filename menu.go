package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Menu() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("--- Menu Principal ---")
	fmt.Println("A - Information du personnage")
	fmt.Println("B - Accéder à l'inventaire du personnage")
	fmt.Println("C - Accéder au marchand")
	fmt.Println("D - Quitter le jeu")
	fmt.Println()

	for {
		fmt.Print("-> ")
		textmenu, _ := reader.ReadString('\n')
		// convert CRLF to LF
		textmenu = strings.Replace(textmenu, "\r\n", "", -1)

		switch textmenu {
		case "A":
			P1.DisplayInfo()
			AccessInfo()
		case "B":
			P1.DisplayInventory()
			P1.AccessInventory()
		case "C":
			Marchand.DisplayInvMarchand()
			P1.AccessInvMarchand()
		case "D":
			Exit()
		}
	}
}

func Exit() {
	fmt.Println("Vous allez quitter le jeu")
	fmt.Println()
	os.Exit(0)
}

func RetourMenu() {
	fmt.Println("Vous allez retourner au menu principal")
	fmt.Println()
	Menu()
}
