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
	var test int
	for i := range p.inventaire {
		if p.inventaire[i] == obj {
			test++
			if p.Equipement.tete != "" { // Ajout à l'inventaire de l'item précédent
				P1.AddInventory(p.Equipement.tete)
			}
			p.Equipement.tete = obj // Changement d'équipement de tête
			p.RemoveInv(obj)        // Retirer l'item de l'inventaire
			p.hpmax += 10           // Bonus d'équipement
			fmt.Println("Les Hp Max de, ", p.nom, " passe à ", p.hpmax, " en équipant", obj)
			fmt.Println()
		}
	}
	if test == 0 {
		fmt.Println("Vous ne possédez pas cet équipement..")
	}
}

func (p *Personnage) AddEquipementTorse(obj string) { // Fonction d'ajout d'un objet a l'equipement Tete
	var test int
	for i := range p.inventaire {
		if p.inventaire[i] == obj {
			test++
			if p.Equipement.torse != "" {
				P1.AddInventory(p.Equipement.torse)
			}
			p.Equipement.torse = obj
			p.RemoveInv("obj")
			p.hpmax += 25
			fmt.Println("Les Hp Max de, ", p.nom, " passe à ", p.hpmax, " en équipant ", obj)
		}
	}
	if test == 0 {
		fmt.Println("Vous ne possédez pas cet équipement..")
	}
}

func (p *Personnage) AddEquipementPieds(obj string) { // Fonction d'ajout d'un objet a l'equipement Tete
	var test int
	for i := range p.inventaire {
		if p.inventaire[i] == obj {
			test++
			if p.Equipement.pieds != "" {
				P1.AddInventory(p.Equipement.pieds)
			}
			p.Equipement.pieds = obj
			p.RemoveInv(obj)
			p.hpmax += 15
			fmt.Println("Les Hp Max de, ", p.nom, " passe à ", p.hpmax, " en équipant ", obj)
		}
	}
	if test == 0 {
		fmt.Println("Vous ne possédez pas cet équipement..")
	}
}
