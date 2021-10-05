package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (p Personnage) DisplayInventory() { // Fonction affichage de l'inventaire
	fmt.Print("--- Inventaire de ", p.nom, " ---  \n")
	if len(p.inventaire) == 0 {
		fmt.Println("L'inventaire de", p.nom, "est vide...")
	} else {
		for i := range p.inventaire {
			fmt.Print(" Objet ", i+1, " : ", p.inventaire[i], "\n")
		}
	}
	fmt.Println()
}

func (p *Personnage) AccessInventory() { // Fonction d'utilisation d'objet dans l'inventaire
	fmt.Println("Quel objet", p.nom, "veut utiliser ? (Nom de l'objet / Rien)")
	fmt.Println()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	textinv, _ := reader.ReadString('\n')
	textinv = strings.Replace(textinv, "\r\n", "", -1)
	switch textinv {
	case "Potion": // Utilisation de la potion
		P1.TakePot()
		P1.DisplayInventory()
		textinv = ""
		P1.AccessInventory()
	case "Chapeau de l'aventurier":
		P1.AddInventory(p.Equipement.tete)
		P1.AddEquipementTete("Chapeau de l'aventurier")
		p.hpmax += 10
		fmt.Println("Les Hp Max de, ", p.nom, " passe à ", p.hpmax)
		P1.AccessInventory()
	case "Tunique de l'aventurier":
		P1.AddInventory(p.Equipement.torse)
		P1.AddEquipementTorse("Tunique de l'aventurier")
		p.hpmax += 25
		fmt.Println("Les Hp Max de, ", p.nom, " passe à ", p.hpmax)
		P1.AccessInventory()
	case "Bottes de l'aventurier":
		P1.AddInventory(p.Equipement.pieds)
		P1.AddEquipementPieds("Bottes de l'aventurier")
		p.hpmax += 15
		fmt.Println("Les Hp Max de, ", p.nom, " passe à ", p.hpmax)
		P1.AccessInventory()
	case "Rien": // Retour au menu
		RetourMenu()
	default: // Choix d'objet invalide
		fmt.Println(P1.nom, "ne sais pas quoi faire..")
		fmt.Println()
		P1.AccessInventory()
	}
}

func (p *Personnage) AddInventory(obj string) { // Fonction d'ajout d'un objet a l'inventaire
	p.inventaire = append(p.inventaire, obj)
}

func (p *Personnage) RemoveInv(obj string) { // Fonction de retrait d'un objet a l'inventaire
	for i := range p.inventaire {
		if i < len(p.inventaire) {
			if p.inventaire[i] == obj {
				p.inventaire = append(p.inventaire[:i], p.inventaire[i+1:]...)
				break
			}
		}
	}
}
