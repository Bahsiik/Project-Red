package main

type Monstre struct { // Création structure monstre
	nom        string
	hpmax      int
	hp         int
	atk        int
	initiative int
}

func (m *Monstre) InitMonstre(nom string, hpmax int, hp int, atk int, init int) {
	m.nom = nom
	m.hpmax = hpmax
	m.hp = hp
	m.atk = atk
	m.initiative = init
}

var Gobelin Monstre

func GobelinInit(m *Monstre) {
	Gobelin.InitMonstre("Goblin d'entrainement", 40, 40, 5, 10)
}
