package main

import (
	"fmt"
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

func (p *Personnage) AddSkill(obj string) {
	p.skill = append(p.skill, obj)
}

func (p *Personnage) SpellBook(sort string) {
	verif := 0
	for i := range p.skill {
		if p.skill[i] == sort {
			verif += 1
		}
	}
	if verif == 0 {
		p.AddSkill(sort)
	}
}
