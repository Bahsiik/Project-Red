package main

import (
	"fmt"
	"os"
)

func TrainingFight(p *Personnage, m *Monstre) { // Initialisation combat d'entrainement
	fmt.Println(p.nom, " engage le combat d'entrainement")
	fmt.Println()
	fmt.Println("FIGHT !!!")
	for tour := 1; ; tour++ { // Condition de fin de combat
		if p.hp > 0 || m.hp > 0 {
			fmt.Println("======== Tour ", tour, " ========") // Initialisation n° de tours
			fmt.Println("C'est à ", p.nom, "d'agir")
			fmt.Println("Que va faire ", p.nom, " ?")
			fmt.Println("Rien (1) ou Coup de poing (2)") // Choix d'actions du personnage
			fmt.Println()
			textfight := Input()
			switch textfight {
			case "1":
				AttaqueGobelin(p, m)
			case "2":
				CoupPoing(p, m)
				AttaqueGobelin(p, m)
			case "666": // Fin du monde
				fmt.Println(" Vous avez détruit le monde Mao-sama !!!!")
				fmt.Println()
				os.Exit(0)

			}
		}
	}
}
