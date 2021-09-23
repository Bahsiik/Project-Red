package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (p Personnage) DisplayInventory() {
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

func (p *Personnage) AccessInventory() {
	fmt.Println("Quel objet", p.nom, "veut utiliser ? (Nom de l'objet / Rien)")
	fmt.Println()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	text2, _ := reader.ReadString('\n')
	// convert CRLF to LF
	text2 = strings.Replace(text2, "\r\n", "", -1)
	switch text2 {
	case "Potion":
		P1.TakePot()
		P1.DisplayInventory()
		P1.AccessInventory()
	case "Rien":
		Menu()
	default:
		fmt.Println(P1.nom, "ne sais pas quoi faire..")
		fmt.Println()
		P1.AccessInventory()
	}
}

func (p *Personnage) AddInventory(obj string) {
	p.inventaire = append(p.inventaire, obj)
}

func (p *Personnage) RemoveInv(obj string, i int) {
	for i := range p.inventaire {
		if p.inventaire[i] == obj {
			p.inventaire = append(p.inventaire[:i], p.inventaire[i+1:]...)
		}
	}
}
