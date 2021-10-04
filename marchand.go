package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

func (p *Personnage) AccessInvMarchand() { // Fonction d'achat d'objet
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	textmarchand, _ := reader.ReadString('\n')
	textmarchand = strings.Replace(textmarchand, "\r\n", "", -1)
	switch textmarchand {
	case "Rien": // Retour au menu (aucun objet choisi)
		fmt.Println("Très bien, passez une bonne journée.")
		fmt.Println()
		RetourMenu()
	case "1": // Achat de la potion
		ContinueMarchandInv("Potion")
	case "2": // Achat de la potion de poison
		ContinueMarchandInv("Potion de poison")
	case "3": // Achat du skill "Boule de Feu"
		if p.VerifSkill("Livre de sort: Boule de Feu") {
			ContinueMarchandSkill("Livre de sort: Boule de Feu")
		} else {
			fmt.Println("désolé monsieur, je ne peut pas vous fournir cet article..")
			P1.AccessInvMarchand()
		}

	default: // Choix d'objet à acheter invalide
		fmt.Println("Désolé, je n'ai pas cette article, veuillez faire un autre choix.")
		fmt.Println()
		P1.AccessInvMarchand()
	}
}

func ContinueMarchandInv(choix string) { // Fonction d'ajout de l'objet choisi dans l'inventaire + choix de continuer les achats ou non
	fmt.Println("Tenez Monsieur")
	fmt.Println("Besoin d'autres choses messire ? (Oui/Non)")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	textmarchand2, _ := reader.ReadString('\n')
	textmarchand2 = strings.Replace(textmarchand2, "\r\n", "", -1)
	P1.AddInventory(choix) // Ajout de l'objet
	switch textmarchand2 {
	case "Oui": // Continuation des achats
		fmt.Println("Que désirez-vous d'autres chacal ?")
		fmt.Println()
		P1.AccessInvMarchand()
	case "Non": // Retour au menu
		fmt.Println("Très bien, au revoir")
		fmt.Println()
		RetourMenu()
	}
}

func ContinueMarchandSkill(choix string) { // Fonction d'ajout du sort choisi dans l'inventaire + choix de continuer les achats ou non
	fmt.Println("Tenez Monsieur")
	fmt.Println("Besoin d'autres choses messire ? (Oui/Non)")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	textmarchand2, _ := reader.ReadString('\n')
	textmarchand2 = strings.Replace(textmarchand2, "\r\n", "", -1)
	P1.SpellBook(choix) // Ajout du sort
	switch textmarchand2 {
	case "Oui": // Continuation des achats
		fmt.Println("Que désirez-vous d'autres chacal ?")
		fmt.Println()
		P1.AccessInvMarchand()
	case "Non": // Retour au menu
		fmt.Println("Très bien, au revoir")
		fmt.Println()
		RetourMenu()
	}
}
