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
	fmt.Println("C - Quitter le jeu")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\r\n", "", -1)

		if strings.Compare("A", text) == 0 {
			P1.DisplayInfo()
		} else if strings.Compare("B", text) == 0 {
			P1.AccessInventory()
			fmt.Println("Que va faire", P1.nom, " ?")
		} else if strings.Compare("Potion", text) == 0 {
			P1.TakePot()
		} else if strings.Compare("C", text) == 0 {
			fmt.Println("Vous allez quitter le jeu")
			os.Exit(0)
		}

	}
}
