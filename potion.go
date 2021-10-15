package main

import (
	"fmt"
	"time"
)

func (p *Personnage) TakeHealPot() { // Fonction de prise de potion de soins
	var test int
	for i := range p.inventaire {
		if i < len(p.inventaire) {
			if p.inventaire[i] == "Potion de soin" { // Incrémentattion du compteur de potions par rapport à l'inventaire
				test++
				if p.hp == p.hpmax { // Condition de non prise de potion en cas de FullHp
					fmt.Println(p.nom, "est déja full HP ! Il ne peut pas prendre de potion..")
					fmt.Println()
				} else {
					fmt.Println("----- ", p.nom, "prend une potion de vie -----")
					p.RemoveInv("Potion de soin") // Retrait de la potion de l'inventaire après son utilisation
					p.hp += 50                    // Gain de santé par potion
					if p.hp >= p.hpmax {          // Condition en cas d'excédent de soin
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
		fmt.Println(p.nom, "n'as malheureusement pas de potion de soin...")
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

func (p *Personnage) TakeManaPot() { // Fonction de prise de potion de soins
	var test int
	for i := range p.inventaire {
		if i < len(p.inventaire) {
			if p.inventaire[i] == "Potion de mana" { // Incrémentattion du compteur de potions par rapport à l'inventaire
				test++
				if p.hp == p.hpmax { // Condition de non prise de potion en cas de FullHp
					fmt.Println("Le mana de", p.nom, " est déjà plein !")
					fmt.Println()
				} else {
					fmt.Println("----- ", p.nom, "prend une potion de mana -----")
					p.RemoveInv("Potion de mana") // Retrait de la potion de l'inventaire après son utilisation
					p.hp += 50                    // Gain de santé par potion
					if p.mana >= p.manamax {      // Condition en cas d'excédent de soin
						p.hp = 100
						fmt.Println("----- ", p.nom, "est maintenant full mana ! -----")
						fmt.Println()
					} else { // Affichage des données final après prise de potion sous condition
						fmt.Println("----- ", p.nom, "a maintenant", p.mana, "mana sur", p.manamax, "mana -----")
						fmt.Println()
					}
				}
				break
			}
		}
	}
	if test == 0 { // Dans l'absence de potions dans l'inventaire
		fmt.Println(p.nom, "n'as malheureusement pas de potion de mana...")
	}
}

func (p *Personnage) PoisonPotComb(m *Monstre) { // Fonction potion de poison
	var test2 int
	for i := range p.inventaire {
		if i < len(p.inventaire) {
			if p.inventaire[i] == "Potion de poison" { // Incrémentattion du compteur de potions par rapport à l'inventaire
				test2++
				fmt.Println(p.nom, " utilise une potion de poison sur ", m.nom)
				fmt.Println()
				p.RemoveInv("Potion de poison")
				test := 0
				for i := 0; i < 3; i++ { // Initialisation des 3s d'effet de la potion
					if test == 0 { // Dégats/sec
						m.hp -= 10
						fmt.Println()
						fmt.Println(m.nom, " a maintenant ", m.hp, "HP sur", m.hpmax, "HP.") // Affichage pv monstre fin tour
						fmt.Println()
						if m.hp <= 0 {
							fmt.Println()
							fmt.Println("👑👑👑👑👑👑 ", p.nom, "a gagné le combat :))) uwu 👑👑👑👑👑👑") // Message fin de game
							fmt.Println()
							GainExp(p, m)
							GainMoney(p, m)
							fmt.Println()
							m.hp = m.hpmax // Réinitialisation pv monstre
							RetourMenu()
						}
						time.Sleep(1 * time.Second) // Timer dégats par secondes
					}
				}
			}
		}
	}
	if test2 == 0 { // Dans l'absence de potions dans l'inventaire de combat
		fmt.Println()
		fmt.Println(p.nom, "n'as malheureusement pas cette potion...")
		fmt.Println()
		p.CharTurn(m)
	}
}
