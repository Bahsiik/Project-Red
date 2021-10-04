package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Affiche du menu de sélection
func Menu() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("--- Menu Principal ---")
	fmt.Println("A - Information du personnage")
	fmt.Println("B - Accéder à l'inventaire du personnage")
	fmt.Println("C - Accéder aux sorts du personnage")
	fmt.Println("D - Accéder au marchand")
	fmt.Println("E - Quitter le jeu")
	fmt.Println()

	// Lecture choix de menu
	for {
		fmt.Print("-> ")
		textmenu, _ := reader.ReadString('\n')
		// convert CRLF to LF
		textmenu = strings.Replace(textmenu, "\r\n", "", -1)

		switch textmenu {
		case "A":
			// Affichage infos perso
			P1.DisplayInfo()
			AccessInfo()
			// Affichage inventaire perso + accès à ce dernier
		case "B":
			P1.DisplayInventory()
			P1.AccessInventory()
			// Affichage des sorts + accès à ces derniers
		case "C":
			P1.DisplaySkill()
			P1.AccessSkill()
			// Affichage inventaire marchand + accès à ce dernier
		case "D":
			Marchand.DisplayInvMarchand()
			P1.AccessInvMarchand()
			// Sortie du jeu
		case "E":
			Exit()
		}
	}
}

// Commande sortie du jeu
func Exit() {
	fmt.Println("Vous allez quitter le jeu")
	fmt.Println()
	os.Exit(0)
}

// Commande retour au menu
func RetourMenu() {
	fmt.Println("Vous allez retourner au menu principal")
	fmt.Println()
	Menu()
}
