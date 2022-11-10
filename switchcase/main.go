package main
import (
	"github.com/01-edu/z01"
	"os"
)
func main () {
//func SwitchCase(c byte) byte {
//var c byte
//strAsByte := []byte(str)
args := os.Args
//var c rune
	for i := 1; i <= count; i++ {
		if i != 0 {
			c := []rune(args[i])
			for _, r := range c {
				if c >='0' && c <= '9' {
					return c - '0'
				}
				if c >= 'a' && c <= 'f' {
					return c - 'a' + 32
				}
				if c >= 'A' && c <= 'F' {
					return c - 'A' - 32
				}

			}
		}
	}
	return 0
}
