package main

import (
	"fmt"
	"os"
)

func TrainingFight(p *Personnage, m *Monstre) { // Initialisation ocmbat d'entrainement
	fmt.Println(p.nom, " engage le combat d'entrainement")
	fmt.Println()
	fmt.Println("FIGHT !!!")
	for tour := 1; ; tour++ { // Condition de fin de combat
		if p.hp > 0 || m.hp > 0 {
			fmt.Println("======== Tour ", tour, " ========")
			textfight := Input()
			switch textfight {
			case "Rien":
				fmt.Println("C'est au tour de ", m.nom)
				fmt.Println()
				fmt.Println(m.nom, " attaque ", p.nom)
				p.hp -= 25
				fmt.Println(p.nom, " a maintenant ", p.hp, " Hp")
				fmt.Println()
				if p.hp <= 0 {
					p.Death()
					fmt.Println()
					RetourMenu()
				}
			case "Coup de poing":
				fmt.Println(p.nom, " effectue un Coup de poing")
				m.hp -= 10
				fmt.Println(m.nom, " a maintenant ", m.hp, " Hp")
				fmt.Println()
				if m.hp <= 0 {
					fmt.Println(p.nom, "a gagné le combat :))) uwu")
					fmt.Println()
					RetourMenu()
				}
			case "666":
				fmt.Println(" Vous avez détruit le monde Mao-sama !!!!")
				fmt.Println()
				os.Exit(0)

			}
		}
	}
}
