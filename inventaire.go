package main

import (
	"fmt"
)

func (p Personnage) DisplayInventory() { // Fonction affichage de l'inventaire
	fmt.Print("--- Inventaire de ", p.nom, " ---  \n")
	if len(p.inventaire) == 0 {
		fmt.Println("L'inventaire de", p.nom, "est vide...")
	} else {
		for i := range p.inventaire {
			fmt.Print(" Objet ", i+1, " : ", p.inventaire[i], "\n")
		}
	}
	fmt.Println()
}

func (p *Personnage) AccessInventory() { // Fonction d'utilisation d'objet dans l'inventaire
	fmt.Println("Quel objet", p.nom, "veut utiliser ? (Nom de l'objet / Rien)")
	fmt.Println()
	textinv := Input()
	switch textinv {
	case "Potion": // Utilisation de la potion
		P1.TakePot()
		P1.DisplayInventory()
		P1.AccessInventory()
	case "Potion de poison":
		P1.PoisonPot()
		P1.DisplayInventory()
		P1.AccessInventory()
	case "Chapeau de l'aventurier":
		P1.AddEquipementTete(textinv)
		fmt.Println()
		P1.DisplayInventory()
		P1.AccessInventory()
	case "Tunique de l'aventurier":
		P1.AddEquipementTorse(textinv)
		fmt.Println()
		P1.DisplayInventory()
		P1.AccessInventory()
	case "Bottes de l'aventurier":
		P1.AddEquipementPieds(textinv)
		fmt.Println()
		P1.DisplayInventory()
		P1.AccessInventory()
	case "Livre de sort : Boule de Feu":
		if p.VerifSpellBook("Livre de sort : Boule de Feu") {
			P1.SpellBook("Boule de Feu", "Livre de sort : Boule de Feu")
			fmt.Println()
			P1.DisplayInventory()
			P1.AccessInventory()
		} else {
			fmt.Println(P1.nom, "ne sais pas quoi faire..")
			fmt.Println()
			P1.AccessInventory()
		}
	case "Rien": // Retour au menu
		RetourMenu()
	default: // Choix d'objet invalide
		fmt.Println(P1.nom, "ne sais pas quoi faire..")
		fmt.Println()
		P1.AccessInventory()
	}
}

func (p *Personnage) AddInventory(obj string) { // Fonction d'ajout d'un objet a l'inventaire
	p.inventaire = append(p.inventaire, obj)
}

func (p *Personnage) RemoveInv(obj string) { // Fonction de retrait d'un objet a l'inventaire
	for i := range p.inventaire {
		if i < len(p.inventaire) {
			if p.inventaire[i] == obj {
				p.inventaire = append(p.inventaire[:i], p.inventaire[i+1:]...)
				break
			}
		}
	}
}
