package main

import (
	"fmt"
	"time"
)

var Lucas Monstre

func LucasInit(m *Monstre) {
	Lucas.InitMonstre("Lucas", 1000, 1000, 1, 15)
}

func AttaqueLucas(p *Personnage, m *Monstre) {
	fmt.Println()
	fmt.Println(m.nom, " demande à ", p.nom, "de faire un exercice Ytrack")
	fmt.Println()
	p.hp -= m.atk
	fmt.Println(p.nom, " a maintenant ", p.hp, "HP sur", p.hpmax, "HP.") // Affichage pv perso fin tour
	fmt.Println()
	fmt.Println("3")
	time.Sleep(1 * time.Second)
	fmt.Println("2")
	time.Sleep(1 * time.Second)
	fmt.Println("1")
	time.Sleep(1 * time.Second)
	fmt.Println(m.nom, " vient d'envoyer quelqu'un a l'administration.. Son attaque a doublée !")
	m.atk *= 2
	if p.hp <= 0 { // Condition pv à 0 du perso
		p.Death()
		time.Sleep(1 * time.Second)
		fmt.Println()
		RetourMenu()
	}
}

func LucasPattern(p *Personnage, m *Monstre, tour int) {
	fmt.Println()
	fmt.Println("C'est au tour du ", m.nom, ", bonne chance...")
	fmt.Println()
	AttaqueLucas(p, m)
}
