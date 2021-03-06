package main

import (
	"fmt"
	"time"
)

var Forgeron Personnage

func ForgeronInit(p *Personnage) {
	Forgeron.Init("Mickey", "Forgeron", 666, 666, 666, 666, 0, []string{"Chapeau de l'aventurier : 10ç (1 Plume de Corbeau / 1 Cuir de Sanglier)", "Tunique de l'aventurier : 20ç (2 Fourrure de Loup / 1 Peau de Troll)", "Bottes de l'aventurier : 10ç (1 Fourrure de Loup / 1 Cuir de Sanglier)", "Casque de Dieu-Roi : 100ç (5 Plume de Corbeau / 5 Cuir de Sanglier)", "Armure de Dieu-Roi : 200ç (10 Fourrure de Loup / 10 Peau de Troll)", "Jambières de Dieu-Roi : 100ç (5 Fourrure de Loup / 5 Cuir de Sanglier)"}, 100, []string{"Coup de poing"}, 666, 5, 666, 666, 0, 100)
}

func (p Personnage) DisplayInvForgeron() { // Fonction d'affichage de l'inventaire du forgeron (Articles du magasin)
	fmt.Println("☭☭☭☭☭☭☭☭☭☭☭☭☭☭ Bonjour, que souhaitez vous crafter ? ☭☭☭☭☭☭☭☭☭☭☭☭☭☭")
	fmt.Println()
	for i := range p.inventaire {
		fmt.Print("Article ", i+1, " : ", p.inventaire[i], "\n")
	}
	fmt.Println()
	fmt.Println("0 - Retour au menu précédent")
	fmt.Println()
}

func (p Personnage) DisplayInvForgeron2() { // Fonction d'affichage de l'inventaire du forgeron apres achats (Articles du magasin)
	fmt.Println("☭☭☭☭☭☭☭☭☭☭☭☭☭☭ Que souhaitez vous crafter d'autres ? ☭☭☭☭☭☭☭☭☭☭☭☭☭☭")
	fmt.Println()
	if len(p.inventaire) == 0 {
		fmt.Println("Désolé, je n'ai rien a vous proposer...")
	} else {
		for i := range p.inventaire {
			fmt.Print("Article ", i+1, " : ", p.inventaire[i], "\n")
		}
	}
	fmt.Println()
	fmt.Println("0 - Retour au menu précédent")
	fmt.Println()
}

func (p *Personnage) AccessInvForgeron() { //Fonction craft d'objet
	textforgeron := Input()
	switch textforgeron {
	case "0": // Retour au menu (aucun objet choisi)
		fmt.Println("Très bien, passez une bonne journée.")
		fmt.Println()
		time.Sleep(1 * time.Second)
		p.Achats()
	case "1": // Craft Chapeau de l'aventurier
		p.CraftItem("Plume de Corbeau", "Cuir de Sanglier", 1, 1, 10, "Chapeau de l'aventurier")
	case "2":
		p.CraftItem("Fourrure de Loup", "Peau de Troll", 2, 1, 20, "Tunique de l'aventurier")
	case "3":
		p.CraftItem("Fourrure de Loup", "Cuir de Sanglier", 1, 1, 10, "Bottes de l'aventurier")
	case "4":
		p.CraftItem("Plume de Corbeau", "Cuir de Sanglier", 5, 5, 100, "Casque de Dieu-Roi")
	case "5":
		p.CraftItem("Fourrure de Loup", "Peau de Troll", 7, 5, 200, "Armure de Dieu-Roi")
	case "6":
		p.CraftItem("Fourrure de Loup", "Cuir de Sanglier", 5, 5, 100, "Jambières de Dieu-Roi")
	default:
		fmt.Println()
		fmt.Print("☭☭☭☭☭☭☭☭☭☭☭☭☭☭ Désolé mais je ne comprend pas, veuillez faire un autre choix ☭☭☭☭☭☭☭☭☭☭☭☭☭☭ \n")
		fmt.Println()
		Forgeron.DisplayInvForgeron2()
		p.AccessInvForgeron()
	}
}

func (p Personnage) ContinueForgeronInv(choix string) { // Fonction d'ajout de l'objet choisi dans l'inventaire + choix de continuer les achats ou non
	fmt.Println("Et voilà monsieur, c'est prêt !")
	fmt.Print("Il reste ", p.money, " ç à ", p.nom, "\n")
	fmt.Println()
	fmt.Println("Besoin d'autres choses messire ? (Oui/Non)")
	fmt.Println()
	p.AddInventory(choix) // Ajout de l'objet
	for {
		textforgeron2 := Input()
		switch textforgeron2 {
		case "Oui": // Continuation des achats
			Forgeron.DisplayInvForgeron2()
			fmt.Println()
			p.AccessInvForgeron()
		case "Non": // Retour au menu
			fmt.Println("☭☭☭☭☭☭☭☭☭☭☭☭☭☭ Très bien, au revoir ☭☭☭☭☭☭☭☭☭☭☭☭☭☭")
			fmt.Println()
			p.Achats()
		}
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
		p.EchecCraft()
	} else if p.money < prix { // Cas où il manque de l'argent
		fmt.Print("Je suis désolé mais tu n'as pas assez d'argent... \n")
		p.EchecCraft()
	}
}

func (p *Personnage) EchecCraft() { // Choix continuation achat ou non
	fmt.Println("☭☭☭☭☭☭☭☭☭☭☭☭☭☭ Besoin d'autres choses messire ? (Oui/Non) ☭☭☭☭☭☭☭☭☭☭☭☭☭☭")
	for {
		textforgeronechec := Input()
		switch textforgeronechec {
		case "Oui": // Continuation des achats
			Forgeron.DisplayInvForgeron2()
			fmt.Println()
			p.AccessInvForgeron()
		case "Non": // Retour au menu
			fmt.Println("Très bien, au revoir")
			fmt.Println()
			time.Sleep(1 * time.Second)
			p.Achats()
		}
	}
}
