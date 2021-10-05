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
		p.CraftItem("Plume de Corbeau", "Cuir de Sanglier", 1, 1, 5, "Chapeau de l'aventurier")
	}
}

func (p Personnage) ContinueForgeronInv(choix string) { // Fonction d'ajout de l'objet choisi dans l'inventaire + choix de continuer les achats ou non
	fmt.Println("Et voilà monsieur, c'est prêt !")
	fmt.Print("Il reste ", p.money, "ç à ", p.nom)
	fmt.Println("Besoin d'autres choses messire ? (Oui/Non)")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	textforgeron2, _ := reader.ReadString('\n')
	textforgeron2 = strings.Replace(textforgeron2, "\r\n", "", -1)
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

func (p *Personnage) CraftItem(elem1 string, elem2 string, nbr1 int, nbr2 int, prix int, equipement string) {
	var test1 int
	var test2 int
	for i := range p.inventaire {
		if i < len(p.inventaire) {
			if p.inventaire[i] == elem1 {
				test1++
			} else if p.inventaire[i] == elem2 {
				test2++
			}
		}
	}
	if test1 >= nbr1 && test2 >= nbr2 && p.money >= prix {
		p.money -= prix
		p.RemoveInv(elem1)
		p.RemoveInv(elem2)
		p.AddInventory(equipement)
		p.ContinueForgeronInv(equipement)
	} else if test1 < nbr1 && test2 < nbr2 {
		fmt.Print("Je suis désolé mais tu n'as pas les composants nécessaires...")
	} else if p.money < prix {
		fmt.Print("Je suis désolé mais tu n'as pas assez d'argent...")
	}
}
