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
var Licorn Monstre
var Dragon Monstre

func GobelinInit(m *Monstre) {
	Gobelin.InitMonstre("Goblin d'entrainement", 40, 40, 5, 10)
}

func LicornInit(m *Monstre) {
	Licorn.InitMonstre("Licorne d'entrainement", 60, 60, 10, 20)
}

func DragonInit(m *Monstre) {
	Dragon.InitMonstre("Dragon d'entrainement", 200, 200, 20, 5)
}
