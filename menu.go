package main

import (
	"fmt"
	"os"
	"time"
)

func Menu() { // Affiche du menu de sélection
	fmt.Println("--- Menu Principal ---")
	fmt.Println("A - Information du personnage")
	fmt.Println("B - Accéder à l'inventaire du personnage")
	fmt.Println("C - Accéder à l'équipement du personnage")
	fmt.Println("D - Accéder aux sorts du personnage")
	fmt.Println("E - Accéder au marchand")
	fmt.Println("F - Accéder au Forgeron")
	fmt.Println("G - Combat d'entrainement")
	fmt.Println("H - Quitter le jeu")
	fmt.Println()

	for { // Lecture choix de menu
		textmenu := Input()

		switch textmenu {
		case "A": // Affichage infos perso
			EffacerTerminal()
			P1.DisplayInfo()
			AccessInfo()
		case "B": // Affichage inventaire perso + accès à ce dernier
			EffacerTerminal()
			P1.DisplayInventory()
			P1.AccessInventory()
		case "C": // Affichage de l'equipement du personnage
			EffacerTerminal()
			P1.DisplayEquipment()
			P1.AccessEquipment()
		case "D": // Affichage des sorts + accès à ces derniers
			EffacerTerminal()
			P1.DisplaySkill()
			P1.AccessSkill()
		case "E": // Affichage inventaire marchand + accès à ce dernier
			EffacerTerminal()
			Marchand.DisplayInvMarchand()
			P1.AccessInvMarchand()
		case "F": // Affichage Forge
			EffacerTerminal()
			Forgeron.DisplayInventory()
			P1.AccessInvForgeron()
		case "G": // Lancement du combat d'entrainement
			EffacerTerminal()
			TrainingFight(&P1, &Gobelin)
		case "H": // Sortie du jeu
			Exit()
		}
	}
}

func Exit() { // Commande sortie du jeu
	fmt.Println("Vous allez quitter le jeu")
	time.Sleep(1 * time.Second)
	fmt.Println()
	EffacerTerminal()
	os.Exit(0)
}

func RetourMenu() { // Commande retour au menu
	fmt.Println("Vous allez retourner au menu principal")
	time.Sleep(1 * time.Second)
	fmt.Println()
	EffacerTerminal()
	Menu()
}
