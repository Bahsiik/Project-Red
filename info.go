package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (p Personnage) DisplayInfo() {
	fmt.Println("--- Informations du personnage ---")
	fmt.Println("Nom --> ", p.nom)
	fmt.Println("Classe --> ", p.classe)
	fmt.Println("Niveau --> ", p.niveau)
	fmt.Println("HP maximum --> ", p.hpmax)
	fmt.Println("HP actuels --> ", p.hp)
	fmt.Println()
	AccessInfo()
}

func AccessInfo() {
	fmt.Println("(Retour au menu --> Tapez retour)")
	fmt.Println()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	textinfo, _ := reader.ReadString('\n')
	// convert CRLF to LF
	textinfo = strings.Replace(textinfo, "\r\n", "", -1)
	switch textinfo {
	case "Retour":
		RetourMenu()
	}
}
