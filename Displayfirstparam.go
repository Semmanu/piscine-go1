Des instructions
Écrivez  un  programme  qui  affiche  son  premier  argument , s'il  y en  a un.

paquet  principal
importer (
	"os"
	"github.com/01-edu/z01"
)
fonction  principale () {
	str  :=  os . Args                 //l'argument de la ligne de commande est transformé en tableau de string séparé par des quotes
	for  _ , char  :=  range  str [ 1 ] { //pour n'importer quelle position des valeurs char de notre string en position 1 dans le tableau str
		//la ligne du haut a transformé notre string 1, premier paramètre, en runes appelées char que notre z01 peut imprimer
		z01 . PrintRune ( char )
	}
	z01 . PrintRune ( '\n' )
}
