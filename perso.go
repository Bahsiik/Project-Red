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
var Forgeron Personnage

func PersoInit(p *Personnage) { // Fonction pour initialiser les personnages
	P1.Init("Byleth", "Roturier", 1, 100, 50, []string{"Epée", "Armure légère", "Potion", "Potion"}, []string{"Coup de poing"}, 100, "chapeau de caca", "torse de caca", "pieds de caca")
	Marchand.Init("Jeff Besos", "Marchand", 777, 777, 777, []string{"Potion : 3ç", "Potion de poison : 6ç", "Livre de sort: Boule de feu : 25ç", "Fourrure de Loup : 4ç", "Peau de troll : 7ç", "Cuir de Sanglier : 3ç", "Plume de Corbeau : 1ç"}, []string{"Coup de poing"}, 999, "Chapeau Gucci", "Veste Luis Vuitton", "Chaussures Geox")
	Forgeron.Init("Mickey", "Forgeron", 666, 666, 666, []string{"Chapeau de l'aventurier : 5ç", "Tunique de l'aventurier : 5ç", "Bottes de l'aventurier : 5ç"}, []string{"Coup de poing"}, 666, "", "", "")
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
