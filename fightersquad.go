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
				fmt.Println()
				fmt.Println("C'est au ", m.nom, "d'agir") // Tour du Gobelin
				fmt.Println(m.nom, " attaque ", p.nom)
				p.hp -= 5
				fmt.Println(p.nom, " a maintenant ", p.hp, " Hp") // Affichage pv perso fin tour
				fmt.Println()
				if p.hp <= 0 { // Condition pv à 0 du perso
					p.Death()
					fmt.Println()
					RetourMenu()
				}
			case "2":
				fmt.Println(p.nom, " effectue un Coup de poing")
				m.hp -= 10
				fmt.Println(m.nom, " a maintenant ", m.hp, " Hp") // Affichage pv monstre fin tour
				fmt.Println()
				fmt.Println("C'est au ", m.nom, "d'agir")
				fmt.Println(m.nom, " attaque ", p.nom)
				p.hp -= m.atk
				fmt.Println(p.nom, " a maintenant ", p.hp, " Hp")
				fmt.Println()
				if m.hp <= 0 {
					fmt.Println(p.nom, "a gagné le combat :))) uwu") // Message fin de game
					fmt.Println()
					m.hp = m.hpmax // Réinitialisation pv monstre
					RetourMenu()
				} else if p.hp <= 0 {
					p.Death()
					fmt.Println()
					RetourMenu()
				}
			case "666": // Fin du monde
				fmt.Println(" Vous avez détruit le monde Mao-sama !!!!")
				fmt.Println()
				os.Exit(0)

			}
		}
	}
}
