package main

import (
	"fmt"
)

func (p Personnage) DisplayInfo() { // Affichage des informations du personnages
	fmt.Println("--- Informations du personnage ---")
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
	fmt.Println("(Retour au menu --> Tapez Retour)")
	fmt.Println()
	textinfo := Input()
	switch textinfo {
	case "Retour":
		RetourMenu() // Retour au menu si le joueur tape "Retour"
	default:
		AccessInfo()
	}
}
