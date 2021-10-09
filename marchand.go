package main

import (
	"fmt"
)

var Marchand Personnage

func MarchandInit(p *Personnage) {
	Marchand.Init("Jeff Besos", "Marchand", 777, 777, 777, []string{"Potion : 3ç", "Potion de poison : 6ç", "Livre de sort: Boule de feu : 25ç", "Fourrure de Loup : 4ç", "Peau de troll : 7ç", "Cuir de Sanglier : 3ç", "Plume de Corbeau : 1ç"}, []string{"Coup de poing"}, 999)
}

func (p Personnage) DisplayInvMarchand() { // Fonction d'affichage de l'inventaire du marchand (Articles du magasin)
	fmt.Println("Bonjour, que souhaitez vous acheter ? (Numéro d'article / Rien)")
	if len(p.inventaire) == 0 {
		fmt.Println("Désolé, je n'ai rien a vous vendre...")
	} else {
		for i := range p.inventaire {
			fmt.Print(" Article ", i+1, " : ", p.inventaire[i], "\n")
		}
	}
	fmt.Println()
}

func (p Personnage) DisplayInvMarchand2() { // Fonction d'affichage de l'inventaire du marchand (Articles du magasin)
	fmt.Println("Que souhaitez vous acheter d'autres ? (Numéro d'article / Rien)")
	if len(p.inventaire) == 0 {
		fmt.Println("Désolé, je n'ai rien a vous vendre...")
	} else {
		for i := range p.inventaire {
			fmt.Print(" Article ", i+1, " : ", p.inventaire[i], "\n")
		}
	}
	fmt.Println()
}

func (p *Personnage) AccessInvMarchand() { // Fonction d'achat d'objet
	textmarchand := Input()
	switch textmarchand {
	case "Rien": // Retour au menu (aucun objet choisi)
		fmt.Println("Très bien, passez une bonne journée.")
		fmt.Println()
		RetourMenu()
	case "1": // Achat de la potion
		if (P1.money - 3) >= 0 { // Retrait cout d'achats pour personnage
			P1.money -= 3
			fmt.Println("Argent restant : ", P1.money, " Cacas d'or")
			fmt.Println()
			ContinueMarchandInv("Potion")
		} else {
			p.Pauvre()
		}
	case "2": // Achat de la potion de poison
		if (P1.money - 6) >= 0 {
			P1.money -= 6
			fmt.Println("Argent restant : ", P1.money, " Cacas d'or")
			fmt.Println()
			ContinueMarchandInv("Potion de poison")
		} else {
			p.Pauvre()
		}
	case "3": // Achat du skill "Boule de Feu"
		if p.VerifSpellBook("Livre de sort : Boule de Feu") && p.VerifSkill("Boule de Feu") {
			if (P1.money - 25) >= 0 {
				P1.money -= 25
				fmt.Println("Argent restant : ", P1.money, " Cacas d'or")
				fmt.Println()
				ContinueMarchandInv("Livre de sort : Boule de Feu")
			} else {
				p.Pauvre()
			}
		} else {
			fmt.Println("désolé monsieur, je ne peut pas vous fournir cet article..")
			P1.AccessInvMarchand()
		}
	case "4": // Achat de la Fourrure de Loup
		if (P1.money - 4) >= 0 {
			P1.money -= 4
			fmt.Println("Argent restant : ", P1.money, " Cacas d'or")
			fmt.Println()
			ContinueMarchandInv("Fourrure de Loup")
		} else {
			p.Pauvre()
		}
	case "5": // Achat de la Peau de troll
		if (P1.money - 7) >= 0 {
			P1.money -= 7
			fmt.Println("Argent restant : ", P1.money, " Cacas d'or")
			fmt.Println()
			ContinueMarchandInv("Peau de Troll")
		} else {
			p.Pauvre()
		}
	case "6": // Achat du Cuir de Sanglier
		if (P1.money - 3) >= 0 {
			P1.money -= 3
			fmt.Println("Argent restant : ", P1.money, " Cacas d'or")
			fmt.Println()
			ContinueMarchandInv("Cuir de Sanglier")
		} else {
			p.Pauvre()
		}
	case "7": // Achat de la Plume de Corbeau
		if (P1.money - 1) >= 0 {
			P1.money -= 1
			fmt.Println("Argent restant : ", P1.money, " Cacas d'or")
			fmt.Println()
			ContinueMarchandInv("Plume de Corbeau")
		} else {
			p.Pauvre()
		}
	default: // Choix d'objet à acheter invalide
		fmt.Println()
		fmt.Println("Désolé, je n'ai pas cette article, veuillez faire un autre choix.")
		fmt.Println()
		Marchand.DisplayInvMarchand2()
		Marchand.AccessInvMarchand()
	}
}

func ContinueMarchandInv(choix string) { // Fonction d'ajout de l'objet choisi dans l'inventaire + choix de continuer les achats ou non
	fmt.Println("Tenez Monsieur")
	fmt.Println("Besoin d'autres choses messire ? (Oui/Non)")
	P1.AddInventory(choix) // Ajout de l'objet
	for {
		textmarchand2 := Input()
		switch textmarchand2 {
		case "Oui": // Continuation des achats
			Marchand.DisplayInvMarchand2()
			fmt.Println()
			P1.AccessInvMarchand()
		case "Non": // Retour au menu
			fmt.Println("Très bien, au revoir")
			fmt.Println()
			RetourMenu()
		}
	}
}
