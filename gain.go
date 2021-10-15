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
			p.GainStats(15, 15, 3, 3)
		}
		if p.classe == "Elfe" { // Gain niveau donne bonus suivant classe
			p.GainStats(10, 20, 2, 4)
		}
		if p.classe == "Nain" { // Gain niveau donne bonus suivant classe
			p.GainStats(20, 10, 4, 2)
		}
		fmt.Println()
		if p.exp > p.expmax {
			GainNiveau(p)
		}
	}
}

func (p *Personnage) GainStats(ghp int, gmana int, gatk int, gpuissance int) {
	fmt.Print("HPMAX : ", p.hpmax, " >>> ")
	p.hpmax += ghp
	fmt.Println(p.hpmax, " (+", ghp, ")")
	time.Sleep(1 * time.Second)
	fmt.Print("ManaMax : ", p.manamax, " >>> ")
	p.manamax += gmana
	fmt.Println(p.manamax, " (+", gmana, ")")
	time.Sleep(1 * time.Second)
	fmt.Print("Attaque : ", p.atk, " >>> ")
	p.atk += gatk
	fmt.Println(p.atk, " (+", gatk, ")")
	time.Sleep(1 * time.Second)
	fmt.Print("Puissance : ", p.puissance, " >>> ")
	p.puissance += gpuissance
	fmt.Println(p.puissance, " (+", gpuissance, ")")
}
