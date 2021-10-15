package main

import (
	"fmt"
	"time"
)

func AttaqueLicorn(p *Personnage, m *Monstre) {
	fmt.Println()
	fmt.Println(m.nom, " donne un coup de sabot à ", p.nom)
	fmt.Println()
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

func AttaqueCritLicorn(p *Personnage, m *Monstre) {
	m.atk *= 2
	fmt.Println()
	fmt.Println(m.nom, " lance une Attaque arc-en-ciel et inflige ", m.atk, " de coup critique à ", p.nom)
	fmt.Println()
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

func LicornPattern(p *Personnage, m *Monstre, tour int) {
	fmt.Println()
	fmt.Println("C'est au tour du ", m.nom)
	fmt.Println()
	if tour%3 == 0 {
		AttaqueCritLicorn(p, m)
	} else {
		AttaqueLicorn(p, m)
	}
}
