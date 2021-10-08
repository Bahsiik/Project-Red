package main

import (
	"fmt"
)

func (p *Personnage) CharTurn(m *Monstre) {
	fmt.Println("Que va faire ", p.nom, " ?")
	fmt.Println("1 - Attaquer")
	fmt.Println("2 - Inventaire")
	textmenucharturn := Input()
	switch textmenucharturn {
	case "1":
		fmt.Println("1 - Attaque de base")
		fmt.Println("2 - Coup de poing")
		textattcharturn := Input()
		switch textattcharturn {
		case "1":
			AttaqueBasique(p, m)
		case "2":
			CoupPoing(p, m)
		}
	case "2":
		p.DisplayInventory()
		p.AccessInvFight(m)
	default: // Choix d'action invalide
		fmt.Println()
		fmt.Println("Désolé, cette commande est invalide, veuillez faire un autre choix.")
		fmt.Println()
		P1.CharTurn(m)
	}
}

func CoupPoing(p *Personnage, m *Monstre) {
	fmt.Println(p.nom, " effectue un Coup de poing")
	m.hp -= 10
	fmt.Println(m.nom, " a maintenant ", m.hp, "HP sur", m.hpmax, "HP.") // Affichage pv monstre fin tour
	if m.hp <= 0 {
		fmt.Println(p.nom, "a gagné le combat :))) uwu") // Message fin de game
		fmt.Println()
		m.hp = m.hpmax // Réinitialisation pv monstre
		RetourMenu()
	}
}

func AttaqueBasique(p *Personnage, m *Monstre) {
	fmt.Println(p.nom, " effectue une attaque basique")
	m.hp -= 5
	fmt.Println(m.nom, " a maintenant ", m.hp, "HP sur", m.hpmax, "HP.") // Affichage pv monstre fin tour
	if m.hp <= 0 {
		fmt.Println(p.nom, "a gagné le combat :))) uwu") // Message fin de game
		fmt.Println()
		m.hp = m.hpmax // Réinitialisation pv monstre
		RetourMenu()
	}
}

func BouleFeu(p *Personnage, m *Monstre) {
	fmt.Println(p.nom, " lance une boule de feu !!!!!!!!!!!!")
	m.hp -= 20
	fmt.Println(m.nom, " a maintenant ", m.hp, "HP sur", m.hpmax, "HP.") // Affichage pv monstre fin tour
	if m.hp <= 0 {
		fmt.Println(p.nom, "a gagné le combat :))) uwu") // Message fin de game
		fmt.Println()
		m.hp = m.hpmax // Réinitialisation pv monstre
		RetourMenu()
	}
}

func AttaqueGobelin(p *Personnage, m *Monstre) {
	fmt.Println()
	fmt.Println("C'est au ", m.nom, "d'agir") // Tour du Gobelin
	fmt.Println(m.nom, " attaque ", p.nom)
	p.hp -= m.atk
	fmt.Println(p.nom, " a maintenant ", p.hp, "HP sur", p.hpmax, "HP.") // Affichage pv perso fin tour
	fmt.Println()
	if p.hp <= 0 { // Condition pv à 0 du perso
		p.Death()
		fmt.Println()
		RetourMenu()
	}
}

func AttaqueCritGobelin(p *Personnage, m *Monstre) {
	fmt.Println()
	fmt.Println("C'est au ", m.nom, "d'agir") // Tour du Gobelin
	fmt.Println(m.nom, " inflige une attaque critique à ", p.nom)
	p.hp -= m.atk * 2
	fmt.Println(p.nom, " a maintenant ", p.hp, "HP sur", p.hpmax, "HP.") // Affichage pv perso fin tour
	fmt.Println()
	if p.hp <= 0 { // Condition pv à 0 du perso
		p.Death()
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
	case "Retour":
		p.CharTurn(m)
	}
}
