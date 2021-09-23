package main

import "fmt"

func (p *Personnage) TakePot() {
	var test int
	for i := range p.inventaire {
		if i < len(p.inventaire) {
			if p.inventaire[i] == "Potion" {
				test++
				if p.hp == p.hpmax {
					fmt.Println(p.nom, "est déja full HP ! Il ne peut pas prendre de potion..")
				} else {
					fmt.Println(p.nom, "prend une potion de vie.")
					p.inventaire = RemoveInv(p.inventaire, "Potion", i)
					p.hp += 50
					if p.hp >= p.hpmax {
						p.hp = 3000
						fmt.Println(p.nom, "est maintenant full HP !")
					} else {
						fmt.Println(p.nom, "a maintenant", p.hp, "HP sur", p.hpmax, "HP.")
					}
				}
			}
		}
	}
	if test == 0 {
		fmt.Println(p.nom, "n'as malheureusement pas de potion...")
	}
}
