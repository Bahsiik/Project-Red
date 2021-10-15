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
		p.DisplayRetirerEquip()
	case "0":
		EffacerTerminal()
		fmt.Println()
		fmt.Println("------ Vous allez retourner au menu précédent ------")
		time.Sleep(1 * time.Second)
		p.Home()
	default:
		fmt.Println("------ ", p.nom, "ne sais pas quoi faire.. ------")
		fmt.Println()
		p.SwipeMenuEquipement()
	}
}

func (p *Personnage) AddEquipementTete(obj string) { // Fonction d'ajout d'un objet a l'equipement Tete
	var test int
	for i := range p.inventaire {
		if i < len(p.inventaire) {
			if p.inventaire[i] == obj {
				test++
				if p.Equipement.tete != "" { // Ajout à l'inventaire de l'item précédent
					p.RetirerEquipementTete()
				}
				if obj == "Chapeau de l'aventurier" {
					p.Equipement.tete = obj // Changement d'équipement de tête
					p.RemoveInv(obj)        // Retirer l'item de l'inventaire
					fmt.Println("En équipant ", obj, " ", p.nom, "à augmenté ses stats :")
					p.GainStats(15, 10, 5, 5)
				}
				if obj == "Casque de Dieu-Roi" {
					p.Equipement.tete = obj // Changement d'équipement de tête
					p.RemoveInv(obj)        // Retirer l'item de l'inventaire
					fmt.Println("En équipant ", obj, " ", p.nom, "à augmenté ses stats :")
					p.GainStats(150, 100, 15, 15)
				}
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
		if i < len(p.inventaire) {
			if p.inventaire[i] == obj {
				test++
				if p.Equipement.torse != "" {
					p.RetirerEquipementTorse()
				}
				if obj == "Tunique de l'aventurier" {
					p.Equipement.torse = obj
					p.RemoveInv(obj)
					fmt.Println("En équipant ", obj, " ", p.nom, "à augmenté ses stats :")
					p.GainStats(25, 25, 10, 10)
				}
				if obj == "Armure de Dieu-Roi" {
					p.Equipement.torse = obj
					p.RemoveInv(obj)
					fmt.Println("En équipant ", obj, " ", p.nom, "à augmenté ses stats :")
					p.GainStats(250, 250, 20, 20)
				}
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
		if i < len(p.inventaire) {
			if p.inventaire[i] == obj {
				test++
				if p.Equipement.pieds != "" {
					p.RetirerEquipementPieds()
				}
				if obj == "Bottes de l'aventurier" {
					p.Equipement.pieds = obj
					p.RemoveInv(obj)
					fmt.Println("En équipant ", obj, " ", p.nom, "à augmenté ses stats :")
					p.GainStats(10, 15, 5, 5)
				}
				if obj == "Jambières de Dieu-Roi" {
					p.Equipement.pieds = obj
					p.RemoveInv(obj)
					fmt.Println("En équipant ", obj, " ", p.nom, "à augmenté ses stats :")
					p.GainStats(100, 150, 15, 15)
				}
			}
		}
	}
	if test == 0 {
		fmt.Println("------ Vous ne possédez pas cet équipement.. ------")
		fmt.Println()
	}
}

func (p *Personnage) DisplayRetirerEquip() {
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
		p.RetirerEquipementTete()
		p.SwipeMenuEquipement()
	case "2":
		p.RetirerEquipementTorse()
		p.SwipeMenuEquipement()
	case "3":
		p.RetirerEquipementPieds()
		p.SwipeMenuEquipement()
	}
}

func (p *Personnage) RetirerEquipementTete() { // Retrait équipement équiper pour le mettre dans inventaire si place dispo
	if p.VerifTailleInv() {
		if p.Equipement.tete == "" {
			fmt.Println("Vous n'avez rien équipé sur votre tête..")
			fmt.Println()
			p.SwipeMenuEquipement()
		} else {
			if p.Equipement.tete == "Chapeau de l'aventurier" {
				p.AddInventory(p.Equipement.tete)
				p.Equipement.tete = ""
				p.PerteStats(15, 10, 5, 5)
				fmt.Println()
			}
			if p.Equipement.tete == "Casque de Dieu-Roi" {
				p.AddInventory(p.Equipement.tete)
				p.Equipement.tete = ""
				p.PerteStats(150, 100, 15, 15)
				fmt.Println()
			}
		}
	} else {
		fmt.Println("Vous devez d'abord vider votre inventaire..")
		fmt.Println()
		p.SwipeMenuEquipement()
	}
}

func (p *Personnage) RetirerEquipementTorse() {
	if p.VerifTailleInv() {
		if p.Equipement.torse == "" {
			fmt.Println("Vous n'avez rien équipé sur votre torse..")
			fmt.Println()
			p.SwipeMenuEquipement()
		} else {
			if p.Equipement.torse == "Tunique de l'aventurier" {
				p.AddInventory(p.Equipement.torse)
				p.Equipement.torse = ""
				p.PerteStats(15, 10, 5, 5)
				fmt.Println()
			}
			if p.Equipement.torse == "Armure de Dieu-Roi" {
				p.AddInventory(p.Equipement.torse)
				p.Equipement.torse = ""
				p.PerteStats(150, 100, 15, 15)
				fmt.Println()
			}
		}
	} else {
		fmt.Println("Vous devez d'abord vider votre inventaire..")
		fmt.Println()
		p.SwipeMenuEquipement()
	}
}

func (p *Personnage) RetirerEquipementPieds() {
	if p.VerifTailleInv() {
		if p.Equipement.pieds == "" {
			fmt.Println("Vous n'avez rien équipé à vos pieds..")
			fmt.Println()
			p.SwipeMenuEquipement()
		} else {
			if p.Equipement.pieds == "Bottes de l'aventurier" {
				p.AddInventory(p.Equipement.pieds)
				p.Equipement.pieds = ""
				p.PerteStats(10, 15, 5, 5)
				fmt.Println()
			}
			if p.Equipement.pieds == "Jambières de Dieu-Roi" {
				p.AddInventory(p.Equipement.pieds)
				p.Equipement.pieds = ""
				p.PerteStats(100, 150, 15, 15)
				fmt.Println()
			}
		}
	} else {
		fmt.Println("Vous devez d'abord vider votre inventaire..")
		fmt.Println()
		p.SwipeMenuEquipement()
	}
}

func (p *Personnage) SwipeMenuEquipement() { // retour choix equipement
	time.Sleep(1 * time.Second)
	EffacerTerminal()
	p.DisplayEquipment()
	p.AccessEquipment()
}
