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
	skill      []string
	money      int
	Equipement struct {
		tete  string
		torse string
		pieds string
	}
}

func (p *Personnage) Init(nom string, classe string, niveau int, hpmax int, hp int, inventaire []string, skill []string, money int, tete string, torse string, pieds string) { // Paramètres à initialiser
	p.nom = nom
	p.classe = classe
	p.niveau = niveau
	p.hpmax = hpmax
	p.hp = hp
	p.inventaire = inventaire
	p.skill = skill
	p.money = money
	p.Equipement.tete = tete
	p.Equipement.torse = torse
	p.Equipement.pieds = pieds
}

// Création variable Perso 1 et Marchand
var P1 Personnage
var Marchand Personnage

func PersoInit(p *Personnage) { // Fonction pour initialiser les personnages
	P1.Init("Byleth", "Roturier", 1, 100, 50, []string{"Epée", "Armure légère", "Potion", "Potion"}, []string{"Coup de poing"}, 100, "chapeau de caca", "torse de caca", "pieds de caca")
	Marchand.Init("Jeff Besos", "Marchand", 777, 777, 777, []string{"Potion", "Potion de poison", "Livre de sort: Boule de feu", "Fourrure de Loup", "Peau de troll", "Cuir de Sanglier", "Plume de Corbeau"}, []string{"Coup de poing"}, 999, "Chapeau Gucci", "Veste Luis Vuitton", "Chaussures Geox")
}

func (p *Personnage) Death() { // Système de mort et de résurection
	if p.hp <= 0 {
		fmt.Println("Wasted ! Retente ta chance chacal..")
		p.hp = p.hpmax / 2
		fmt.Println("Grâce à John Cena,", p.nom, "a réssucité avec", p.hp, "HP")
	}
}
