package main

import "fmt"

func (p Personnage) DisplayInfo() {
	fmt.Println("--- Informations du personnage ---")
	fmt.Println("Nom du personnage --> ", p.nom)
	fmt.Println("Classe du personnage --> ", p.classe)
	fmt.Println("Niveau du personnage --> ", p.niveau)
	fmt.Println("HP maximum du personnage --> ", p.hpmax)
	fmt.Println("HP actuels du personnage --> ", p.hp)
}

func RemoveInv(inv []string, obj string, i int) []string {
	var newinv []string
	for i := range inv {
		if inv[i] == obj {
			newinv = append(inv[:i], inv[i+1:]...)
		}
	}
	return newinv
}
