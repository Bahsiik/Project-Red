package main

import (
	"fmt"
	"time"
)

var Alan Monstre // Déclaration Alan

func AlanInit(m *Monstre) { // Initialisation Alan
	Alan.InitMonstre("Alan", 1000, 1000, 40, 25)
}

func AttaqueAlan(p *Personnage, m *Monstre) { // Script Attaque 1 Alan
	fmt.Println(m.nom, " délaisse ", p.nom, " pour aller voir les marcoms... Cela inflige ", m.atk, " points de dégats")
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

func AttaqueCritAlan(p *Personnage, m *Monstre) { // Script Attaque 2 Alan
	fmt.Println(m.nom, " force ", p.nom, " à  utiliser Go Playground... Cela inflige ", m.atk*3, " points de dégats")
	time.Sleep(1 * time.Second)
	p.hp -= m.atk * 3
	fmt.Println(p.nom, " a maintenant ", p.hp, "HP sur", p.hpmax, "HP.") // Affichage pv perso fin tour
	fmt.Println()
	if p.hp <= 0 { // Condition pv à 0 du perso
		p.Death()
		time.Sleep(1 * time.Second)
		fmt.Println()
		RetourMenu()
	}
}

func AlanPattern(p *Personnage, m *Monstre, tour int) { // Pattern de Alan
	fmt.Println()
	fmt.Println("C'est au tour du ", m.nom, ", bonne chance...")
	fmt.Println()
	time.Sleep(1 * time.Second)
	if tour%3 == 0 { // Init condition ocup critique
		AttaqueCritAlan(p, m)
	} else {
		AttaqueAlan(p, m)
	}
}
