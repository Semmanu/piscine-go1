Instructions
Write a program that takes a string, and displays the first z character it encounters in it, 
followed by a newline ('\n'). If there are no z characters in the string, the program just writes z followed by a newline ('\n'). 
If the number of arguments is not 1, the program displays a z followed by a newline ('\n').

package main

import "github.com/01-edu/z01"

func main() {
	xyz := 'z'
	z01.PrintRune(xyz)
	z01.PrintRune('\n')
}
