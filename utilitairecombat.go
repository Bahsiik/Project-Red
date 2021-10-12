package main

import (
	"fmt"
	"time"
)

func (p *Personnage) CharTurn(m *Monstre) {
	fmt.Println("C'est à ", p.nom, "d'agir")
	fmt.Println("Que va faire ", p.nom, " ?")
	fmt.Println("1 - Attaquer")
	fmt.Println("2 - Utiliser un sort")
	fmt.Println("3 - Inventaire")
	fmt.Println("4 - Fuir")
	fmt.Println()
	textmenucharturn := Input()
	switch textmenucharturn {
	case "1":
		AttaqueBasique(p, m)
	case "2":
		fmt.Println("0 - Retour")
		fmt.Println("1 - Coup de poing")
		verif := 0
		for i := range p.skill {
			if p.skill[i] == "Boule de Feu" {
				verif += 1
			}
		}
		if verif != 0 {
			fmt.Println("2 - Boule de Feu")
		}
		textattcharturn := Input()
		switch textattcharturn {
		case "1":
			CoupPoing(p, m)
		case "2":
			BouleFeu(p, m)
		case "0":
			P1.CharTurn(m)
		default: // Choix d'action invalide
			fmt.Println()
			fmt.Println("Désolé, cette commande est invalide, veuillez faire un autre choix.")
			fmt.Println()
			P1.CharTurn(m)
		}
	case "3":
		p.DisplayInventory()
		p.AccessInvFight(m)
	case "4":
		RetourMenu()
	default: // Choix d'action invalide
		fmt.Println()
		fmt.Println("Désolé, cette commande est invalide, veuillez faire un autre choix.")
		fmt.Println()
		P1.CharTurn(m)
	}
}

func CoupPoing(p *Personnage, m *Monstre) {
	if p.mana >= 5 {
		fmt.Println(p.nom, " effectue un Coup de poing")
		m.hp -= 8
		p.mana -= 5
		fmt.Println(m.nom, " a maintenant ", p.mana, "Mana sur", p.manamax, "Mana.")
		fmt.Println(m.nom, " a maintenant ", m.hp, "HP sur", m.hpmax, "HP.") // Affichage pv monstre fin tour
		fmt.Println()
		DeathMonstre(p, m)
	} else {
		fmt.Println(p.nom, " n'a pas assez de mana pour ce sort...")
		fmt.Println()
		P1.CharTurn(m)
	}
}

func AttaqueBasique(p *Personnage, m *Monstre) {
	fmt.Println(p.nom, " effectue une attaque basique")
	m.hp -= 5
	fmt.Println(m.nom, " a maintenant ", m.hp, "HP sur", m.hpmax, "HP.") // Affichage pv monstre fin tour
	DeathMonstre(p, m)
}

func BouleFeu(p *Personnage, m *Monstre) {
	if p.mana >= 10 {
		verif := 0
		for i := range p.skill {
			if p.skill[i] == "Boule de Feu" {
				verif += 1
			}
		}
		if verif == 0 {
			fmt.Println()
			fmt.Println(p.nom, " ne possède pas ce sort...")
			fmt.Println()
			P1.CharTurn(m)
		} else {
			fmt.Println(p.nom, " lance une boule de feu !!!!!!!!!!!!")
			m.hp -= 18
			p.mana -= 10
			fmt.Println(m.nom, " a maintenant ", p.mana, "Mana sur", p.manamax, "Mana.")
			fmt.Println(m.nom, " a maintenant ", m.hp, "HP sur", m.hpmax, "HP.") // Affichage pv monstre fin tour
			DeathMonstre(p, m)
		}
	} else {
		fmt.Println(p.nom, " n'a pas assez de mana pour ce sort...")
		fmt.Println()
		P1.CharTurn(m)
	}
}

func AttaqueGobelin(p *Personnage, m *Monstre) {
	fmt.Println(m.nom, " attaque ", p.nom)
	p.hp -= m.atk
	fmt.Println(p.nom, " a maintenant ", p.hp, "HP sur", p.hpmax, "HP.") // Affichage pv perso fin tour
	fmt.Println()
	if p.hp <= 0 { // Condition pv à 0 du perso
		p.Death()
		time.Sleep(1 * time.Second)
		fmt.Println()
		RetourMenu()
	}
}

func AttaqueCritGobelin(p *Personnage, m *Monstre) {
	fmt.Println(m.nom, " inflige une attaque critique à ", p.nom)
	p.hp -= m.atk * 2
	fmt.Println(p.nom, " a maintenant ", p.hp, "HP sur", p.hpmax, "HP.") // Affichage pv perso fin tour
	fmt.Println()
	if p.hp <= 0 { // Condition pv à 0 du perso
		p.Death()
		time.Sleep(1 * time.Second)
		fmt.Println()
		RetourMenu()
	}
}

func (p *Personnage) AccessInvFight(m *Monstre) {
	fmt.Println("Quel objet", p.nom, "veut utiliser ? (Nom de l'objet / Retour)")
	fmt.Println()
	textinvfight := Input()
	switch textinvfight {
	case "Potion":
		p.TakePot()
	case "Potion de poison":
		p.PoisonPotComb(m)
	case "Retour":
		p.CharTurn(m)
	default:
		fmt.Println(P1.nom, "ne sais pas quoi faire..")
		fmt.Println()
		p.CharTurn(m)
	}
}

func (p *Personnage) PoisonPotComb(m *Monstre) { // Fonction potion de poison
	var test2 int
	for i := range p.inventaire {
		if i < len(p.inventaire) {
			if p.inventaire[i] == "Potion de poison" { // Incrémentattion du compteur de potions par rapport à l'inventaire
				test2++
				fmt.Println(p.nom, " utilise une potion de poison sur ", m.nom)
				p.RemoveInv("Potion de poison")
				test := 0
				for i := 0; i < 3; i++ { // Initialisation des 3s d'effet de la potion
					if test == 0 { // Dégats/sec
						m.hp -= 10
						fmt.Println(m.nom, " a maintenant ", m.hp, "HP sur", m.hpmax, "HP.") // Affichage pv monstre fin tour
						if m.hp <= 0 {
							fmt.Println(p.nom, "a gagné le combat :))) uwu") // Message fin de game
							p.exp += 10
							fmt.Println()
							m.hp = m.hpmax // Réinitialisation pv monstre
							RetourMenu()
						}
						time.Sleep(1 * time.Second) // Timer dégats par secondes
					}
				}
			}
		}
	}
	if test2 == 0 { // Dans l'absence de potions dans l'inventaire de combat
		fmt.Println(p.nom, "n'as malheureusement pas cette potion...")
		fmt.Println()
		p.CharTurn(m)
	}
}
