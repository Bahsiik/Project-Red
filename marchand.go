package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (p Personnage) DisplayInvMarchand() {
	fmt.Println("Bonjour, que souhaitez vous acheter ? (Numéro d'article / Rien)")
	if len(p.inventaire) == 0 {
		fmt.Println("Désolé, je n'ai rien a vous vendre...")
	} else {
		for i := range p.inventaire {
			fmt.Print(" Article ", i+1, " : ", p.inventaire[i], "\n")
		}
	}
	fmt.Println()
}

func (p *Personnage) AccessInvMarchand() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	text3, _ := reader.ReadString('\n')
	// convert CRLF to LF
	text3 = strings.Replace(text3, "\r\n", "", -1)
	switch text3 {
	case "Rien":
		fmt.Println("Très bien, passez une bonne journée.")
		fmt.Println()
		Menu()
	case "1":
		fmt.Println("Tenez Monsieur")
		fmt.Println("Besoin d'autres choses messire ? (Oui/Non)")
		fmt.Print("-> ")
		text4, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text4 = strings.Replace(text4, "\r\n", "", -1)
		P1.AddInventory("Potion")
		switch text4 {
		case "Oui":
			fmt.Println("Que désirez-vous d'autres chacal ?")
			fmt.Println()
			P1.AccessInvMarchand()
		case "Non":
			fmt.Println("Très bien, au revoir")
			fmt.Println()
			Menu()
		}
	case "2":
		fmt.Println("Désolé, cet article n'est pas encore disponible à la vente.")
		fmt.Println()
		P1.AccessInvMarchand()
	default:
		fmt.Println("Désolé, je n'ai pas cette article, veuillez faire un autre choix.")
		fmt.Println()
		P1.AccessInvMarchand()
	}
}
