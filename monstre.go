package main

import (
	"fmt"
	"time"
)

type Monstre struct { // Création structure monstre
	nom        string
	hpmax      int
	hp         int
	atk        int
	initiative int
}

func (m *Monstre) InitMonstre(nom string, hpmax int, hp int, atk int, init int) { // Initialisation monstre
	m.nom = nom
	m.hpmax = hpmax
	m.hp = hp
	m.atk = atk
	m.initiative = init
}

func DeathMonstre(p *Personnage, m *Monstre) { // Condition mort du monstre
	if m.hp <= 0 {
		fmt.Println("👑👑👑👑👑👑 ", p.nom, "a gagné le combat 👑👑👑👑👑👑") // Message fin de game
		time.Sleep(1 * time.Second)
		GainExp(p, m) // Gain EXp
		time.Sleep(1 * time.Second)
		GainNiveau(p) // Gain Niveau suivant Exp
		time.Sleep(1 * time.Second)
		GainMoney(p, m) // Gain d'argent
		time.Sleep(3 * time.Second)
		fmt.Println()
		m.hp = m.hpmax // Réinitialisation pv monstre
		RetourMenu()
	}
}
