package main

import (
	"fmt"
)

var Forgeron Personnage

func ForgeronInit(p *Personnage) {
	Forgeron.Init("Mickey", "Forgeron", 666, 666, 666, []string{"Chapeau de l'aventurier : 5ç (1 Plume de Corbeau / 1 Cuir de Sanglier)", "Tunique de l'aventurier : 5ç (2 Fourrure de Loup / 1 Peau de Troll)", "Bottes de l'aventurier : 5ç (1 Fourrure de Loup / 1 Cuir de Sanglier)"}, []string{"Coup de poing"}, 666, "", "", "")
}

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
	textforgeron := Input()
	switch textforgeron {
	case "Rien": // Retour au menu (aucun objet choisi)
		fmt.Println("Très bien, passez une bonne journée.")
		fmt.Println()
		RetourMenu()
	case "1": // Craft Chapeau de l'aventurier
		p.CraftItem("Plume de Corbeau", "Cuir de Sanglier", 1, 1, 5, "Chapeau de l'aventurier")
	case "2":
		p.CraftItem("Fourrure de Loup", "Peau de Troll", 2, 1, 5, "Tunique de l'aventurier")
	case "3":
		p.CraftItem("Fourrure de Loup", "Cuir de Sanglier", 1, 1, 5, "Bottes de l'aventurier")
	default:
		fmt.Print("désolé mais je ne comprend pas \n")
	}
}

func (p Personnage) ContinueForgeronInv(choix string) { // Fonction d'ajout de l'objet choisi dans l'inventaire + choix de continuer les achats ou non
	fmt.Println("Et voilà monsieur, c'est prêt !")
	fmt.Print("Il reste ", p.money, " ç à ", p.nom, "\n")
	fmt.Println("Besoin d'autres choses messire ? (Oui/Non)")
	textforgeron2 := Input()
	P1.AddInventory(choix) // Ajout de l'objet
	switch textforgeron2 {
	case "Oui": // Continuation des achats
		fmt.Println("Que voulez vous fabriquer d'autre ?")
		fmt.Println()
		P1.AccessInvMarchand()
	case "Non": // Retour au menu
		fmt.Println("Très bien, au revoir")
		fmt.Println()
		RetourMenu()
	}
}

func (p *Personnage) CraftItem(elem1 string, elem2 string, nbr1 int, nbr2 int, prix int, equip string) { // Fonction de fabrication d'objet
	var test1 int // Variable pour vérifier la présence du composant 1
	var test2 int // Variable pour vérifier la présence du composant 2
	for i := range p.inventaire {
		if i < len(p.inventaire) {
			if p.inventaire[i] == elem1 {
				test1++
			} else if p.inventaire[i] == elem2 {
				test2++
			}
		}
	}
	if test1 >= nbr1 && test2 >= nbr2 && p.money >= prix { // Si tout les éléments requis sont la, on retire l'argent et les composants puis on ajoute l'objet fabriqué
		p.money -= prix
		for i := 0; i < nbr1; i++ {
			p.RemoveInv(elem1)
		}
		for i := 0; i < nbr2; i++ {
			p.RemoveInv(elem2)
		}
		p.ContinueForgeronInv(equip)
	} else if test1 < nbr1 && test2 < nbr2 { // Cas où il manque des composants
		fmt.Print("Je suis désolé mais tu n'as pas les composants nécessaires... \n")
	} else if p.money < prix { // Cas où il manque de l'argent
		fmt.Print("Je suis désolé mais tu n'as pas assez d'argent... \n")
	}
}
