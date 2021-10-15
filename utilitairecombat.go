package main

import (
	"fmt"
	"time"
)

func (p *Personnage) CharTurn(m *Monstre) {
	fmt.Println()
	fmt.Println("C'est à ", p.nom, "d'agir")
	fmt.Println("Que va faire ", p.nom, " ?")
	fmt.Println()
	fmt.Println("1 - Attaquer")
	fmt.Println("2 - Utiliser un sort")
	fmt.Println("3 - Inventaire")
	fmt.Println()
	fmt.Println("0 - Abandonner")
	fmt.Println()
	textmenucharturn := Input()
	switch textmenucharturn {
	case "1":
		AttaqueBasique(p, m)
		fmt.Println()
	case "2":
		fmt.Println()
		fmt.Println("1 - Coup de poing")
		fmt.Println()
		if p.VerifSkill("Boule de Feu") {
			fmt.Println("2 - Boule de Feu")
			fmt.Println()
		}
		fmt.Println("0 - Retour")
		fmt.Println()
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
	case "0":
		RetourMenu()
	default: // Choix d'action invalide
		fmt.Println()
		fmt.Println("Désolé, cette commande est invalide, veuillez faire un autre choix.")
		fmt.Println()
		P1.CharTurn(m)
	}
}

func (p *Personnage) AccessInvFight(m *Monstre) {
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
				fmt.Println()
				p.RemoveInv("Potion de poison")
				test := 0
				for i := 0; i < 3; i++ { // Initialisation des 3s d'effet de la potion
					if test == 0 { // Dégats/sec
						m.hp -= 10
						fmt.Println()
						fmt.Println(m.nom, " a maintenant ", m.hp, "HP sur", m.hpmax, "HP.") // Affichage pv monstre fin tour
						fmt.Println()
						if m.hp <= 0 {
							fmt.Println()
							fmt.Println("👑👑👑👑👑👑 ", p.nom, "a gagné le combat :))) uwu 👑👑👑👑👑👑") // Message fin de game
							fmt.Println()
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
		fmt.Println()
		fmt.Println(p.nom, "n'as malheureusement pas cette potion...")
		fmt.Println()
		p.CharTurn(m)
	}
}
