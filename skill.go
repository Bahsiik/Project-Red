package main

import (
	"fmt"
)

func (p Personnage) DisplaySkill() { // Affichage des sorts possédés
	fmt.Print("--- Les sorts de ", p.nom, " sont ---  \n")
	if len(p.skill) == 0 {
		fmt.Println(" ", p.nom, "ne possède pas de sorts...")
	} else {
		for i := range p.skill {
			fmt.Print(" Sort ", i+1, " : ", p.skill[i], "\n")
		}
	}
	fmt.Println()
}

func (p *Personnage) AccessSkill() { // Commande d'accès aux sorts
	fmt.Println("Quel sort", p.nom, "veut utiliser ? ( Rien)")
	fmt.Println()
	textskill := Input()
	switch textskill {
	case "Rien":
		RetourMenu()
	default:
		fmt.Println(P1.nom, "ne sais pas quoi faire..")
		fmt.Println()
		P1.AccessInventory()
	}
}

func (p *Personnage) SpellBook(sort string, livre string) { // Initialisation livre de sorts
	if p.VerifSkill(sort) {
		p.skill = append(p.skill, sort)
		fmt.Println(p.nom, "a appris le sort", sort)
		p.RemoveInv(livre)
	} else {
		fmt.Println(p.nom, "connait déjà ce sort..")
	}
}

func (p *Personnage) VerifSpellBook(livre string) bool { // Vérification si le livre est présent (true=présent / false=absent)
	for i := range p.inventaire {
		if p.inventaire[i] == livre {
			return true
		}
	}
	return false
}

func (p *Personnage) VerifSkill(sort string) bool { // Vérification si le sort est présent dans les skill (true=présent / false=absent)
	for i := range p.skill {
		if p.skill[i] == sort {
			return true
		}
	}
	return false
}
