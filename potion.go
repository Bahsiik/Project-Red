package main

import (
	"fmt"
	"time"
)

func (p *Personnage) TakePot() {
	var test int
	for i := range p.inventaire {
		if i < len(p.inventaire) {
			if p.inventaire[i] == "Potion" {
				test++
				if p.hp == p.hpmax {
					fmt.Println(p.nom, "est déja full HP ! Il ne peut pas prendre de potion..")
					fmt.Println()
				} else {
					fmt.Println(p.nom, "prend une potion de vie.")
					p.RemoveInv("Potion", i)
					p.hp += 50
					if p.hp >= p.hpmax {
						p.hp = 100
						fmt.Println(p.nom, "est maintenant full HP !")
						fmt.Println()
					} else {
						fmt.Println(p.nom, "a maintenant", p.hp, "HP sur", p.hpmax, "HP.")
						fmt.Println()
					}
				}
				break
			}
		}
	}
	if test == 0 {
		fmt.Println(p.nom, "n'as malheureusement pas de potion...")
	}
}

func (p *Personnage) PoisonPot() {
	test := 0
	for i := 0; i < 3; i++ {
		if test == 0 {
			p.hp -= 10
			fmt.Println(p.nom, "a maintenant", p.hp, "Hp")
			if p.hp == 0 {
				p.Death()
				test++
			}
			time.Sleep(1 * time.Second)
		}
	}
}
