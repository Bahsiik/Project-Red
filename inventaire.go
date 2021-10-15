package main

import (
	"fmt"
	"time"
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
	fmt.Println("----- Quel objet", p.nom, "veut utiliser ? (Nom de l'objet) -----")
	fmt.Println()
	fmt.Println("0 - Retour à la gestion du personnage")
	fmt.Println()
	textinv := Input()
	switch textinv {
	case "Potion de soin": // Utilisation de la potion
		P1.TakeHealPot()
		P1.DisplayInventory()
		P1.AccessInventory()
	case "Potion de mana":
		P1.TakeManaPot()
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
	case "Livre de sort : Iron Fist":
		if p.VerifSpellBook("Livre de sort : Iron Fist") {
			P1.SpellBook("Iron Fist", "Livre de sort : Iron Fist")
			fmt.Println()
			P1.DisplayInventory()
			P1.AccessInventory()
		} else {
			fmt.Println(P1.nom, "ne sais pas quoi faire..")
			fmt.Println()
			P1.AccessInventory()
		}
	case "Livre de sort : Charge du Berserker":
		if p.VerifSpellBook("Livre de sort : Charge du Berserker") {
			P1.SpellBook("Charge du Berserker", "Livre de sort : Charge du Berserker")
			fmt.Println()
			P1.DisplayInventory()
			P1.AccessInventory()
		} else {
			fmt.Println(P1.nom, "ne sais pas quoi faire..")
			fmt.Println()
			P1.AccessInventory()
		}
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
	case "Livre de sort : Blizzard":
		if p.VerifSpellBook("Livre de sort : Blizzard") {
			P1.SpellBook("Blizzard", "Livre de sort : Blizzard")
			fmt.Println()
			P1.DisplayInventory()
			P1.AccessInventory()
		} else {
			fmt.Println(P1.nom, "ne sais pas quoi faire..")
			fmt.Println()
			P1.AccessInventory()
		}
	case "Livre de sort : Décharge énergétique":
		if p.VerifSpellBook("Livre de sort : Décharge énergétique") {
			P1.SpellBook("Décharge énergétique", "Livre de sort : Décharge énergétique")
			fmt.Println()
			P1.DisplayInventory()
			P1.AccessInventory()
		} else {
			fmt.Println(P1.nom, "ne sais pas quoi faire..")
			fmt.Println()
			P1.AccessInventory()
		}
	case "0": // Retour au menu
		EffacerTerminal()
		fmt.Println()
		fmt.Println("-------- Vous allez retourner au menu précédent --------")
		fmt.Println()
		time.Sleep(1 * time.Second)
		Home()
	default: // Choix d'objet invalide
		fmt.Println(P1.nom, "ne sais pas quoi faire..")
		fmt.Println()
		time.Sleep(1 * time.Second)
		EffacerTerminal()
		P1.DisplayInventory()
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
		fmt.Println(p.nom, "a acheté une nouvelle sacoche Gucci. Il peut désormais transporté 10 objets supplémentaires !")
		fmt.Println()
	} else {
		fmt.Println("Je suis désolé mais vous ne pouvez pas porter une sacoche de plus..")
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("Besoin d'autres choses messire ? (Oui/Non)")
	fmt.Println()
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

func (p *Personnage) AccessInvFight(m *Monstre) { // Utilisation de l'inventaire en cours de combat
	fmt.Println()
	fmt.Println("------ Quel objet", p.nom, "veut utiliser ? (Nom de l'objet / Retour) ------")
	fmt.Println()
	textinvfight := Input()
	switch textinvfight {
	case "Potion de soin":
		p.TakeHealPot()
	case "Potion de mana":
		p.TakeManaPot()
	case "Potion de poison":
		p.PoisonPotComb(m)
	case "Retour":
		p.CharTurn(m)
	default: // Condition par défaut
		fmt.Println(P1.nom, "ne sais pas quoi faire..")
		fmt.Println()
		p.CharTurn(m)
	}
}
