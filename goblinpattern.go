package main

func GoblinPattern(p *Personnage, m *Monstre, tour int) {
	if tour%3 == 0 {
		AttaqueCritGobelin(p, m)
	} else {
		AttaqueGobelin(p, m)
	}
}
