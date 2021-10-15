package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func EffacerTerminal() { // Commande d'effacement du terminal
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
}
