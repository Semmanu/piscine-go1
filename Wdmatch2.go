Instructions
Write a program that takes two string and checks whether it is possible to write the first string with characters from the second string. This rewrite must respect the order in which these characters appear in the second string.

If it is possible, the program displays the string followed by a newline ('\n'), otherwise it simply displays nothing.

If the number of arguments is different from 2, the program displays nothing.

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func ok(s1 string, s2 string) bool {
	runes1 := []rune(s1) //transformer les strings en tableaux de runes
	runes2 := []rune(s2)
	var rest string
	count := 0
	for i := 0; i < len(runes1); i++ { //créer i suivant la longueur de rune 1
		for j := count; j < len(runes2); j++ { //j initialisé à count, suivant la longueur de rune 2
			if runes1[i] == runes2[j] {
				rest += string(runes1[i]) //rest(qui est nul) = rest + la valeur de la rune postion i retransformée en string
				j = len(runes2) - 1
			}
			count++ //count fait avancer l'index de rune 2
		}
	}
	return s1 == rest //rest qui prend à chaque fois la valeur, ainsi il prend les valeur correspondantes dans l'ordre, dès que c !, vanhouan
}

func main() {
	if len(os.Args) == 3 {
		if ok(os.Args[1], os.Args[2]) {
			for _, char := range os.Args[1] { //en mode PrintStr
				z01.PrintRune(char)
			}
		}
	}
}
