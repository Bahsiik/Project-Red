package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
var Marchand Personnage

func PersoInit(p *Personnage) {
	P1.Init("Byleth", "Roturier", 1, 100, 50, []string{"Epée", "Armure légère", "Potion", "Potion"})
	Marchand.Init("Jeff Besos", "Marchand", 777, 777, 777, []string{"Potion", "Potion de poison"})
}

func (p Personnage) DisplayInfo() {
	fmt.Println("--- Informations du personnage ---")
	fmt.Println("Nom --> ", p.nom)
	fmt.Println("Classe --> ", p.classe)
	fmt.Println("Niveau --> ", p.niveau)
	fmt.Println("HP maximum --> ", p.hpmax)
	fmt.Println("HP actuels --> ", p.hp)
	fmt.Println()
	AccessInfo()
}

func AccessInfo() {
	fmt.Println("(Retour au menu --> Tapez retour)")
	fmt.Println()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	textainfo, _ := reader.ReadString('\n')
	// convert CRLF to LF
	textainfo = strings.Replace(textainfo, "\r\n", "", -1)
	switch textainfo {
	case "Retour":
		RetourMenu()
	}
}

func (p *Personnage) Death() {
	if p.hp <= 0 {
		fmt.Println("Wasted ! Retente ta chance chacal..")
		p.hp = p.hpmax / 2
		fmt.Println("Grâce à John Cena,", p.nom, "a réssucité avec", p.hp, "HP")
	}
}
