package main

import (
	"fmt"
	"time"
)

var Licorne Monstre // declaration licorne

func LicorneInit(m *Monstre) { // initalisation licorne
	Licorne.InitMonstre("Licorne", 100, 100, 10, 20)
}

func AttaqueLicorne(p *Personnage, m *Monstre) { //script attaque 1 licorne
	fmt.Println(m.nom, " donne un coup de sabot à ", p.nom, ". Cela inflige ", m.atk, " points de dégats")
	time.Sleep(1 * time.Second)
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

func AttaqueCritLicorne(p *Personnage, m *Monstre) { //script attaque 2 licorne
	fmt.Println(m.nom, " lance une Attaque arc-en-ciel et inflige ", m.atk*2, " points de dégats à ", p.nom)
	time.Sleep(1 * time.Second)
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

func LicornePattern(p *Personnage, m *Monstre, tour int) { // pattern licorne
	fmt.Println()
	fmt.Println("C'est au tour du ", m.nom)
	fmt.Println()
	time.Sleep(1 * time.Second)
	if tour%3 == 0 {
		AttaqueCritLicorne(p, m)
	} else {
		AttaqueLicorne(p, m)
	}
}
