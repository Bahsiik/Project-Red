package main

import "fmt"

func CoupPoing(p *Personnage, m *Monstre) {
	fmt.Println(p.nom, " effectue un Coup de poing")
	m.hp -= 10
	fmt.Println(m.nom, " a maintenant ", m.hp, " Hp") // Affichage pv monstre fin tour
	if m.hp <= 0 {
		fmt.Println(p.nom, "a gagné le combat :))) uwu") // Message fin de game
		fmt.Println()
		m.hp = m.hpmax // Réinitialisation pv monstre
		RetourMenu()
	}
}

func AttaqueBasique(p *Personnage, m *Monstre) {
	fmt.Println(p.nom, " effectue un Coup de poing")
	m.hp -= 5
	fmt.Println(m.nom, " a maintenant ", m.hp, " Hp") // Affichage pv monstre fin tour
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
	fmt.Println(m.nom, " a maintenant ", m.hp, " Hp") // Affichage pv monstre fin tour
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
	fmt.Println(p.nom, " a maintenant ", p.hp, " Hp") // Affichage pv perso fin tour
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
	fmt.Println(p.nom, " a maintenant ", p.hp, " Hp") // Affichage pv perso fin tour
	fmt.Println()
	if p.hp <= 0 { // Condition pv à 0 du perso
		p.Death()
		fmt.Println()
		RetourMenu()
	}
}
