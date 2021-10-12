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

func DeathMonstre(p *Personnage, m *Monstre) {
	if m.hp <= 0 {
		fmt.Println(p.nom, "a gagné le combat :))) uwu") // Message fin de game
		GainExp(p, m)
		GainNiveau(p)
		time.Sleep(1 * time.Second)
		fmt.Println()
		m.hp = m.hpmax // Réinitialisation pv monstre
		RetourMenu()
	}
}

func GainExp(p *Personnage, m *Monstre) {
	switch m.nom {
	case "Goblin d'entrainement":
		p.exp += 1000
		fmt.Println(p.nom, " a gagné", p.exp, " Xp")
	}
}

func GainNiveau(p *Personnage) {
	if p.exp == p.expmax {
		p.niveau += 1
		fmt.Println(p.nom, " est maintenant niveau : ", p.niveau)
		p.expmax *= 2
		p.exp = 0
	} else if p.exp > p.expmax {
		p.niveau += 1
		a := p.exp - p.expmax
		p.exp = 0
		p.exp += a
		p.expmax *= 2
		fmt.Println(p.nom, " est maintenant niveau : ", p.niveau, " et a ", p.exp, " Xp sur ", p.expmax, " XpMax")
		if p.exp > p.expmax {
			GainNiveau(p)
		}
	}
}

func EffacerTerminal() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
}
