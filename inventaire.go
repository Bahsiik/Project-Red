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
		textinv = ""
		P1.AccessInventory()
	case "Chapeau de l'aventurier": // Craft du Chapeau de laventurier
		P1.AddInventory(p.Equipement.tete)              // Ajout à l'inventaire de l'item précédent
		P1.AddEquipementTete("Chapeau de l'aventurier") // Changement d'équipement de tête
		p.RemoveInv("Chapeau de l'aventurier")          // Retirer l'item de l'inventaire
		p.hpmax += 10                                   // Bonus d'équipement
		fmt.Println("Les Hp Max de, ", p.nom, " passe à ", p.hpmax, " en équipant Chapeau de l'aventurier")
		fmt.Println()
		P1.AccessInventory()
	case "Tunique de l'aventurier":
		P1.AddInventory(p.Equipement.torse)
		P1.AddEquipementTorse("Tunique de l'aventurier")
		p.RemoveInv("Tunique de l'aventurier")
		p.hpmax += 25
		fmt.Println("Les Hp Max de, ", p.nom, " passe à ", p.hpmax, " en équipant Tunique de l'aventurier")
		fmt.Println()
		P1.AccessInventory()
	case "Bottes de l'aventurier":
		P1.AddInventory(p.Equipement.pieds)
		P1.AddEquipementPieds("Bottes de l'aventurier")
		p.RemoveInv("Bottes de l'aventurier")
		p.hpmax += 15
		fmt.Println("Les Hp Max de, ", p.nom, " passe à ", p.hpmax, " en équipant Bottes de l'aventurier")
		fmt.Println()
		P1.AccessInventory()
	case "Livre de sort : Boule de Feu":
		P1.SpellBook("Boule de Feu")
		fmt.Println()
		P1.AccessInventory()
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
