package main

import (
	"fmt"
)

// Création structure Perso
type Personnage struct {
	nom        string
	classe     string
	niveau     int
	hpmax      int
	hp         int
	inventaire []string
	skill      []string
	money      int
}

// Paramètres à initialiser
func (p *Personnage) Init(nom string, classe string, niveau int, hpmax int, hp int, inventaire []string, skill []string, money int) {
	p.nom = nom
	p.classe = classe
	p.niveau = niveau
	p.hpmax = hpmax
	p.hp = hp
	p.inventaire = inventaire
	p.skill = skill
	p.money = money
}

// Création variable Perso 1 et Marchand
var P1 Personnage
var Marchand Personnage

// Fonction pour initialiser
func PersoInit(p *Personnage) {
	P1.Init("Byleth", "Roturier", 1, 100, 50, []string{"Epée", "Armure légère", "Potion", "Potion"}, []string{"Coup de poing"}, 100)
	Marchand.Init("Jeff Besos", "Marchand", 777, 777, 777, []string{"Potion", "Potion de poison", "Livre de sort: Boule de feu"}, []string{"Coup de poing"}, 999)
}

// Système de mort et de résurection
func (p *Personnage) Death() {
	if p.hp <= 0 {
		fmt.Println("Wasted ! Retente ta chance chacal..")
		p.hp = p.hpmax / 2
		fmt.Println("Grâce à John Cena,", p.nom, "a réssucité avec", p.hp, "HP")
	}
}
