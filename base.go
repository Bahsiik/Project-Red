package main

import "fmt"

func main() { // Fonction principale qui lance le jeu
	EffacerTerminal()
	fmt.Println("~~~~~~~~ Bienvenue dans ce jeu de dingue ! ~~~~~~~~")
	fmt.Println()
	P1.CharCreation()       // Initialisation du perso
	MarchandInit(&Marchand) // Initialisation du marchand
	ForgeronInit(&Forgeron) // Initialisation du forgeron
	GobelinInit(&Gobelin)   // Initialisation du gobelin
	LicorneInit(&Licorne)   // Initialisation de la plus belle des Licornees
	DragonInit(&Dragon)     // Initialisation DragonBlanc aux yeux bleus
	AlanInit(&Alan)         // Initialisation ALAN
	LucasInit(&Lucas)       // Initialisation LUCAS
	Menu()                  // Affichage du Menu
}
