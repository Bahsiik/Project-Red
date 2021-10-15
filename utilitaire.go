package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func Input() string { // Fonction pour récupérer le texte écrit dans le cmd et l'utiliser
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r\n", "", -1)
	return text
}

func IsLetter(s string) bool { // Fonction vérification alphanumérique
	phrase := []rune(s)
	counter := 0
	if s == "" {
		return false
	}
	for i := 0; i < len(phrase); i++ {
		if (phrase[i] >= 65 && phrase[i] <= 90) || (phrase[i] >= 97 && phrase[i] <= 122) {
			counter++
		}
	}
	return counter == len(s)
}

func Capitalize(s string) string { // Fonction mise en majuscule première lettre chaine de caractère
	var resultat string
	for i := range s {
		if IsAlpha(string(s[i])) { // Si lettre
			if IsUpper(string(s[i])) { // Si majuscule
				if i == 0 {
					resultat = resultat + string(s[i])
					continue
				} else if IsUpper(string(s[i-1])) {
					resultat = resultat + ToLower(string(s[i]))
					continue
				} else if IsLower(string(s[i-1])) {
					resultat = resultat + ToLower(string(s[i]))
					continue
				} else if s[i-1] >= 48 && s[i-1] <= 57 {
					resultat = resultat + ToLower(string(s[i]))
					continue
				} else {
					resultat = resultat + string(s[i])
					continue
				}
			} else if IsLower(string(s[i])) { // Si minuscule
				if i == 0 {
					resultat = resultat + ToUpper(string(s[i]))
					continue
				} else if !(IsAlpha(string(s[i-1]))) {
					resultat = resultat + ToUpper(string(s[i]))
					continue
				} else {
					resultat = resultat + string(s[i])
					continue
				}
			}
			resultat = resultat + string(s[i])
			continue
		}
		resultat = resultat + string(s[i])
	}
	return resultat
}

func IsAlpha(s string) bool { // Fonction reconnsaissance si lettre
	phrase := []rune(s)
	counter := 0
	for i := 0; i < len(phrase); i++ {
		if (phrase[i] >= 65 && phrase[i] <= 90) || (phrase[i] >= 97 && phrase[i] <= 122) || (phrase[i] >= 48 && phrase[i] <= 57) {
			counter++
		}
	}
	return counter == len(s)
}

func IsLower(s string) bool { // Fonction reconnsaissance minuscule
	phrase := []rune(s)
	counter := 0
	for i := 0; i < len(s); i++ {
		if phrase[i] >= 97 && phrase[i] <= 122 {
			counter++
		}
	}
	return counter == len(s)
}

func IsUpper(s string) bool { // Fonction reconnsaissance majuscule
	phrase := []rune(s)
	counter := 0
	for i := 0; i < len(s); i++ {
		if phrase[i] >= 65 && phrase[i] <= 90 {
			counter++
		}
	}
	return counter == len(s)
}

func ToLower(s string) string { // Mise en minuscule
	phrase := []rune(s)
	for i := 0; i < len(s); i++ {
		if phrase[i] >= 65 && phrase[i] <= 90 {
			phrase[i] += 32
		}
	}
	return string(phrase)
}

func ToUpper(s string) string { // Mise en majuscule
	phrase := []rune(s)
	for i := 0; i < len(s); i++ {
		if phrase[i] >= 97 && phrase[i] <= 122 {
			phrase[i] -= 32
		}
	}
	return string(phrase)
}

func DeathMonstre(p *Personnage, m *Monstre) { // Condition mort du monstre
	if m.hp <= 0 {
		fmt.Println("👑👑👑👑👑👑 ", p.nom, "a gagné le combat :))) uwu 👑👑👑👑👑👑") // Message fin de game
		GainExp(p, m)                                                      // Gain EXp
		GainNiveau(p)                                                      // Gain Niveau suivant Exp
		GainMoney(p, m)                                                    // Gain d'argent
		time.Sleep(1 * time.Second)
		fmt.Println()
		m.hp = m.hpmax // Réinitialisation pv monstre
		RetourMenu()
	}
}

func GainExp(p *Personnage, m *Monstre) { // Gain Exp suivant les monstres vaincus
	switch m.nom {
	case "Goblin":
		p.exp += 25
		fmt.Println("------ ", p.nom, " a gagné", p.exp, " Xp ------")
	case "Licorn":
		p.exp += 100
		fmt.Println("------ ", p.nom, " a gagné", p.exp, " Xp ------")
	case "Dragon":
		p.exp += 300
		fmt.Println("------ ", p.nom, " a gagné", p.exp, " Xp ------")
	case "Alan":
		p.exp += 500
		fmt.Println("------ ", p.nom, " a gagné", p.exp, " Xp ------")
	case "Lucas":
		p.exp += 1000
		fmt.Println("------ ", p.nom, " a gagné", p.exp, " Xp ------")
	}
}

func GainMoney(p *Personnage, m *Monstre) { // Gain Money suivant les monstres vaincus
	switch m.nom {
	case "Goblin":
		p.money += 10
		fmt.Println("------ ", p.nom, " a gagné", p.money, " ç ------")
	case "Licorn":
		p.money += 25
		fmt.Println("------ ", p.nom, " a gagné", p.money, " ç ------")
	case "Dragon":
		p.money += 75
		fmt.Println("------ ", p.nom, " a gagné", p.money, " ç ------")
	case "Alan":
		p.money += 100
		fmt.Println("------ ", p.nom, " a gagné", p.exp, " Xp ------")
	case "Lucas":
		p.money += 200
		fmt.Println("------ ", p.nom, " a gagné", p.exp, " Xp ------")
	}
}

func GainNiveau(p *Personnage) { // Gain Niveau suivant les monstres vaincus
	if p.exp == p.expmax {
		p.niveau += 1
		fmt.Println("------ ", p.nom, " est maintenant niveau : ", p.niveau, " ------")
		time.Sleep(1 * time.Second)
		p.expmax *= 1.5
		p.exp = 0
	} else if p.exp > p.expmax { // Condition si excédent de gain d'exp
		p.niveau += 1
		a := p.exp - p.expmax
		p.exp = 0
		p.exp += a
		p.expmax *= 1.5
		fmt.Println("------ ", p.nom, " est maintenant niveau : ", p.niveau, " et a ", p.exp, " Xp sur ", p.expmax, " XpMax ------")
		if p.classe == "Humain" { // Gain niveau donne bonus suivant classe
			fmt.Print("Attaque : ", p.atk, " >>> ")
			p.atk += 2
			fmt.Println(p.atk)
			fmt.Print("Puissance : ", p.puissance, " >>> ")
			p.puissance += 2
			fmt.Println(p.puissance)
		}
		if p.classe == "Elfe" { // Gain niveau donne bonus suivant classe
			fmt.Print("Attaque : ", p.atk, " >>> ")
			p.atk += 1
			fmt.Println(p.atk)
			fmt.Print("Puissance : ", p.puissance, " >>> ")
			p.puissance += 3
			fmt.Println(p.puissance)
		}
		if p.classe == "Nain" { // Gain niveau donne bonus suivant classe
			fmt.Print("Attaque : ", p.atk, " >>> ")
			p.atk += 3
			fmt.Println(p.atk)
			fmt.Print("Puissance : ", p.puissance, " >>> ")
			p.puissance += 1
			fmt.Println(p.puissance)
		}
		fmt.Println()
		time.Sleep(1 * time.Second)
		if p.exp > p.expmax {
			GainNiveau(p)
		}
	}
}

func EffacerTerminal() { // Commande d'effacement du terminal
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
}
