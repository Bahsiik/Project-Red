package main

import (
	"fmt"
	"time"
)

var Dragon Monstre // Declaration monstre dragon

func DragonInit(m *Monstre) { // Initialisation Dragon
	Dragon.InitMonstre("Dragon", 200, 200, 20, 5)
}

func AttaqueDragon(p *Personnage, m *Monstre) { //Script attaque1 dragon
	fmt.Println(m.nom, " donne un coup de sabot à ", p.nom, "qui inflige ", m.atk, "points de dégats")
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

func AttaqueCritDragon1(p *Personnage, m *Monstre) { //Script attaque 2 dragon
	fmt.Println(m.nom, " lance un coup de queue et inflige ", m.atk*2, " points de dégats à ", p.nom)
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

func AttaqueCritDragon2(p *Personnage, m *Monstre) { //Script attaque 3 dragon
	fmt.Println(m.nom, " lance un jet enflammé et inflige ", m.atk*5, " points de dégats à ", p.nom)
	time.Sleep(1 * time.Second)
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

func DragonPattern(p *Personnage, m *Monstre, tour int) { // Pattern du dragon
	fmt.Println()
	fmt.Println("C'est au tour du ", m.nom)
	fmt.Println()
	time.Sleep(1 * time.Second)
	if tour%5 == 0 {
		AttaqueCritDragon2(p, m)
	} else if tour%2 == 0 {
		AttaqueCritDragon1(p, m)
	} else {
		AttaqueDragon(p, m)
	}
}
