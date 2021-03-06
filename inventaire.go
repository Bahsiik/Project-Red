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
	case "Potion de soin":
		p.TakeHealPot()
		p.DisplayInventory()
		p.AccessInventory()
	case "Potion de mana":
		p.TakeManaPot()
		p.DisplayInventory()
		p.AccessInventory()
	case "Potion de poison":
		p.PoisonPot()
		p.DisplayInventory()
		p.AccessInventory()
	case "Chapeau de l'aventurier":
		p.AddEquipementTete(textinv)
		fmt.Println()
		p.DisplayInventory()
		p.AccessInventory()
	case "Tunique de l'aventurier":
		p.AddEquipementTorse(textinv)
		fmt.Println()
		p.DisplayInventory()
		p.AccessInventory()
	case "Bottes de l'aventurier":
		p.AddEquipementPieds(textinv)
		fmt.Println()
		p.DisplayInventory()
		p.AccessInventory()
	case "Casque de Dieu-Roi":
		p.AddEquipementTete(textinv)
		fmt.Println("TESST")
		fmt.Println()
		p.DisplayInventory()
		p.AccessInventory()
	case "Armure de Dieu-Roi":
		p.AddEquipementTorse(textinv)
		fmt.Println()
		p.DisplayInventory()
		p.AccessInventory()
	case "Jambières de Dieu-Roi":
		p.AddEquipementPieds(textinv)
		fmt.Println()
		p.DisplayInventory()
		p.AccessInventory()
	case "Livre de sort : Iron Fist":
		p.UseSpellBook("Iron Fist")
	case "Livre de sort : Charge du Berserker":
		p.UseSpellBook("Charge du Berserker")
	case "Livre de sort : Boule de Feu":
		p.UseSpellBook("Boule de Feu")
	case "Livre de sort : Blizzard":
		p.UseSpellBook("Blizzard")
	case "Livre de sort : Décharge énergétique":
		p.UseSpellBook("Décharge énergétique")
	case "0": // Retour au menu
		EffacerTerminal()
		fmt.Println()
		fmt.Println("-------- Vous allez retourner au menu précédent --------")
		fmt.Println()
		time.Sleep(1 * time.Second)
		p.Home()
	default: // Choix d'objet invalide
		fmt.Println(p.nom, "ne sais pas quoi faire..")
		fmt.Println()
		time.Sleep(1 * time.Second)
		EffacerTerminal()
		p.DisplayInventory()
		p.AccessInventory()
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
			p.AccessInvMarchand()
		case "Non": // Retour au menu
			fmt.Println("Très bien, au revoir")
			fmt.Println()
			RetourMenu()
		}
	}
}

func (p *Personnage) AccessInvFight(m *Monstre) { // Utilisation de l'inventaire en cours de combat
	fmt.Println()
	fmt.Println("------ Quel objet", p.nom, "veut utiliser ? (Nom de l'objet) ------")
	fmt.Println()
	fmt.Println("0 - Retour à la gestion du personnage")
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
		fmt.Println(p.nom, "ne sais pas quoi faire..")
		fmt.Println()
		p.CharTurn(m)
	}
}
