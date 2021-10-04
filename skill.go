package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (p Personnage) DisplaySkill() {
	fmt.Print("--- Les sorts de ", p.nom, " sont ---  \n")
	if len(p.skill) == 0 {
		fmt.Println(" ", p.nom, "ne possède pas de sorts...")
	} else {
		for i := range p.skill {
			fmt.Print(" Sort ", i+1, " : ", p.skill[i], "\n")
		}
	}
	fmt.Println()
}

func (p *Personnage) AccessSkill() {
	fmt.Println("Quel sort", p.nom, "veut utiliser ? ( Rien)")
	fmt.Println()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	textskill, _ := reader.ReadString('\n')
	// convert CRLF to LF
	textskill = strings.Replace(textskill, "\r\n", "", -1)
	switch textskill {
	case "Rien":
		RetourMenu()
	default:
		fmt.Println(P1.nom, "ne sais pas quoi faire..")
		fmt.Println()
		P1.AccessInventory()
	}
}

func (p *Personnage) SpellBook(sort string) {
	verif := 0
	for i := range p.skill {
		if p.skill[i] == sort {
			verif += 1
		}
	}
	if verif == 0 {
		p.skill = append(p.skill, sort)
	}
}

func (p *Personnage) VerifSkill(sort string) bool {
	for i := range p.skill {
		if p.skill[i] == sort {
			return false
		}
	}
	return true
}
