package main

import "fmt"

func main() { // Fonction principale qui lance le jeu
	EffacerTerminal()
	fmt.Println("Bienvenue dans ce jeu de dingue !")
	P1.CharCreation()       // Initialisation du perso
	MarchandInit(&Marchand) // Initialisation du marchand
	ForgeronInit(&Forgeron) // Initialisation du forgeron
	GobelinInit(&Gobelin)   // Initialisation du gobelin
	Menu()                  // Affichage du Menu
}
