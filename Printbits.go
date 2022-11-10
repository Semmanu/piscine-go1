/*Write a program that takes a number as argument, and prints it in binary value without a newline at the end.
If the the argument is not a number, the program should print 00000000.
*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for a := range os.Args[1] {
		if a >= 'a' && a <= 'z' || a >= 'A' && a <= 'Z' {
			fmt.Printf("00000000")
		}
	}
	if len(os.Args) == 2 {
		nbr, _ := strconv.Atoi(os.Args[1])
		printBits(byte(nbr))

	}

}
func printBits(octe byte) {
	fmt.Printf("%08b", octe)
}
