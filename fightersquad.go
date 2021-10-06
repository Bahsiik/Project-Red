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
			fmt.Println("C'est à ", p.nom, "d'agir")
			fmt.Println("Que va faire ", p.nom, " ?")
			fmt.Println("Rien (1) ou Coup de poing (2)")
			textfight := Input()
			switch textfight {
			case "1":
				fmt.Println()
				fmt.Println("C'est au ", m.nom, "d'agir")
				fmt.Println(m.nom, " attaque ", p.nom)
				p.hp -= m.atk
				fmt.Println(p.nom, " a maintenant ", p.hp, " Hp")
				fmt.Println()
				if p.hp <= 0 {
					p.Death()
					fmt.Println()
					RetourMenu()
				}
			case "2":
				fmt.Println(p.nom, " effectue un Coup de poing")
				m.hp -= 10
				fmt.Println(m.nom, " a maintenant ", m.hp, " Hp")
				if m.hp <= 0 {
					fmt.Println(p.nom, "a gagné le combat :))) uwu")
					fmt.Println()
					m.hp = m.hpmax
					RetourMenu()
				}
				fmt.Println()
				fmt.Println("C'est au ", m.nom, "d'agir")
				fmt.Println(m.nom, " attaque ", p.nom)
				p.hp -= m.atk
				fmt.Println(p.nom, " a maintenant ", p.hp, " Hp")
				fmt.Println()
				if p.hp <= 0 {
					p.Death()
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
