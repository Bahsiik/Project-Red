package main

import (
	"fmt"
	"time"
)

var Lucas Monstre //déclaration lucas

func LucasInit(m *Monstre) { //initialisation lucas
	Lucas.InitMonstre("Lucas", 1000, 1000, 1, 15)
}

func AttaqueLucas(p *Personnage, m *Monstre) { // script attaque Lucas
	fmt.Println(m.nom, " demande à ", p.nom, "de faire un exercice Ytrack. Cela inflige ", m.atk, " points de dégats")
	time.Sleep(1 * time.Second)
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

func LucasPattern(p *Personnage, m *Monstre, tour int) { // pattern lucas
	fmt.Println()
	fmt.Println("C'est au tour du ", m.nom, ", bonne chance...")
	fmt.Println()
	time.Sleep(1 * time.Second)
	AttaqueLucas(p, m)
}
