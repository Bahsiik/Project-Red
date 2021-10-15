package main

import (
	"fmt"
	"time"
)

func TrainingFight(p *Personnage, m *Monstre) { // Initialisation combat d'entrainement
	fmt.Println(p.nom, " engage le combat d'entrainement")
	fmt.Println()
	fmt.Println("FIGHT !!!")
	for tour := 1; ; tour++ { // Condition de fin de combat
		if p.hp > 0 || m.hp > 0 {
			time.Sleep(3 * time.Second)
			EffacerTerminal()
			fmt.Println("======== Tour ", tour, " ========") // Initialisation n° de tours
			fmt.Println(p.nom, " :", p.hp, "/", p.hpmax, "HP  |||", m.nom, " :", m.hp, "/", m.hpmax, "HP")
			fmt.Println()
			if p.initiative >= m.initiative {
				p.CharTurn(m)
				time.Sleep(1 * time.Second)
				GoblinPattern(p, m, tour)
				time.Sleep(1 * time.Second)
			} else {
				GoblinPattern(p, m, tour)
				time.Sleep(1 * time.Second)
				p.CharTurn(m)
				time.Sleep(1 * time.Second)
			}
		}
	}
}
