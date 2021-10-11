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

func IsLetter(s string) bool {
	phrase := []rune(s)
	counter := 0
	for i := 0; i < len(phrase); i++ {
		if (phrase[i] >= 65 && phrase[i] <= 90) || (phrase[i] >= 97 && phrase[i] <= 122) {
			counter++
		}
	}
	if counter == len(s) {
		return true
	}
	return false
}

func Capitalize(s string) string {
	var resultat string
	for i := range s {
		if IsAlpha(string(s[i])) {
			if IsUpper(string(s[i])) {
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
			} else if IsLower(string(s[i])) {
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

func IsAlpha(s string) bool {
	phrase := []rune(s)
	counter := 0
	for i := 0; i < len(phrase); i++ {
		if (phrase[i] >= 65 && phrase[i] <= 90) || (phrase[i] >= 97 && phrase[i] <= 122) || (phrase[i] >= 48 && phrase[i] <= 57) {
			counter++
		}
	}
	if counter == len(s) {
		return true
	}
	return false
}

func IsLower(s string) bool {
	phrase := []rune(s)
	counter := 0
	for i := 0; i < len(s); i++ {
		if phrase[i] >= 97 && phrase[i] <= 122 {
			counter++
		}
	}
	if counter == len(s) {
		return true
	}
	return false
}

func IsUpper(s string) bool {
	phrase := []rune(s)
	counter := 0
	for i := 0; i < len(s); i++ {
		if phrase[i] >= 65 && phrase[i] <= 90 {
			counter++
		}
	}
	if counter == len(s) {
		return true
	}
	return false
}

func ToLower(s string) string {
	phrase := []rune(s)
	for i := 0; i < len(s); i++ {
		if phrase[i] >= 65 && phrase[i] <= 90 {
			phrase[i] += 32
		}
	}
	return string(phrase)
}

func ToUpper(s string) string {
	phrase := []rune(s)
	for i := 0; i < len(s); i++ {
		if phrase[i] >= 97 && phrase[i] <= 122 {
			phrase[i] -= 32
		}
	}
	return string(phrase)
}
