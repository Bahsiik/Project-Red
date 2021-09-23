package main

import "fmt"

type Personnage struct {
	nom        string
	classe     string
	niveau     int
	hpmax      int
	hp         int
	inventaire []string
}

func (p *Personnage) Init(nom string, classe string, niveau int, hpmax int, hp int, inventaire []string) {
	p.nom = nom
	p.classe = classe
	p.niveau = niveau
	p.hpmax = hpmax
	p.hp = hp
	p.inventaire = inventaire
}

var P1 Personnage

func PersoInit(p *Personnage) {
	P1.Init("Byleth", "Roturier", 1, 100, 50, []string{"Epée", "Armure légère", "Potion", "Potion"})
}

func (p Personnage) DisplayInfo() {
	fmt.Println("--- Informations du personnage ---")
	fmt.Println("Nom --> ", p.nom)
	fmt.Println("Classe --> ", p.classe)
	fmt.Println("Niveau --> ", p.niveau)
	fmt.Println("HP maximum --> ", p.hpmax)
	fmt.Println("HP actuels --> ", p.hp)
	fmt.Println()
}
