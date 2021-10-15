package main

import (
	"fmt"
	"time"
)

var Marchand Personnage

func MarchandInit(p *Personnage) {
	Marchand.Init("Jeff Besos", "Marchand", 777, 777, 777, 0, 777, []string{"Potion de soin : 3ç", "Potion de mana : 3ç", "Potion de poison : 6ç", "Livre de sort: Boule de feu : 25ç", "Livre de sort: Iron Fist : 150ç", "Livre de sort : Blizzard : 200ç", "Livre de sort: Charge de Berserker : 500ç", "Livre de sort : Décharge énergétique : 750ç", "Fourrure de Loup : 4ç", "Peau de troll : 7ç", "Cuir de Sanglier : 3ç", "Plume de Corbeau : 1ç", "Sacoche de l'aventurier : 30ç"}, 10, []string{"Coup de poing"}, 999, 20, 999, 999, 0, 100)
}

func (p Personnage) DisplayInvMarchand() { // Fonction d'affichage de l'inventaire du marchand (Articles du magasin)
	fmt.Println("çççççççç Bonjour, que souhaitez vous acheter ? çççççççç")
	fmt.Println()
	if len(p.inventaire) == 0 {
		fmt.Println("Désolé, je n'ai rien a vous vendre...")
	} else {
		for i := range p.inventaire {
			fmt.Print("Article ", i+1, " : ", p.inventaire[i], "\n")
		}
	}
	fmt.Println()
	fmt.Println("0 - Retour au menu précédent")
	fmt.Println()
}

func (p Personnage) DisplayInvMarchand2() { // Fonction d'affichage de l'inventaire du marchand (Articles du magasin)
	fmt.Println("çççççççç Que souhaitez vous acheter d'autres ? çççççççç")
	fmt.Println()
	for i := range p.inventaire {
		fmt.Print("Article ", i+1, " : ", p.inventaire[i], "\n")
	}
	fmt.Println()
	fmt.Println("0 - Retour au menu précédent")
	fmt.Println()
}

func (p *Personnage) AccessInvMarchand() { // Fonction d'achat d'objet
	textmarchand := Input()
	switch textmarchand {
	case "0": // Retour au menu (aucun objet choisi)
		fmt.Println("çççççççç Très bien, passez une bonne journée. çççççççç")
		fmt.Println()
		time.Sleep(1 * time.Second)
		p.Achats()
	case "1": // Achat de la potion de vie
		p.AchatItem(3, "Potion de soin")
	case "2": // Achat de la potion de mana
		p.AchatItem(3, "Potion de mana")
	case "3": // Achat de la potion de poison
		p.AchatItem(3, "Potion de poison")
	case "4": // Achat du skill "Boule de Feu"
		p.AchatSkill(25, "Boule de Feu")
	case "5": // Achat du skill "Iron Fist"
		p.AchatSkill(150, "Iron Fist")
	case "6": // Achat du skill "Blizzard"
		p.AchatSkill(200, "Blizzard")
	case "7": // Achat du skill "Charge de Berserker"
		p.AchatSkill(500, "Charge de Berserker")
	case "8": // Achat du skill "Décharge énergétique"
		p.AchatSkill(750, "Décharge énergétique")
	case "9": // Achat de la Fourrure de Loup
		p.AchatItem(4, "Fourrure de Loup")
	case "10": // Achat de la Peau de Troll
		p.AchatItem(7, "Peau de Troll")
	case "11": // Achat du Cuir de Sanglier
		p.AchatItem(3, "Cuir de Sanglier")
	case "12": // Achat de la Plume de Corbeau
		p.AchatItem(1, "Plume de Corbeau")
	case "13": // Achat d'une sacoche (augmentation de l'inventaire, limité a 3 achats)
		if (p.money - 30) >= 0 {
			p.money -= 30
			fmt.Println("Argent restant : ", p.money, " Cacas d'or")
			fmt.Println()
			p.UpgradeInventorySlot()
		} else {
			p.Pauvre()
		}
	default: // Choix d'objet à acheter invalide
		fmt.Println()
		fmt.Println("Désolé, je n'ai pas cette article, veuillez faire un autre choix.")
		fmt.Println()
		Marchand.DisplayInvMarchand2()
		p.AccessInvMarchand()
	}
}

func (p *Personnage) ContinueMarchandInv(choix string) { // Fonction d'ajout de l'objet choisi dans l'inventaire + choix de continuer les achats ou non
	fmt.Println("Tenez Monsieur")
	fmt.Println()
	fmt.Println("Besoin d'autres choses messire ? (Oui/Non)")
	fmt.Println()
	p.AddInventory(choix) // Ajout de l'objet
	for {
		textmarchand2 := Input()
		switch textmarchand2 {
		case "Oui": // Continuation des achats
			Marchand.DisplayInvMarchand2()
			fmt.Println()
			p.AccessInvMarchand()
		case "Non": // Retour au menu
			fmt.Println()
			fmt.Println("Très bien, au revoir")
			fmt.Println()
			time.Sleep(1 * time.Second)
			p.Achats()
		}
	}
}

func (p *Personnage) AchatItem(prix int, item string) {
	if p.VerifTailleInv() {
		if (p.money - prix) >= 0 { // Retrait cout d'achats pour personnage
			p.money -= prix
			fmt.Println("Argent restant : ", p.money, " Cacas d'or")
			fmt.Println()
			p.ContinueMarchandInv(item)
		} else {
			p.Pauvre()
		}
	} else {
		fmt.Println("Désolez mais vous ne pouvez rien transporter de plus..")
		p.Achats()
	}
}

func (p *Personnage) AchatSkill(prix int, skill string) {
	if p.VerifTailleInv() {
		if !p.VerifSpellBook("Livre de sort : "+skill) && !p.VerifSkill(skill) {
			if (p.money - prix) >= 0 {
				p.money -= prix
				fmt.Println("Argent restant : ", p.money, " Cacas d'or")
				fmt.Println()
				p.ContinueMarchandInv("Livre de sort : " + skill)
			} else {
				p.Pauvre()
			}
		} else {
			fmt.Println("Désolé monsieur, je ne peut pas vous fournir cet article..")
			p.AccessInvMarchand()
		}
	} else {
		fmt.Println("Désolez mais vous ne pouvez rien transporter de plus..")
		p.Achats()
	}
}
