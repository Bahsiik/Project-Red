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
	mana       int
	manamax    int
}

func (p *Personnage) Init(nom string, classe string, niveau int, hpmax int, hp int, inventaire []string, tailleinv int, skill []string, money int, init int, mana int, manamax int) { // Paramètres à initialiser
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
	p.mana = mana
	p.manamax = manamax
}

// Création variable globale Perso 1
var P1 Personnage

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

func NameCreation() string {

	fmt.Println("Veuillez choisir votre nom :")
	textnom := Input()
	if !IsLetter(textnom) {
		fmt.Println("Ce nom n'est pas valide, choisissez en un autre")
		return NameCreation()
	}
	return Capitalize(textnom)
}

func ClassChoice() string {
	var class string
	fmt.Println("Veuillez choisir votre classe :")
	fmt.Println("1 - Humain")
	fmt.Println("2 - Elfe")
	fmt.Println("3 - Nain")
	textclass := Input()
	switch textclass {
	case "1":
		class = "Humain"
	case "2":
		class = "Elfe"
	case "3":
		class = "Nain"
	default:
		fmt.Println()
		fmt.Println("Choix de classe invalide, refaites votre choix")
		ClassChoice()
	}
	return class
}

func (p *Personnage) CharCreation() {
	var hpmax int
	var init int
	var manamax int
	nom := NameCreation()
	classe := ClassChoice()
	niveau := 1
	if classe == "Humain" {
		hpmax = 100
		init = 15
		manamax = 50
	}
	if classe == "Elfe" {
		hpmax = 80
		init = 20
		manamax = 100
	}
	if classe == "Nain" {
		hpmax = 120
		init = 5
		manamax = 25
	}
	hp := hpmax / 2
	inventaire := []string{"Potion", "Potion"}
	tailleinv := 10
	skill := []string{"Coup de poing"}
	money := 100
	mana := manamax
	P1.Init(nom, classe, niveau, hpmax, hp, inventaire, tailleinv, skill, money, init, mana, manamax)
}
