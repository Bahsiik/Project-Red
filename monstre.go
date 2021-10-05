package main

type Monstre struct { // Création structure monstre
	nom   string
	hpmax int
	hp    int
	atk   int
}

func (m *Monstre) InitMonstre(nom string, hpmax int, hp int, atk int) {
	m.nom = nom
	m.hpmax = hpmax
	m.hp = hp
	m.atk = atk
}

var Goblin Monstre

func MonstreInit(m *Monstre) {
	Goblin.InitMonstre("Goblin d'entrainement", 40, 40, 5)
}
