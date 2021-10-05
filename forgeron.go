package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func (p *Personnage) AccessInvForgeron() { //Fonction craft d'objet
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	textforgeron, _ := reader.ReadString('\n')
	textforgeron = strings.Replace(textforgeron, "\r\n", "", -1)
	switch textforgeron {
	case "Rien": // Retour au menu (aucun objet choisi)
		fmt.Println("Très bien, passez une bonne journée.")
		fmt.Println()
		RetourMenu()
	case "1": // Craft Chapeau de l'aventurier
		var test int
		for i := range p.inventaire { 
			if i < len(p.inventaire) {
				if p.inventaire[i] == "Plume de Corbeau : 1ç"  && p.inventaire[i] == "Cuir de Sanglier : 3ç" { // Incrémentattion du compteur de composants par rapport à l'inventaire
					test++	
					if (P1.money - 5) >= 0 {
						fmt.Println("Argent restant : ", P1.money, " Cacas d'or")
						fmt.Println()
						ContinueForegronInv("Chapeau de l'aventurier")
					} else {
						p.Pauvre()
					} 
					}
				}
			}
			if test == 0 { // Dans l'absence des composants dans l'inventaire
		fmt.Println(p.nom, "n'as malheureusement pas les composants nécessaire...")
	}
	case "2" :// Craft Tunique de l'aventurier
	var test int
		for i := range p.inventaire { 
			if i < len(p.inventaire) {
				if p.inventaire[i] == "Fourrure de Loup : 4ç"  && p.inventaire[i] == "Peau de troll : 7ç" { // Incrémentattion du compteur de composants par rapport à l'inventaire
					for c := range p.inventaire {
						if c != i && c < len(p.inventaire[i]) {	
							if p.inventaire[c] == "Fourrure de Loup : 4ç"{					
								test++	
									if (P1.money - 5) >= 0 {
										fmt.Println("Argent restant : ", P1.money, " Cacas d'or")
										fmt.Println()
										ContinueForegronInv("Tunique de l'aventurier")
									} else {
										p.Pauvre()
									}
								}
							}
						} 
					}
				}
			}
			if test == 0 { // Dans l'absence des composants dans l'inventaire
		fmt.Println(p.nom, "n'as malheureusement pas les composants nécessaire...")
	}
	case "3" : // Craft Bottes de l'aventurier
		var test int
		for i := range p.inventaire { 
			if i < len(p.inventaire) {
				if p.inventaire[i] == "Cuir de Sanglier : 3ç"  && p.inventaire[i] == "Fourrure de Loup : 4ç" { // Incrémentattion du compteur de composants par rapport à l'inventaire
					test++	
					if (P1.money - 5) >= 0 {
						fmt.Println("Argent restant : ", P1.money, " Cacas d'or")
						fmt.Println()
						ContinueForgeronInv("Bottes de l'aventurier")
					} else {
						p.Pauvre()
					} 
					}
				}
			}
			if test == 0 { // Dans l'absence des composants dans l'inventaire
		fmt.Println(p.nom, "n'as malheureusement pas les composants nécessaire...")
	}

}

	case "2" :
	for i := range p.inventaire {
		if i < len(p.inventaire) {
			if p.inventaire[i] == "Potion" { // Incrémentattion du compteur de potions par rapport à l'inventaire
				test++
}
