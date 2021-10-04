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
	textmarchand, _ := reader.ReadString('\n')
	// convert CRLF to LF
	textmarchand = strings.Replace(textmarchand, "\r\n", "", -1)
	switch textmarchand {
	case "Rien":
		fmt.Println("Très bien, passez une bonne journée.")
		fmt.Println()
		RetourMenu()
	case "1":
		ContinueMarchandInv("Potion")
	case "2":
		ContinueMarchandInv("Potion de poison")
	case "3":
		ContinueMarchandSkill("Livre de sort: Boule de Feu")
	default:
		fmt.Println("Désolé, je n'ai pas cette article, veuillez faire un autre choix.")
		fmt.Println()
		P1.AccessInvMarchand()
	}
}

func ContinueMarchandInv(choix string) {
	fmt.Println("Tenez Monsieur")
	fmt.Println("Besoin d'autres choses messire ? (Oui/Non)")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	textmarchand2, _ := reader.ReadString('\n')
	// convert CRLF to LF
	textmarchand2 = strings.Replace(textmarchand2, "\r\n", "", -1)
	P1.AddInventory(choix)
	switch textmarchand2 {
	case "Oui":
		fmt.Println("Que désirez-vous d'autres chacal ?")
		fmt.Println()
		P1.AccessInvMarchand()
	case "Non":
		fmt.Println("Très bien, au revoir")
		fmt.Println()
		RetourMenu()
	}
}

func ContinueMarchandSkill(choix string) {
	fmt.Println("Tenez Monsieur")
	fmt.Println("Besoin d'autres choses messire ? (Oui/Non)")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	textmarchand2, _ := reader.ReadString('\n')
	// convert CRLF to LF
	textmarchand2 = strings.Replace(textmarchand2, "\r\n", "", -1)
	P1.SpellBook(choix)
	switch textmarchand2 {
	case "Oui":
		fmt.Println("Que désirez-vous d'autres chacal ?")
		fmt.Println()
		P1.AccessInvMarchand()
	case "Non":
		fmt.Println("Très bien, au revoir")
		fmt.Println()
		RetourMenu()
	}
}
