package main

import (
	"fmt"
	"time"
)

func (p Personnage) DisplayInfo() { // Affichage des informations du personnages
	fmt.Println("--- Informations du personnage ---")
	fmt.Println()
	fmt.Println("Nom --> ", p.nom)
	fmt.Println("Classe --> ", p.classe)
	fmt.Println("Niveau --> ", p.niveau)
	fmt.Println("HP --> ", p.hp, " / ", p.hpmax)
	fmt.Println("Mana --> ", p.mana, " / ", p.manamax)
	fmt.Println("Attaque --> ", p.atk)
	fmt.Println("Puissance --> ", p.puissance)
	fmt.Println("Initiative --> ", p.initiative)
	fmt.Println("Points  d'expérience --> ", p.exp, " / ", p.expmax)
	fmt.Println("Argent actuel --> ", p.money, " Caca d'or")
	fmt.Println()
	p.AccessInfo()
}

func (p *Personnage) AccessInfo() { // Fonction pour quitter les infos du personnage
	fmt.Println()
	fmt.Println("0 - Retour à la gestion du personnage")
	fmt.Println()
	textinfo := Input()
	switch textinfo {
	case "0":
		EffacerTerminal()
		fmt.Println()
		fmt.Println("------ Vous allez retourner au menu précédent ------")
		time.Sleep(1 * time.Second)
		p.Home() // Retour au menu si le joueur tape "Retour"
	default:
		EffacerTerminal()
		p.DisplayInfo()
		p.AccessInfo()
	}
}
