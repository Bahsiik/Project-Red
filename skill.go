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

func (p *Personnage) SpellBook(sort string) { // Initialisation livre de sorts
	verif := 0
	for i := range p.skill {
		if p.skill[i] == sort {
			verif += 1
		}
	}
	if verif == 0 {
		p.skill = append(p.skill, sort)
		fmt.Println(p.nom, "a appris le sort Boule de Feu ")
	} else {
		fmt.Println(p.nom, "connait déjà ce sort..")
	}
}

func (p *Personnage) VerifSkill(sort string) bool { // Vérification unicité des sorts dans le livre
	for i := range p.skill {
		if p.skill[i] == sort {
			return false
		}
	}
	return true
}
