package main

import (
	"fmt"
)

func (p Personnage) DisplayEquipment() { // Fonction d'affichage de l'equipement du personnage
	fmt.Print("--- Equipement de ", p.nom, " ---  \n")
	fmt.Print("Tête : ", p.Equipement.tete, "\n")
	fmt.Print("Torse : ", p.Equipement.torse, "\n")
	fmt.Print("Pieds : ", p.Equipement.pieds, "\n")
	fmt.Println()
}

func (p *Personnage) AccessEquipment() { // Fonction pour modifier l'equipement du personnage
	fmt.Print("Que va faire ", p.nom, " ? (Rien)\n")
	textequip := Input()
	switch textequip {
	case "Rien":
		RetourMenu()
	default:
		fmt.Println(P1.nom, "ne sais pas quoi faire..")
		fmt.Println()
		P1.AccessEquipment()
	}
}

func (p *Personnage) AddEquipementTete(obj string) { // Fonction d'ajout d'un objet a l'equipement Tete
	p.Equipement.tete = obj
}

func (p *Personnage) AddEquipementTorse(obj string) { // Fonction d'ajout d'un objet a l'equipement Tete
	p.Equipement.torse = obj
}

func (p *Personnage) AddEquipementPieds(obj string) { // Fonction d'ajout d'un objet a l'equipement Tete
	p.Equipement.pieds = obj
}
