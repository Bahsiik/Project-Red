package main

import (
	"fmt"
	"time"
)

var Gobelin Monstre

func GobelinInit(m *Monstre) {
	Gobelin.InitMonstre("Goblin d'entrainement", 40, 40, 5, 10)
}

func AttaqueGobelin(p *Personnage, m *Monstre) {
	fmt.Println()
	fmt.Println(m.nom, " attaque ", p.nom)
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

func AttaqueCritGobelin(p *Personnage, m *Monstre) {
	fmt.Println()
	fmt.Println(m.nom, " inflige ", m.atk*2, " de dégats critique à ", p.nom)
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

func GoblinPattern(p *Personnage, m *Monstre, tour int) {
	fmt.Println()
	fmt.Println()
	fmt.Println("C'est au tour du ", m.nom)
	fmt.Println()
	if tour%4 == 0 {
		AttaqueCritGobelin(p, m)
	} else {
		AttaqueGobelin(p, m)
	}
}
