package main

import (
	"fmt"
	"time"
)

var Dragon Monstre

func DragonInit(m *Monstre) {
	Dragon.InitMonstre("Dragon d'entrainement", 200, 200, 20, 5)
}

func AttaqueDragon(p *Personnage, m *Monstre) {
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

func AttaqueCritDragon1(p *Personnage, m *Monstre) {
	fmt.Println()
	fmt.Println(m.nom, " lance un coup de queu et inflige ", m.atk*2, " de coup critique à ", p.nom)
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

func AttaqueCritDragon2(p *Personnage, m *Monstre) {
	fmt.Println()
	fmt.Println(m.nom, " lance un jet enflammé et inflige ", m.atk*5, " de coup critique à ", p.nom)
	fmt.Println()
	p.hp -= m.atk * 5
	fmt.Println(p.nom, " a maintenant ", p.hp, "HP sur", p.hpmax, "HP.") // Affichage pv perso fin tour
	fmt.Println()
	if p.hp <= 0 { // Condition pv à 0 du perso
		p.Death()
		time.Sleep(1 * time.Second)
		fmt.Println()
		RetourMenu()
	}
}

func DragonPattern(p *Personnage, m *Monstre, tour int) {
	fmt.Println()
	fmt.Println("C'est au tour du ", m.nom)
	fmt.Println()
	if tour%5 == 0 {
		AttaqueCritDragon2(p, m)
	} else if tour%2 == 0 {
		AttaqueCritDragon1(p, m)
	} else {
		AttaqueDragon(p, m)
	}
}
