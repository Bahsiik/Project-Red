package main

import (
	"fmt"
	"time"
)

var Gobelin Monstre // déclaration gobelin

func GobelinInit(m *Monstre) { // initialisation gobelin
	Gobelin.InitMonstre("Gobelin", 50, 50, 5, 10)
}

func AttaqueGobelin(p *Personnage, m *Monstre) { // script attaque 1 gobelin
	fmt.Println(m.nom, " attaque ", p.nom, ". Cela inflige ", m.atk, " points de dégats")
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

func AttaqueCritGobelin(p *Personnage, m *Monstre) { // script attaque 2 gobelin
	fmt.Println(m.nom, " inflige ", m.atk*2, " de dégats critique à ", p.nom)
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

func GobelinPattern(p *Personnage, m *Monstre, tour int) { // pattern gobelin
	fmt.Println()
	fmt.Println("C'est au tour du ", m.nom)
	fmt.Println()
	time.Sleep(1 * time.Second)
	if tour%4 == 0 {
		AttaqueCritGobelin(p, m)
	} else {
		AttaqueGobelin(p, m)
	}
}
