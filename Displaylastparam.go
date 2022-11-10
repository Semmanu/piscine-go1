Instructions
Write a program that displays its last argument, if there is one.

package main
import (
	"os"
	"github.com/01-edu/z01"
)
func main() {
	str := os.Args //transformer l'argument de la ligne en tableau de strings dénombrables
	count := 0
	for range str { //nous comptons les strings de str
		count++ //count est le nombre de strings donc paramètres qui augmente de 01 dès qu'on en trouve
	}
	i := count - 1                //n'oublions pas que range commence l'index à partir de 0, i the last position of Args is len(Arg) -01
	for _, char := range str[i] { //choisir les runes de notre string en last position
		z01.PrintRune(char) // les afficher
	}
	z01.PrintRune('\n')
}
