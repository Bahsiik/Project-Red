package main

import (
	"fmt"
	"time"
)

var Licorne Monstre

func LicorneInit(m *Monstre) {
	Licorne.InitMonstre("Licorne", 60, 60, 10, 20)
}

func AttaqueLicorne(p *Personnage, m *Monstre) {
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

func AttaqueCritLicorne(p *Personnage, m *Monstre) {
	fmt.Println()
	fmt.Println(m.nom, " lance une Attaque arc-en-ciel et inflige ", m.atk*2, " de coup critique à ", p.nom)
	fmt.Println()
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

func LicornePattern(p *Personnage, m *Monstre, tour int) {
	fmt.Println()
	fmt.Println("C'est au tour du ", m.nom)
	fmt.Println()
	if tour%3 == 0 {
		AttaqueCritLicorne(p, m)
	} else {
		AttaqueLicorne(p, m)
	}
}
