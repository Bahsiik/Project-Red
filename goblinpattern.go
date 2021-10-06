package main

import "fmt"

func GoblinPattern(p *Personnage, m *Monstre, tour int) {
	fmt.Println("C'est au tour du ", m.nom)
	fmt.Println()
	if tour%3 == 0 {
		AttaqueCritGobelin(p, m)
	} else {
		AttaqueGobelin(p, m)
	}
}
