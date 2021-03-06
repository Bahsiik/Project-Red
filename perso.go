package main

import (
	"fmt"
	"time"
)

type Personnage struct { // Création structure Perso
	nom        string
	classe     string
	niveau     int
	hpmax      int
	hp         int
	atk        int
	puissance  int
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
	exp        float32
	expmax     float32
}

func (p *Personnage) Init(nom string, classe string, niveau int, hpmax int, hp int, atk int, puissance int, inventaire []string, tailleinv int, skill []string, money int, init int, mana int, manamax int, exp float32, expmax float32) { // Paramètres à initialiser
	p.nom = nom
	p.classe = classe
	p.niveau = niveau
	p.hpmax = hpmax
	p.hp = hp
	p.atk = atk
	p.puissance = puissance
	p.inventaire = inventaire
	p.tailleinv = tailleinv
	p.skill = skill
	p.money = money
	p.initiative = init
	p.mana = mana
	p.manamax = manamax
	p.exp = exp
	p.expmax = expmax
}

// Création variable globale Perso 1
var P1 Personnage

func (p *Personnage) Death() { // Système de mort et de résurection
	if p.hp <= 0 {
		fmt.Println("Wasted ! Retente ta chance chacal..")
		time.Sleep(1 * time.Second)
		p.hp = p.hpmax / 2
		fmt.Println("Grâce à John Cena,", p.nom, "a réssucité avec", p.hp, "HP")
		time.Sleep(3 * time.Second)
	}
}

func (p *Personnage) Pauvre() { // Système de pauvreté
	fmt.Println("Vous n'avez pas assez d'argent...")
	fmt.Println()
	time.Sleep(2 * time.Second)
	p.Achats()
}

func NameCreation() string { // Système de création de nom
	fmt.Println("***** Veuillez choisir votre nom : *****")
	fmt.Println()
	textnom := Input()
	if !IsLetter(textnom) {
		fmt.Println()
		fmt.Println("Ce nom n'est pas valide, choisissez en un autre")
		fmt.Println()
		return NameCreation()
	}
	EffacerTerminal()
	return Capitalize(textnom)
}

func (p *Personnage) ClassChoice() { // Système de choix de classe
	fmt.Println()
	fmt.Println("***** Veuillez choisir votre classe : *****")
	fmt.Println()
	fmt.Println("1 - Humain")
	fmt.Println("2 - Elfe")
	fmt.Println("3 - Nain")
	fmt.Println()
	textclass := Input()
	switch textclass {
	case "1":
		p.classe = "Humain"
	case "2":
		p.classe = "Elfe"
	case "3":
		p.classe = "Nain"
	default:
		fmt.Println()
		fmt.Println("Choix de classe invalide, refaites votre choix")
		fmt.Println()
		p.ClassChoice()
	}
	EffacerTerminal()
}

func (p *Personnage) CharCreation() { // Fonction de création de perso selon classe
	p.nom = NameCreation()
	p.ClassChoice()
	switch p.classe {
	case "Humain":
		p.hpmax = 100
		p.initiative = 15
		p.manamax = 50
		p.atk = 4
		p.puissance = 4
	case "Elfe":
		p.hpmax = 80
		p.initiative = 20
		p.manamax = 75
		p.atk = 2
		p.puissance = 6
	case "Nain":
		p.hpmax = 120
		p.initiative = 5
		p.manamax = 25
		p.atk = 6
		p.puissance = 2
	}
	p.niveau = 1
	p.hp = p.hpmax / 2
	p.inventaire = []string{"Potion de soin", "Potion de soin", "Potion de mana"}
	p.tailleinv = 10
	p.skill = []string{"Coup de poing"}
	p.money = 20
	p.mana = p.manamax
	p.exp = 0
	p.expmax = 100
}
