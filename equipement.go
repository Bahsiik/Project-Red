package main

import (
	"fmt"
	"time"
)

func (p Personnage) DisplayEquipment() { // Fonction d'affichage de l'equipement du personnage
	fmt.Print("--- Equipement de ", p.nom, " ---  \n")
	fmt.Print("Tête : ", p.Equipement.tete, "\n")
	fmt.Print("Torse : ", p.Equipement.torse, "\n")
	fmt.Print("Pieds : ", p.Equipement.pieds, "\n")
	fmt.Println()
}

func (p *Personnage) AccessEquipment() { // Fonction pour modifier l'equipement du personnage
	fmt.Println("------ Que va faire ", p.nom, " ? ------")
	fmt.Println()
	fmt.Println("1 - Retirer de l'équipement")
	fmt.Println()
	fmt.Println("0 - Retour à la gestion du personnage")
	textequip := Input()
	switch textequip {
	case "1":
		DisplayRetirerEquip()
	case "0":
		EffacerTerminal()
		fmt.Println()
		fmt.Println("------ Vous allez retourner au menu précédent ------")
		time.Sleep(1 * time.Second)
		Home()
	default:
		fmt.Println("------ ", P1.nom, "ne sais pas quoi faire.. ------")
		fmt.Println()
		SwipeMenuEquipement()
	}
}

func (p *Personnage) AddEquipementTete(obj string) { // Fonction d'ajout d'un objet a l'equipement Tete
	var test int
	for i := range p.inventaire {
		if p.inventaire[i] == obj {
			test++
			if p.Equipement.tete != "" { // Ajout à l'inventaire de l'item précédent
				P1.AddInventory(p.Equipement.tete)
			}
			if obj == "Chapeau de l'aventurier" {
				p.Equipement.tete = obj // Changement d'équipement de tête
				p.RemoveInv(obj)        // Retirer l'item de l'inventaire
				p.hpmax += 10           // Bonus d'équipement
				fmt.Println("Les Hp Max de ", p.nom, " passe à ", p.hpmax, " en équipant", obj)
				fmt.Println()
			}
		}
	}
	if test == 0 {
		fmt.Println("------ Vous ne possédez pas cet équipement.. ------")
		fmt.Println()
	}
}

func (p *Personnage) AddEquipementTorse(obj string) { // Fonction d'ajout d'un objet a l'equipement torse
	var test int
	for i := range p.inventaire {
		if p.inventaire[i] == obj {
			test++
			if p.Equipement.torse != "" {
				P1.AddInventory(p.Equipement.torse)
			}
			if obj == "Tunique de l'aventurier" {
				p.Equipement.torse = obj
				p.RemoveInv(obj)
				p.hpmax += 25
				fmt.Println("Les Hp Max de ", p.nom, " passe à ", p.hpmax, " en équipant ", obj)
			}
		}
	}
	if test == 0 {
		fmt.Println("------ Vous ne possédez pas cet équipement.. ------")
		fmt.Println()
	}
}

func (p *Personnage) AddEquipementPieds(obj string) { // Fonction d'ajout d'un objet a l'equipement pieds
	var test int
	for i := range p.inventaire {
		if p.inventaire[i] == obj {
			test++
			if p.Equipement.pieds != "" {
				P1.AddInventory(p.Equipement.pieds)
			}
			if obj == "Bottes de l'aventurier" {
				p.Equipement.pieds = obj
				p.RemoveInv(obj)
				p.hpmax += 15
				fmt.Println("Les Hp Max de ", p.nom, " passe à ", p.hpmax, " en équipant ", obj)
			}
		}
	}
	if test == 0 {
		fmt.Println("------ Vous ne possédez pas cet équipement.. ------")
		fmt.Println()
	}
}

func DisplayRetirerEquip() {
	fmt.Println()
	fmt.Println("------ Quelle partie d'équipement voulez vous retirer ? ------")
	fmt.Println()
	fmt.Println("1 - Tête")
	fmt.Println("2 - Torse")
	fmt.Println("3 - Pieds")
	fmt.Println()
	textretirerequip := Input()
	switch textretirerequip {
	case "1":
		P1.RetirerEquipementTete()
		SwipeMenuEquipement()
	case "2":
		P1.RetirerEquipementTorse()
		SwipeMenuEquipement()
	case "3":
		P1.RetirerEquipementPieds()
		SwipeMenuEquipement()
	}
}

func (p *Personnage) RetirerEquipementTete() { // Retrait équipement équiper pour le mettre dans inventaire si place dispo
	if p.VerifTailleInv() {
		if p.Equipement.tete == "" {
			fmt.Println("Vous n'avez rien équipé sur votre tête..")
			fmt.Println()
			SwipeMenuEquipement()
		} else {
			if p.Equipement.tete == "Chapeau de l'aventurier" {
				p.AddInventory(p.Equipement.tete)
				p.Equipement.tete = ""
				p.hpmax -= 10
				if p.hp > p.hpmax {
					p.hp = p.hpmax
				}
				fmt.Println(p.nom, " perd 10 HP maximum")
				fmt.Println()
			}
		}
	} else {
		fmt.Println("Vous devez d'abord vider votre inventaire..")
		fmt.Println()
		SwipeMenuEquipement()
	}
}

func (p *Personnage) RetirerEquipementTorse() {
	if p.VerifTailleInv() {
		if p.Equipement.torse == "" {
			fmt.Println("Vous n'avez rien équipé sur votre torse..")
			fmt.Println()
			SwipeMenuEquipement()
		} else {
			if p.Equipement.torse == "Tunique de l'aventurier" {
				p.AddInventory(p.Equipement.torse)
				p.Equipement.torse = ""
				p.hpmax -= 25
				if p.hp > p.hpmax {
					p.hp = p.hpmax
				}
				fmt.Println(p.nom, " perd 25 HP maximum")
				fmt.Println()
			}
		}
	} else {
		fmt.Println("Vous devez d'abord vider votre inventaire..")
		fmt.Println()
		SwipeMenuEquipement()
	}
}

func (p *Personnage) RetirerEquipementPieds() {
	if p.VerifTailleInv() {
		if p.Equipement.pieds == "" {
			fmt.Println("Vous n'avez rien équipé à vos pieds..")
			fmt.Println()
			SwipeMenuEquipement()
		} else {
			if p.Equipement.pieds == "Bottes de l'aventurier" {
				p.AddInventory(p.Equipement.pieds)
				p.Equipement.pieds = ""
				p.hpmax -= 25
				if p.hp > p.hpmax {
					p.hp = p.hpmax
				}
				fmt.Println(p.nom, " perd 15 HP maximum")
				fmt.Println()
			}
		}
	} else {
		fmt.Println("Vous devez d'abord vider votre inventaire..")
		fmt.Println()
		SwipeMenuEquipement()
	}
}

func SwipeMenuEquipement() { // retour choix equipement
	time.Sleep(1 * time.Second)
	EffacerTerminal()
	P1.DisplayEquipment()
	P1.AccessEquipment()
}
