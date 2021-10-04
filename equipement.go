package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (p Personnage) DisplayEquipment() { // Fonction d'affichage de l'equipement du personnage
	fmt.Print("--- Equipement de ", p.nom, " ---  \n")
	fmt.Print("Tête : ", p.Equipement.tete, "\n")
	fmt.Print("Torse : ", p.Equipement.torse, "\n")
	fmt.Print("Pieds : ", p.Equipement.pieds, "\n")
	fmt.Println()
}

func (p *Personnage) AccessEquipment() { // Fonction pour modifier l'equipement du personnage
	fmt.Print("Que va faire ", p.nom, " ? (Rien)\n")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	textequ, _ := reader.ReadString('\n')
	textequ = strings.Replace(textequ, "\r\n", "", -1)
	switch textequ {
	case "Rien":
		RetourMenu()
	default:
		fmt.Println(P1.nom, "ne sais pas quoi faire..")
		fmt.Println()
		P1.AccessEquipment()
	}
}
