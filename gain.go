package main

import (
	"fmt"
	"time"
)

func GainExp(p *Personnage, m *Monstre) { // Gain Exp suivant les monstres vaincus
	switch m.nom {
	case "Gobelin":
		p.exp += 25
		fmt.Println("------ ", p.nom, " a gagné", p.exp, " pts d'XP ------")
	case "Licorne":
		p.exp += 100
		fmt.Println("------ ", p.nom, " a gagné", p.exp, " pts d'XP ------")
	case "Dragon":
		p.exp += 300
		fmt.Println("------ ", p.nom, " a gagné", p.exp, " pts d'XP ------")
	case "Alan":
		p.exp += 500
		fmt.Println("------ ", p.nom, " a gagné", p.exp, " pts d'XP ------")
	case "Lucas":
		p.exp += 1000
		fmt.Println("------ ", p.nom, " a gagné", p.exp, " pts d'XP ------")
	}
}

func GainMoney(p *Personnage, m *Monstre) { // Gain Money suivant les monstres vaincus
	switch m.nom {
	case "Gobelin":
		p.money += 10
		fmt.Println("------ ", p.nom, " a gagné 10ç ------")
	case "Licorne":
		p.money += 25
		fmt.Println("------ ", p.nom, " a gagné 25ç ------")
	case "Dragon":
		p.money += 75
		fmt.Println("------ ", p.nom, " a gagné 75ç ------")
	case "Alan":
		p.money += 100
		fmt.Println("------ ", p.nom, " a gagné 100ç ------")
	case "Lucas":
		p.money += 200
		fmt.Println("------ ", p.nom, " a gagné 200ç ------")
	}
}

func GainNiveau(p *Personnage) { // Gain Niveau suivant les monstres vaincus
	if p.exp > p.expmax || p.exp == p.expmax { // Condition si excédent de gain d'exp
		p.niveau += 1
		a := p.exp - p.expmax
		p.exp = 0
		p.exp += a
		p.expmax *= 1.5
		fmt.Println("------ ", p.nom, " est maintenant niveau : ", p.niveau, " (", p.exp, "/", p.expmax, ") ------")
		if p.classe == "Humain" { // Gain niveau donne bonus suivant classe
			fmt.Print("HPMAX : ", p.hpmax, " >>> ")
			p.hpmax += 15
			fmt.Println(p.hpmax, " (+15)")
			time.Sleep(1 * time.Second)
			fmt.Print("ManaMax : ", p.manamax, " >>> ")
			p.manamax += 15
			fmt.Println(p.manamax, " (+15)")
			time.Sleep(1 * time.Second)
			fmt.Print("Attaque : ", p.atk, " >>> ")
			p.atk += 3
			fmt.Println(p.atk, " (+3)")
			time.Sleep(1 * time.Second)
			fmt.Print("Puissance : ", p.puissance, " >>> ")
			p.puissance += 3
			fmt.Println(p.puissance, " (+3)")
		}
		if p.classe == "Elfe" { // Gain niveau donne bonus suivant classe
			fmt.Print("HPMAX : ", p.hpmax, " >>> ")
			p.hpmax += 10
			fmt.Println(p.hpmax, " (+10)")
			time.Sleep(1 * time.Second)
			fmt.Print("ManaMax : ", p.manamax, " >>> ")
			p.manamax += 20
			fmt.Println(p.manamax, " (+20)")
			time.Sleep(1 * time.Second)
			fmt.Print("Attaque : ", p.atk, " >>> ")
			p.atk += 2
			fmt.Println(p.atk, " (+2)")
			time.Sleep(1 * time.Second)
			fmt.Print("Puissance : ", p.puissance, " >>> ")
			p.puissance += 4
			fmt.Println(p.puissance, " (+4)")
		}
		if p.classe == "Nain" { // Gain niveau donne bonus suivant classe
			fmt.Print("HPMAX : ", p.hpmax, " >>> ")
			p.hpmax += 20
			fmt.Println(p.hpmax, " (+20)")
			time.Sleep(1 * time.Second)
			fmt.Print("ManaMax : ", p.manamax, " >>> ")
			p.manamax += 10
			fmt.Println(p.manamax, " (+10)")
			time.Sleep(1 * time.Second)
			fmt.Print("Attaque : ", p.atk, " >>> ")
			p.atk += 4
			fmt.Println(p.atk, " (+4)")
			time.Sleep(1 * time.Second)
			fmt.Print("Puissance : ", p.puissance, " >>> ")
			p.puissance += 2
			fmt.Println(p.puissance, " (+2)")
		}
		fmt.Println()
		if p.exp > p.expmax {
			GainNiveau(p)
		}
	}
}
