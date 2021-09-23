package main

import "fmt"

func (p Personnage) accessInventory() {
	fmt.Print("--- Inventaire de ", p.nom, " ---  \n")
	if len(p.inventaire) == 0 {
		fmt.Println("L'inventaire de", p.nom, "est vide...")
	} else {
		for i := range p.inventaire {
			fmt.Print(" Objet ", i+1, " : ", p.inventaire[i], "\n")
		}
	}
}
