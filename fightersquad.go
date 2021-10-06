package main

import "fmt"

func TrainingFight(p *Personnage, m *Monstre) { // Initialisation ocmbat d'entrainement
	fmt.Println(p.nom, " engage le combat d'entrainement")
	fmt.Println()
	fmt.Println("FIGHT !!!")
	for tour := 1; p.hp != 0 || m.hp != 0; tour++ { // Condition de fin de combat
		fmt.Println("======== Tour ", tour, " ========")
		textfight := Input()
		switch textfight {
		case "Rien":
			fmt.Println("C'est au tour de ", m.nom)
			fmt.Println()
			fmt.Println(m.nom, " attaque ", p.nom)
			p.hp -= m.atk
			fmt.Println(p.nom, " a maintenant ", p.hp, " Hp")
			fmt.Println()
		case "Coup de poing":
			fmt.Println(p.nom, " effectue un Coup de poing")
			m.hp -= 10
			fmt.Println(m.nom, " a maintenant ", m.hp, " Hp")
			fmt.Println()
		}
	}
}
