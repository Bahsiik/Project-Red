package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func DisplayInvMarchand() {
	marchandises := []string{"Potion", "Potion de Poison"}
	fmt.Println("Bonjour, que souhaitez vous acheter ?")
	for i := range marchandises {
		fmt.Print(" Article ", i+1, " : ", marchandises[i], "\n")
	}
}

func AccesInvMarchand() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	text2, _ := reader.ReadString('\n')
	// convert CRLF to LF
	text2 = strings.Replace(text2, "\r\n", "", -1)
	switch text2 {
	case "1":
		P1.TakePot()
		P1.DisplayInventory()
		P1.AccessInventory()
	case "2":
		fmt.Println("Désolé, cet article n'est pas encore disponible à la vente.")
	default:
		fmt.Println("Désolé, je n'ai pas cette article.")
		fmt.Println()
		P1.AccessInventory()
	}
}
