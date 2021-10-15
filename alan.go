package main

import (
	"fmt"
	"time"
)

var Alan Monstre

func AlanInit(m *Monstre) {
	Alan.InitMonstre("Alan", 500, 500, 40, 25)
}

func AttaqueAlan(p *Personnage, m *Monstre) {
	fmt.Println()
	fmt.Println(m.nom, " délaisse ", p.nom, " pour aller voir les marcoms...")
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

func AttaqueCritAlan(p *Personnage, m *Monstre) {
	fmt.Println()
	fmt.Println(m.nom, " force ", p.nom, " à  utiliser Go Playground...")
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

func AlanPattern(p *Personnage, m *Monstre, tour int) {
	fmt.Println()
	fmt.Println("C'est au tour du ", m.nom, ", bonne chance...")
	fmt.Println()
	if tour%3 == 0 {
		AttaqueCritAlan(p, m)
	} else {
		AttaqueAlan(p, m)
	}
}
