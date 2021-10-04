package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Menu() { // Affiche du menu de sélection
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("--- Menu Principal ---")
	fmt.Println("A - Information du personnage")
	fmt.Println("B - Accéder à l'inventaire du personnage")
	fmt.Println("C - Accéder aux sorts du personnage")
	fmt.Println("D - Accéder au marchand")
	fmt.Println("E - Accéder au Forgeron")
	fmt.Println("F - Quitter le jeu")
	fmt.Println()

	for { // Lecture choix de menu
		fmt.Print("-> ")
		textmenu, _ := reader.ReadString('\n')
		textmenu = strings.Replace(textmenu, "\r\n", "", -1)

		switch textmenu {
		case "A": // Affichage infos perso
			P1.DisplayInfo()
			AccessInfo()
		case "B": // Affichage inventaire perso + accès à ce dernier
			P1.DisplayInventory()
			P1.AccessInventory()
		case "C": // Affichage des sorts + accès à ces derniers
			P1.DisplaySkill()
			P1.AccessSkill()
		case "D": // Affichage inventaire marchand + accès à ce dernier
			Marchand.DisplayInvMarchand()
			P1.AccessInvMarchand()
		case "E": // Affichage Forge
			Forgeron.DisplayInventory()
			P1.AccessInvForegeron()
		case "F": // Sortie du jeu
			Exit()
		}
	}
}

func Exit() { // Commande sortie du jeu
	fmt.Println("Vous allez quitter le jeu")
	fmt.Println()
	os.Exit(0)
}

func RetourMenu() { // Commande retour au menu
	fmt.Println("Vous allez retourner au menu principal")
	fmt.Println()
	Menu()
}
