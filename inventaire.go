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
		Home()
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

func (p *Personnage) VerifTailleInv() bool { // Fonction pour vérifier la taille de l'inventaire (true=taille correcte / false=taille max atteinte)
	if len(p.inventaire) < p.tailleinv {
		return true
	} else {
		return false
	}
}

func (p *Personnage) UpgradeInventorySlot() { // Fonction pour augmenter la taille de l'inventaire
	if p.tailleinv < 40 {
		p.tailleinv += 10
		fmt.Println(p.nom, "a acheté une nouvelle sacoche. Il peut désormais transporté 10 objets supplémentaires !")
	} else {
		fmt.Println("Je suis désolé mais vous ne pouvez pas porter une sacoche de plus..")
	}
	fmt.Println("Besoin d'autres choses messire ? (Oui/Non)")
	for {
		textmarchand2 := Input()
		switch textmarchand2 {
		case "Oui": // Continuation des achats
			Marchand.DisplayInvMarchand2()
			fmt.Println()
			P1.AccessInvMarchand()
		case "Non": // Retour au menu
			fmt.Println("Très bien, au revoir")
			fmt.Println()
			RetourMenu()
		}
	}
}
