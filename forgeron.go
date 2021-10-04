package main

import (
	"fmt"
)

func (p Personnage) DisplayInvForgeron() { // Fonction d'affichage de l'inventaire du forgeron (Articles du magasin)
	fmt.Println("Bonjour, que souhaitez vous crafter ? (Numéro d'article / Rien)")
	if len(p.inventaire) == 0 {
		fmt.Println("Désolé, je n'ai rien a vous proposer...")
	} else {
		for i := range p.inventaire {
			fmt.Print(" Article ", i+1, " : ", p.inventaire[i], "\n")
		}
	}
	fmt.Println()
}
