package main

import (
	"fmt"
)

type Personnage struct { // Création structure Perso
	nom        string
	classe     string
	niveau     int
	hpmax      int
	hp         int
	inventaire []string
	tailleinv  int
	skill      []string
	money      int
	Equipement struct {
		tete  string
		torse string
		pieds string
	}
	initiative int
}

func (p *Personnage) Init(nom string, classe string, niveau int, hpmax int, hp int, inventaire []string, tailleinv int, skill []string, money int, init int) { // Paramètres à initialiser
	p.nom = nom
	p.classe = classe
	p.niveau = niveau
	p.hpmax = hpmax
	p.hp = hp
	p.inventaire = inventaire
	p.tailleinv = tailleinv
	p.skill = skill
	p.money = money
	p.initiative = init
}

// Création variable Perso 1 et Marchand
var P1 Personnage

func PersoInit(p *Personnage) { // Fonction pour initialiser les personnages
	P1.Init("Byleth", "Humain", 1, 100, 50, []string{"Potion", "Potion"}, 10, []string{"Coup de poing"}, 500, 15)
}

func (p *Personnage) Death() { // Système de mort et de résurection
	if p.hp <= 0 {
		fmt.Println("Wasted ! Retente ta chance chacal..")
		p.hp = p.hpmax / 2
		fmt.Println("Grâce à John Cena,", p.nom, "a réssucité avec", p.hp, "HP")
	}
}

func (p *Personnage) Pauvre() {
	if p.money <= 0 {
		fmt.Println("Vous n'avez plus d'argent...")
	}
}
