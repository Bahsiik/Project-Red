package main

import (
	"fmt"
)

func TrainingFight(p *Personnage, m *Monstre) { // Initialisation combat d'entrainement
	fmt.Println(p.nom, " engage le combat d'entrainement")
	fmt.Println()
	fmt.Println("FIGHT !!!")
	for tour := 1; ; tour++ { // Condition de fin de combat
		if p.hp > 0 || m.hp > 0 {
			fmt.Println("======== Tour ", tour, " ========") // Initialisation n° de tours
			fmt.Println(p.nom, " :", p.hp, "/", p.hpmax, "HP  |||", m.nom, " :", m.hp, "/", m.hpmax, "HP")
			if p.initiative >= m.initiative {
				p.CharTurn(m)
				GoblinPattern(p, m, tour)
			} else {
				p.CharTurn(m)
				GoblinPattern(p, m, tour)
			}
		}
	}
}
