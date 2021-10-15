package main

import (
	"fmt"
	"time"
)

func (p *Personnage) TakePot() { // Fonction de prise de potion de soins
	var test int
	for i := range p.inventaire {
		if i < len(p.inventaire) {
			if p.inventaire[i] == "Potion" { // Incrémentattion du compteur de potions par rapport à l'inventaire
				test++
				if p.hp == p.hpmax { // Condition de non prise de potion en cas de FullHp
					fmt.Println(p.nom, "est déja full HP ! Il ne peut pas prendre de potion..")
					fmt.Println()
				} else {
					fmt.Println("----- ", p.nom, "prend une potion de vie -----")
					p.RemoveInv("Potion") // Retrait de la potion de l'inventaire après son utilisation
					p.hp += 50            // Gain de santé par potion
					if p.hp >= p.hpmax {  // Condition en cas d'excédent de soin
						p.hp = 100
						fmt.Println("----- ", p.nom, "est maintenant full HP ! -----")
						fmt.Println()
					} else { // Affichage des données final après prise de potion sous condition
						fmt.Println("----- ", p.nom, "a maintenant", p.hp, "HP sur", p.hpmax, "HP -----")
						fmt.Println()
					}
				}
				break
			}
		}
	}
	if test == 0 { // Dans l'absence de potions dans l'inventaire
		fmt.Println(p.nom, "n'as malheureusement pas de potion...")
	}
}

func (p *Personnage) PoisonPot() { // Fonction potion de poison
	var test2 int
	for i := range p.inventaire {
		if i < len(p.inventaire) {
			if p.inventaire[i] == "Potion de poison" { // Incrémentattion du compteur de potions par rapport à l'inventaire
				test2++
				fmt.Println(p.nom, " prend une potion de poison.")
				p.RemoveInv("Potion de poison") // Retrait de la potion de l'inventaire après son utilisation
				test := 0
				for i := 0; i < 3; i++ { // Initialisation des 3s d'effet de la potion
					if test == 0 { // Dégats/sec
						p.hp -= 10
						fmt.Println(p.nom, "a maintenant", p.hp, "HP sur", p.hpmax, "HP.")
						if p.hp <= 0 { // Condition de mort
							p.Death()
							test++
						}
						time.Sleep(1 * time.Second) // Timer dégats par secondes
					}
				}
			}
		}
	}
	if test2 == 0 { // Dans l'absence de potions dans l'inventaire
		fmt.Println(p.nom, "n'as malheureusement pas cette potion...")
		fmt.Println()
	}
}
