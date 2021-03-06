package main

import (
	"fmt"
	"time"
)

func AttaqueBasique(p *Personnage, m *Monstre) { // Utilisation  AttaqueBasique
	fmt.Println()
	fmt.Println(p.nom, " effectue une attaque basique")
	m.hp -= p.atk
	time.Sleep(1 * time.Second)
	fmt.Println(m.nom, " a maintenant ", m.hp, "HP sur", m.hpmax, "HP.") // Affichage pv monstre fin tour
	fmt.Println()
	DeathMonstre(p, m)
}

func CoupPoing(p *Personnage, m *Monstre) { // Utilisation Coup de poing
	if p.mana >= 5 {
		fmt.Println()
		fmt.Println(p.nom, " effectue un Coup de poing")
		m.hp -= p.atk * 2
		p.mana -= 5
		time.Sleep(1 * time.Second)
		fmt.Println(p.nom, " a maintenant ", p.mana, "Mana sur", p.manamax, "Mana.")
		fmt.Println(m.nom, " a maintenant ", m.hp, "HP sur", m.hpmax, "HP.") // Affichage pv monstre fin tour
		fmt.Println()
		DeathMonstre(p, m)
	} else {
		fmt.Println(p.nom, " n'a pas assez de mana pour ce sort...")
		fmt.Println()
		p.CharTurn(m)
	}
}

func IronFist(p *Personnage, m *Monstre) { // Utilisation Coup de poing
	if p.mana >= 5 {
		fmt.Println()
		fmt.Println(p.nom, " effectue un Iron Fist !!")
		m.hp -= p.atk * 3
		p.mana -= 10
		time.Sleep(1 * time.Second)
		fmt.Println(p.nom, " a maintenant ", p.mana, "Mana sur", p.manamax, "Mana.")
		fmt.Println(m.nom, " a maintenant ", m.hp, "HP sur", m.hpmax, "HP.") // Affichage pv monstre fin tour
		fmt.Println()
		DeathMonstre(p, m)
	} else {
		fmt.Println(p.nom, " n'a pas assez de mana pour ce sort...")
		fmt.Println()
		p.CharTurn(m)
	}
}

func ChargeBerserker(p *Personnage, m *Monstre) { // Utilisation Coup de poing
	if p.mana >= 5 {
		fmt.Println()
		fmt.Println(p.nom, " effectue une charge du Berserker !!")
		m.hp -= p.atk * 5
		p.mana -= 30
		time.Sleep(1 * time.Second)
		fmt.Println(p.nom, " a maintenant ", p.mana, "Mana sur", p.manamax, "Mana.")
		fmt.Println(m.nom, " a maintenant ", m.hp, "HP sur", m.hpmax, "HP.") // Affichage pv monstre fin tour
		fmt.Println()
		DeathMonstre(p, m)
	} else {
		fmt.Println(p.nom, " n'a pas assez de mana pour ce sort...")
		fmt.Println()
		p.CharTurn(m)
	}
}

func BouleFeu(p *Personnage, m *Monstre) { // Utilisation Boule de Feu
	if p.mana >= 10 {
		verif := 0
		for i := range p.skill {
			if p.skill[i] == "Boule de Feu" {
				verif += 1
			}
		}
		if verif == 0 {
			fmt.Println()
			fmt.Println(p.nom, " ne poss??de pas ce sort...") // Cas absence de ce sort en notre possession
			fmt.Println()
			p.CharTurn(m) // Retour choix perso
		} else {
			fmt.Println()
			fmt.Println(p.nom, " lance une boule de feu !!!!!!!!!!!!")
			m.hp -= p.puissance * 2
			p.mana -= 15
			time.Sleep(1 * time.Second)
			fmt.Println(p.nom, " a maintenant ", p.mana, "Mana sur", p.manamax, "Mana.") // Affichage mana perso fin tour
			fmt.Println(m.nom, " a maintenant ", m.hp, "HP sur", m.hpmax, "HP.")         // Affichage pv monstre fin tour
			fmt.Println()
			DeathMonstre(p, m)
		}
	} else {
		fmt.Println()
		fmt.Println(p.nom, " n'a pas assez de mana pour ce sort...")
		fmt.Println()
		p.CharTurn(m)
	}
}

func Blizzard(p *Personnage, m *Monstre) { // Utilisation Blizzard
	if p.mana >= 25 {
		verif := 0
		for i := range p.skill {
			if p.skill[i] == "Blizzard" {
				verif += 1
			}
		}
		if verif == 0 {
			fmt.Println()
			fmt.Println(p.nom, " ne poss??de pas ce sort...") // Cas absence de ce sort en notre possession
			fmt.Println()
			p.CharTurn(m) // Retour choix perso
		} else {
			fmt.Println()
			fmt.Println(p.nom, " lance un blizzard !!!!!!!!!!!!")
			m.hp -= p.puissance * 4
			p.mana -= 30
			time.Sleep(1 * time.Second)
			fmt.Println(p.nom, " a maintenant ", p.mana, "Mana sur", p.manamax, "Mana.") // Affichage mana perso fin tour
			fmt.Println(m.nom, " a maintenant ", m.hp, "HP sur", m.hpmax, "HP.")         // Affichage pv monstre fin tour
			fmt.Println()
			DeathMonstre(p, m)
		}
	} else {
		fmt.Println()
		fmt.Println(p.nom, " n'a pas assez de mana pour ce sort...")
		fmt.Println()
		p.CharTurn(m)
	}
}

func DechargeEnergetique(p *Personnage, m *Monstre) { // Utilisation D??charge ??nerg??tique
	if p.mana >= 100 {
		verif := 0
		for i := range p.skill {
			if p.skill[i] == "D??charge ??nerg??tique" {
				verif += 1
			}
		}
		if verif == 0 {
			fmt.Println()
			fmt.Println(p.nom, " ne poss??de pas ce sort...") // Cas absence de ce sort en notre possession
			fmt.Println()
			p.CharTurn(m) // Retour choix perso
		} else {
			fmt.Println()
			fmt.Println(p.nom, " lance une d??charge ??nerg??tique !!!!!!!!!!!!")
			m.hp -= p.puissance * 6
			p.mana -= 100
			time.Sleep(1 * time.Second)
			fmt.Println(p.nom, " a maintenant ", p.mana, "Mana sur", p.manamax, "Mana.") // Affichage mana perso fin tour
			fmt.Println(m.nom, " a maintenant ", m.hp, "HP sur", m.hpmax, "HP.")         // Affichage pv monstre fin tour
			fmt.Println()
			DeathMonstre(p, m)
		}
	} else {
		fmt.Println()
		fmt.Println(p.nom, " n'a pas assez de mana pour ce sort...")
		fmt.Println()
		p.CharTurn(m)
	}
}
