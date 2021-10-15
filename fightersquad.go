package main

import (
	"fmt"
	"time"
)

func TrainingFight(p *Personnage, m *Monstre, pattern func(p *Personnage, m *Monstre, tour int)) { // Initialisation combat
	fmt.Println()
	fmt.Println("------ ", p.nom, " engage le combat  ------")
	fmt.Println()
	fmt.Println("FIGHT !!!")
	fmt.Println()
	for tour := 1; ; tour++ { // Condition de fin de combat
		if p.hp > 0 || m.hp > 0 {
			time.Sleep(3 * time.Second)
			EffacerTerminal()
			fmt.Println()
			fmt.Println("======== Tour ", tour, " ========") // Initialisation n° de tours
			fmt.Println()
			fmt.Println(p.nom, " :", p.hp, "/", p.hpmax, "HP  |||", m.nom, " :", m.hp, "/", m.hpmax, "HP")
			fmt.Println()
			if p.initiative >= m.initiative { // Condition perso commence le fight
				p.CharTurn(m)
				time.Sleep(1 * time.Second)
				pattern(p, m, tour) // Suivant le monstre choisi, initie les données de combat
				time.Sleep(1 * time.Second)
			} else { // Condition monstre commence le fight
				pattern(p, m, tour)
				time.Sleep(1 * time.Second)
				p.CharTurn(m)
				time.Sleep(1 * time.Second)
			}
		}
	}
}

func (p *Personnage) CharTurn(m *Monstre) { // Initialisation tour personnage
	fmt.Println()
	fmt.Println("C'est à ", p.nom, "d'agir")
	fmt.Println("Que va faire ", p.nom, " ?")
	fmt.Println()
	fmt.Println("1 - Attaquer")
	fmt.Println("2 - Utiliser un sort")
	fmt.Println("3 - Inventaire")
	fmt.Println()
	fmt.Println("0 - Abandonner")
	fmt.Println()
	textmenucharturn := Input()
	switch textmenucharturn { // Choix perso a effectué
	case "1":
		AttaqueBasique(p, m)
		fmt.Println()
	case "2":
		fmt.Println()
		fmt.Println("1 - Coup de poing")
		if p.VerifSkill("Boule de Feu") {
			fmt.Println("2 - Boule de Feu")
		}
		if p.VerifSkill("Iron Fist") {
			fmt.Println("3 - Iron Fist")
		}
		if p.VerifSkill("Blizzard") {
			fmt.Println("4 - Blizzard")
		}
		if p.VerifSkill("Charge du Berserker") {
			fmt.Println("5 - Charge du Berserker")
		}
		if p.VerifSkill("Décharge énergétique") {
			fmt.Println("6 - Décharge énergétique")
		}
		fmt.Println()
		fmt.Println("0 - Retour")
		fmt.Println()
		textattcharturn := Input()
		switch textattcharturn {
		case "1":
			CoupPoing(p, m)
		case "2":
			BouleFeu(p, m)
		case "3":
			IronFist(p, m)
		case "4":
			Blizzard(p, m)
		case "5":
			ChargeBerserker(p, m)
		case "6":
			DechargeEnergetique(p, m)
		case "0":
			P1.CharTurn(m)
		default: // Choix d'action invalide
			fmt.Println()
			fmt.Println("Désolé, cette commande est invalide, veuillez faire un autre choix.")
			fmt.Println()
			P1.CharTurn(m)
		}
	case "3":
		p.DisplayInventory()
		p.AccessInvFight(m)
	case "0":
		RetourMenu()
	default: // Choix d'action invalide
		fmt.Println()
		fmt.Println("Désolé, cette commande est invalide, veuillez faire un autre choix.")
		fmt.Println()
		P1.CharTurn(m)
	}
}
