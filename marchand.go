package main

import (
	"fmt"
	"time"
)

var Marchand Personnage

func MarchandInit(p *Personnage) {
	Marchand.Init("Jeff Besos", "Marchand", 777, 777, 777, 0, 777, []string{"Potion de soin : 3ç", "Potion de mana : 3ç", "Potion de poison : 6ç", "Livre de sort: Boule de feu : 25ç", "Livre de sort : Blizzard : 200ç", "Livre de sort : Décharge énergétique : 750ç", "Fourrure de Loup : 4ç", "Peau de troll : 7ç", "Cuir de Sanglier : 3ç", "Plume de Corbeau : 1ç", "Sacoche de l'aventurier : 30ç"}, 10, []string{"Coup de poing"}, 999, 20, 999, 999, 0, 100)
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

func (p *Personnage) AccessInvMarchand() { // Fonction d'achat d'objet
	textmarchand := Input()
	switch textmarchand {
	case "0": // Retour au menu (aucun objet choisi)
		fmt.Println("çççççççç Très bien, passez une bonne journée. çççççççç")
		fmt.Println()
		time.Sleep(1 * time.Second)
		Achats()
	case "1": // Achat de la potion
		if p.VerifTailleInv() {
			if (P1.money - 3) >= 0 { // Retrait cout d'achats pour personnage
				P1.money -= 3
				fmt.Println("Argent restant : ", P1.money, " Cacas d'or")
				fmt.Println()
				ContinueMarchandInv("Potion de soin")
			} else {
				p.Pauvre()
			}
		} else {
			fmt.Println("Désolez mais vous ne pouvez rien transporter de plus..")
			Achats()
		}
	case "2": // Achat de la potion
		if p.VerifTailleInv() {
			if (P1.money - 3) >= 0 { // Retrait cout d'achats pour personnage
				P1.money -= 3
				fmt.Println("Argent restant : ", P1.money, " Cacas d'or")
				fmt.Println()
				ContinueMarchandInv("Potion de mana")
			} else {
				p.Pauvre()
			}
		} else {
			fmt.Println("Désolez mais vous ne pouvez rien transporter de plus..")
			Achats()
		}
	case "3": // Achat de la potion de poison
		if p.VerifTailleInv() {
			if (P1.money - 6) >= 0 {
				P1.money -= 6
				fmt.Println("Argent restant : ", P1.money, " Cacas d'or")
				fmt.Println()
				ContinueMarchandInv("Potion de poison")
			} else {
				p.Pauvre()
			}
		} else {
			fmt.Println("Désolez mais vous ne pouvez rien transporter de plus..")
			Achats()
		}
	case "4": // Achat du skill "Boule de Feu"
		if p.VerifTailleInv() {
			if !p.VerifSpellBook("Livre de sort : Boule de Feu") && !p.VerifSkill("Boule de Feu") {
				if (P1.money - 25) >= 0 {
					P1.money -= 25
					fmt.Println("Argent restant : ", P1.money, " Cacas d'or")
					fmt.Println()
					ContinueMarchandInv("Livre de sort : Boule de Feu")
				} else {
					p.Pauvre()
				}
			} else {
				fmt.Println("Désolé monsieur, je ne peut pas vous fournir cet article..")
				P1.AccessInvMarchand()
			}
		} else {
			fmt.Println("Désolez mais vous ne pouvez rien transporter de plus..")
			Achats()
		}
<<<<<<< HEAD
	case "5": // Achat du skill "Blizzard"
		if !p.VerifSpellBook("Livre de sort : Blizzard") && !p.VerifSkill("Blizzard") {
			if (P1.money - 200) >= 0 {
				P1.money -= 200
				fmt.Println("Argent restant : ", P1.money, " Cacas d'or")
				fmt.Println()
				ContinueMarchandInv("Livre de sort : Blizzard")
			} else {
				p.Pauvre()
			}
		} else {
			fmt.Println("Désolé monsieur, je ne peut pas vous fournir cet article..")
			P1.AccessInvMarchand()
		}
	case "6": // Achat du skill "Décharge énergétique"
		if !p.VerifSpellBook("Livre de sort : Décharge énergétique") && !p.VerifSkill("Décharge énergétique") {
			if (P1.money - 750) >= 0 {
				P1.money -= 750
				fmt.Println("Argent restant : ", P1.money, " Cacas d'or")
				fmt.Println()
				ContinueMarchandInv("Livre de sort : Décharge énergétique")
			} else {
				p.Pauvre()
			}
		} else {
			fmt.Println("Désolé monsieur, je ne peut pas vous fournir cet article..")
			P1.AccessInvMarchand()
		}
	case "9": // Achat de la Fourrure de Loup
		if (P1.money - 4) >= 0 {
			P1.money -= 4
			fmt.Println("Argent restant : ", P1.money, " Cacas d'or")
			fmt.Println()
			ContinueMarchandInv("Fourrure de Loup")
=======
	case "5": // Achat de la Fourrure de Loup
		if p.VerifTailleInv() {
			if (P1.money - 4) >= 0 {
				P1.money -= 4
				fmt.Println("Argent restant : ", P1.money, " Cacas d'or")
				fmt.Println()
				ContinueMarchandInv("Fourrure de Loup")
			} else {
				p.Pauvre()
			}
>>>>>>> f468b78bed614db57eb524f3946ce7a03c63cc77
		} else {
			fmt.Println("Désolez mais vous ne pouvez rien transporter de plus..")
			Achats()
		}
<<<<<<< HEAD
	case "12": // Achat de la Peau de troll
		if (P1.money - 7) >= 0 {
			P1.money -= 7
			fmt.Println("Argent restant : ", P1.money, " Cacas d'or")
			fmt.Println()
			ContinueMarchandInv("Peau de Troll")
=======
	case "6": // Achat de la Peau de troll
		if p.VerifTailleInv() {
			if (P1.money - 7) >= 0 {
				P1.money -= 7
				fmt.Println("Argent restant : ", P1.money, " Cacas d'or")
				fmt.Println()
				ContinueMarchandInv("Peau de Troll")
			} else {
				p.Pauvre()
			}
>>>>>>> f468b78bed614db57eb524f3946ce7a03c63cc77
		} else {
			fmt.Println("Désolez mais vous ne pouvez rien transporter de plus..")
			Achats()
		}
<<<<<<< HEAD
	case "13": // Achat du Cuir de Sanglier
		if (P1.money - 3) >= 0 {
			P1.money -= 3
			fmt.Println("Argent restant : ", P1.money, " Cacas d'or")
			fmt.Println()
			ContinueMarchandInv("Cuir de Sanglier")
=======
	case "7": // Achat du Cuir de Sanglier
		if p.VerifTailleInv() {
			if (P1.money - 3) >= 0 {
				P1.money -= 3
				fmt.Println("Argent restant : ", P1.money, " Cacas d'or")
				fmt.Println()
				ContinueMarchandInv("Cuir de Sanglier")
			} else {
				p.Pauvre()
			}
>>>>>>> f468b78bed614db57eb524f3946ce7a03c63cc77
		} else {
			fmt.Println("Désolez mais vous ne pouvez rien transporter de plus..")
			Achats()
		}
<<<<<<< HEAD
	case "14": // Achat de la Plume de Corbeau
		if (P1.money - 1) >= 0 {
			P1.money -= 1
			fmt.Println("Argent restant : ", P1.money, " Cacas d'or")
			fmt.Println()
			ContinueMarchandInv("Plume de Corbeau")
=======
	case "8": // Achat de la Plume de Corbeau
		if p.VerifTailleInv() {
			if (P1.money - 1) >= 0 {
				P1.money -= 1
				fmt.Println("Argent restant : ", P1.money, " Cacas d'or")
				fmt.Println()
				ContinueMarchandInv("Plume de Corbeau")
			} else {
				p.Pauvre()
			}
>>>>>>> f468b78bed614db57eb524f3946ce7a03c63cc77
		} else {
			fmt.Println("Désolez mais vous ne pouvez rien transporter de plus..")
			Achats()
		}
<<<<<<< HEAD
	case "15": // Achat d'une sacoche (augmentation de l'inventaire, limité a 3 achats)
		if (P1.money - 30) >= 0 {
			P1.money -= 30
			fmt.Println("Argent restant : ", P1.money, " Cacas d'or")
			fmt.Println()
			P1.UpgradeInventorySlot()
=======
	case "9": // Achat d'une sacoche (augmentation de l'inventaire, limité a 3 achats)
		if p.VerifTailleInv() {
			if (P1.money - 30) >= 0 {
				P1.money -= 30
				fmt.Println("Argent restant : ", P1.money, " Cacas d'or")
				fmt.Println()
				P1.UpgradeInventorySlot()
			} else {
				p.Pauvre()
			}
>>>>>>> f468b78bed614db57eb524f3946ce7a03c63cc77
		} else {
			fmt.Println("Désolez mais vous ne pouvez rien transporter de plus..")
			Achats()
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
	fmt.Println()
	fmt.Println("Besoin d'autres choses messire ? (Oui/Non)")
	fmt.Println()
	P1.AddInventory(choix) // Ajout de l'objet
	for {
		textmarchand2 := Input()
		switch textmarchand2 {
		case "Oui": // Continuation des achats
			Marchand.DisplayInvMarchand2()
			fmt.Println()
			P1.AccessInvMarchand()
		case "Non": // Retour au menu
			fmt.Println()
			fmt.Println("Très bien, au revoir")
			fmt.Println()
			time.Sleep(1 * time.Second)
			Achats()
		}
	}
}
