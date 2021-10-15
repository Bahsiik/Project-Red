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
	fmt.Println("HP actuels --> ", p.hp, " sur ", p.hpmax)
	fmt.Println("Mana actuel --> ", p.mana, " sur ", p.manamax)
	fmt.Println("Exp actuel --> ", p.exp, " sur ", p.expmax)
	fmt.Println("Argent actuel --> ", p.money)
	fmt.Println()
	AccessInfo()
}

func AccessInfo() { // Fonction pour quitter les infos du personnage
	fmt.Println("0 - Retour à la gestion du personnage")
	fmt.Println()
	textinfo := Input()
	switch textinfo {
	case "0":
		EffacerTerminal()
		fmt.Println("------ Vous allez retourner au menu précédent ------")
		time.Sleep(1 * time.Second)
		Home() // Retour au menu si le joueur tape "Retour"
	default:
		EffacerTerminal()
		P1.DisplayInfo()
		AccessInfo()
	}
}
