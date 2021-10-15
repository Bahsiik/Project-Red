package main

import "fmt"

func CoupPoing(p *Personnage, m *Monstre) { // Initialisation fonction Coup de poing
	if p.mana >= 5 {
		fmt.Println()
		fmt.Println(p.nom, " effectue un Coup de poing")
		m.hp -= p.atk * 2
		p.mana -= 5
		fmt.Println(p.nom, " a maintenant ", p.mana, "Mana sur", p.manamax, "Mana.")
		fmt.Println(m.nom, " a maintenant ", m.hp, "HP sur", m.hpmax, "HP.") // Affichage pv monstre fin tour
		fmt.Println()
		DeathMonstre(p, m)
	} else {
		fmt.Println(p.nom, " n'a pas assez de mana pour ce sort...")
		fmt.Println()
		P1.CharTurn(m)
	}
}

func AttaqueBasique(p *Personnage, m *Monstre) { // Initialisation fonction  AttaqueBasique
	fmt.Println()
	fmt.Println(p.nom, " effectue une attaque basique")
	m.hp -= p.atk
	fmt.Println(m.nom, " a maintenant ", m.hp, "HP sur", m.hpmax, "HP.") // Affichage pv monstre fin tour
	fmt.Println()
	DeathMonstre(p, m)
}

func BouleFeu(p *Personnage, m *Monstre) { // Initialisation fonction Boule de Feu
	if p.mana >= 10 {
		verif := 0
		for i := range p.skill {
			if p.skill[i] == "Boule de Feu" {
				verif += 1
			}
		}
		if verif == 0 {
			fmt.Println()
			fmt.Println(p.nom, " ne possède pas ce sort...") // Cas absence de ce sort en notre possession
			fmt.Println()
			P1.CharTurn(m) // Retour choix perso
		} else {
			fmt.Println()
			fmt.Println(p.nom, " lance une boule de feu !!!!!!!!!!!!")
			m.hp -= p.puissance * 2
			p.mana -= 10
			fmt.Println(p.nom, " a maintenant ", p.mana, "Mana sur", p.manamax, "Mana.") // Affichage mana perso fin tour
			fmt.Println(m.nom, " a maintenant ", m.hp, "HP sur", m.hpmax, "HP.")         // Affichage pv monstre fin tour
			fmt.Println()
			DeathMonstre(p, m)
		}
	} else {
		fmt.Println()
		fmt.Println(p.nom, " n'a pas assez de mana pour ce sort...")
		fmt.Println()
		P1.CharTurn(m)
	}
}
