package main

func main() { // Fonction principale qui lance le jeu
	PersoInit(&P1)          // Initialisation du perso
	MarchandInit(&Marchand) // Initialisation du marchand
	ForgeronInit(&Forgeron) // Initialisation du forgeron
	GobelinInit(&Gobelin)   // Initialisation du gobelin
	Menu()                  // Affichage du Menu
}
