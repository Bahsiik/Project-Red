package main

import (
	"fmt"
	"os"
	"time"
)

func Menu() { // Affiche du menu de sélection
	fmt.Println("------ Menu Principal ------")
	fmt.Println()
	fmt.Println("1 - Gérer le personnage")
	fmt.Println("2 - Faire des achats")
	fmt.Println("3 - Commencer le combat suivant")
	fmt.Println()
	fmt.Println("0 - Quitter le jeu")
	fmt.Println()
	for { // Lecture choix de menu
		textmenu := Input()
		switch textmenu {
		case "1":
			EffacerTerminal()
			Home()
		case "2":
			EffacerTerminal()
			Achats()
		case "3": // Lancement du combat d'entrainement
			EffacerTerminal()
			fmt.Println("ooooo Qui voulez-vous affrontez ? ooooo")
			fmt.Println()
			fmt.Println("1 - Goblin (Easy)")
			fmt.Println("2 - Licorn (Medium)")
			fmt.Println("3 - Dragon (Hard)")
			fmt.Println("4 - Alan (Extrême)")
			fmt.Println("5 - Lucas (Divin)")
			fmt.Println()
			fmt.Println("0 - Retour")
			fmt.Println()
			textfight := Input() // Choix opposant
			switch textfight {
			case "1":
				EffacerTerminal()
				TrainingFight(&P1, &Gobelin, GoblinPattern)
			case "2":
				EffacerTerminal()
				TrainingFight(&P1, &Licorn, LicornPattern)
			case "3":
				EffacerTerminal()
				TrainingFight(&P1, &Dragon, DragonPattern)
			case "4":
				EffacerTerminal()
				TrainingFight(&P1, &Alan, AlanPattern)
			case "5":
				EffacerTerminal()
				TrainingFight(&P1, &Lucas, LucasPattern)
			case "0":
				EffacerTerminal()
				Menu()
			}
		case "0": // Sortie du jeu
			EffacerTerminal()
			Exit()
		}
	}
}

func Exit() { // Commande sortie du jeu
	fmt.Println()
	fmt.Println("-------- Vous allez quitter le jeu --------")
	fmt.Println()
	time.Sleep(1 * time.Second)
	fmt.Println()
	EffacerTerminal()
	os.Exit(0)
}

func RetourMenu() { // Commande retour au menu
	fmt.Println()
	fmt.Println("-------- Vous allez retourner au menu principal --------")
	time.Sleep(1 * time.Second)
	fmt.Println()
	EffacerTerminal()
	Menu()
}

func Home() {
	EffacerTerminal()
	fmt.Println("-------- Que voulez-vous faire ? --------")
	fmt.Println()
	fmt.Println("1 - Information du personnage")
	fmt.Println("2 - Accéder à l'inventaire du personnage")
	fmt.Println("3 - Accéder à l'équipement du personnage")
	fmt.Println("4 - Accéder aux sorts du personnage")
	fmt.Println()
	fmt.Println("0 - Retour au menu principal")
	texthome := Input()
	switch texthome {
	case "1": // Affichage infos perso
		EffacerTerminal()
		P1.DisplayInfo()
		AccessInfo()
	case "2": // Affichage inventaire perso + accès à ce dernier
		EffacerTerminal()
		P1.DisplayInventory()
		P1.AccessInventory()
	case "3": // Affichage de l'equipement du personnage
		EffacerTerminal()
		P1.DisplayEquipment()
		P1.AccessEquipment()
	case "4": // Affichage des sorts + accès à ces derniers
		EffacerTerminal()
		P1.DisplaySkill()
		P1.AccessSkill()
	case "0":
		EffacerTerminal()
		RetourMenu()
	default:
		fmt.Println("-------- Commande invalide --------")
		fmt.Println()
		time.Sleep(1 * time.Second)
		Home()
	}
}

func Achats() {
	EffacerTerminal()
	fmt.Println("-------- Que voulez-vous faire ? --------")
	fmt.Println()
	fmt.Println("1 - Accéder au marchand")
	fmt.Println("2 - Accéder au Forgeron")
	fmt.Println()
	fmt.Println("0 - Retour au menu principal")
	textachats := Input()
	switch textachats {
	case "1": // Affichage inventaire marchand + accès à ce dernier
		EffacerTerminal()
		Marchand.DisplayInvMarchand()
		P1.AccessInvMarchand()
	case "2": // Affichage Forge
		EffacerTerminal()
		Forgeron.DisplayInvForgeron()
		P1.AccessInvForgeron()
	case "0":
		EffacerTerminal()
		RetourMenu()
	default:
		fmt.Println("-------- Commande invalide --------")
		fmt.Println()
		time.Sleep(1 * time.Second)
		Achats()
	}
}
