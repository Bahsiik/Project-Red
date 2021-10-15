package main

import (
	"fmt"
	"time"
)

func TrainingFight(p *Personnage, m *Monstre, pattern func(p *Personnage, m *Monstre, tour int)) { // Initialisation combat d'entrainement
	fmt.Println()
	fmt.Println("------ ", p.nom, " engage le combat d'entrainement ------")
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
