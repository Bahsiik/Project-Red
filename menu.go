package main

import (
	"fmt"
	"os"
	"time"
)

func (p *Personnage) Menu() { // Affiche du menu de sélection
	fmt.Println("------ Menu Principal ------")
	fmt.Println()
	fmt.Println("1 - Gérer le personnage")
	fmt.Println("2 - Faire des achats")
	fmt.Println("3 - Aller chercher un adversaire")
	fmt.Println()
	fmt.Println("0 - Quitter le jeu")
	fmt.Println()
	for { // Lecture choix de menu
		textmenu := Input()
		switch textmenu {
		case "1":
			EffacerTerminal()
			p.Home()
		case "2":
			EffacerTerminal()
			p.Achats()
		case "3": // Lancement du combat
			EffacerTerminal()
			fmt.Println("ooooo Qui voulez-vous affrontez ? ooooo")
			fmt.Println()
			fmt.Println("1 - Gobelin (Easy)")
			fmt.Println("2 - Licorne (Medium)")
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
				p.Fight(&Gobelin, GobelinPattern)
			case "2":
				EffacerTerminal()
				p.Fight(&Licorne, LicornePattern)
			case "3":
				EffacerTerminal()
				p.Fight(&Dragon, DragonPattern)
			case "4":
				EffacerTerminal()
				p.Fight(&Alan, AlanPattern)
			case "5":
				EffacerTerminal()
				p.Fight(&Lucas, LucasPattern)
			case "0":
				EffacerTerminal()
				p.Menu()
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
	P1.Menu()
}

func (p *Personnage) Home() {
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
		p.DisplayInfo()
		p.AccessInfo()
	case "2": // Affichage inventaire perso + accès à ce dernier
		EffacerTerminal()
		p.DisplayInventory()
		p.AccessInventory()
	case "3": // Affichage de l'equipement du personnage
		EffacerTerminal()
		p.DisplayEquipment()
		p.AccessEquipment()
	case "4": // Affichage des sorts + accès à ces derniers
		EffacerTerminal()
		p.DisplaySkill()
		p.AccessSkill()
	case "0":
		EffacerTerminal()
		RetourMenu()
	default:
		fmt.Println("-------- Commande invalide --------")
		fmt.Println()
		time.Sleep(1 * time.Second)
		p.Home()
	}
}

func (p *Personnage) Achats() {
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
		p.AccessInvMarchand()
	case "2": // Affichage Forge
		EffacerTerminal()
		Forgeron.DisplayInvForgeron()
		p.AccessInvForgeron()
	case "0":
		EffacerTerminal()
		RetourMenu()
	default:
		fmt.Println("-------- Commande invalide --------")
		fmt.Println()
		time.Sleep(1 * time.Second)
		p.Achats()
	}
}
