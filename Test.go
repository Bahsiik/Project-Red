package main

import "fmt"

type Personnage struct {
	nom        string
	classe     string
	niveau     int
	hpmax      int
	hp         int
	inventaire []string
}

func main() {
	var p1 Personnage
	p1.Init("Garen", "Combattant", 18, 3000, 2500, []string{"Potion", "Jambières du Berserker", "Arc-Bouclier de l'immortel", "Percepteur", "Lame d'infini", "Salutations de Dominik"})
	p1.displayInfo()
	fmt.Println()
	p1.accessInventory()
	fmt.Println()
	p1.takePot()

}

func (p *Personnage) Init(nom string, classe string, niveau int, hpmax int, hp int, inventaire []string) {
	p.nom = nom
	p.classe = classe
	p.niveau = niveau
	p.hpmax = hpmax
	p.hp = hp
	p.inventaire = inventaire
}

func (p Personnage) displayInfo() {
	fmt.Println("--- Informations du personnage ---")
	fmt.Println("Nom du personnage --> ", p.nom)
	fmt.Println("Classe du personnage --> ", p.classe)
	fmt.Println("Niveau du personnage --> ", p.niveau)
	fmt.Println("HP maximum du personnage --> ", p.hpmax)
	fmt.Println("HP actuels du personnage --> ", p.hp)
}

func (p Personnage) accessInventory() {
	fmt.Print("--- Inventaire de ", p.nom, " ---  \n")
	if len(p.inventaire) == 0 {
		fmt.Println("L'inventaire de", p.nom, "est vide...")
	} else {
		for i := range p.inventaire {
			fmt.Print(" Objet ", i+1, " : ", p.inventaire[i], "\n")
		}
	}
}

func (p *Personnage) takePot() {
	var test int
	for i := range p.inventaire {
		if p.inventaire[i] == "Potion" {
			test++
			if p.hp == p.hpmax {
				fmt.Println(p.nom, "est déja full HP ! Il ne peut pas prendre de potion..")
			} else {
				fmt.Println(p.nom, "prend une potion de vie.")
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
	if test == 0 {
		fmt.Println(p.nom, "n'as malheureusement pas de potion...")
	}
}
