
package main
import (
	"fmt"
	// "strings"
	"os"
	"bufio"
)
func main() {
  for {
		inputReader := bufio.NewReader(os.Stdin)
		input, _ := inputReader.ReadString('\n')
		// fmt.Println(input)
		for i:=len(input)-2; i>= 0; i-- {
			// fmt.Print(string(input[i]))
			if string(input[i]) == " " {
				fmt.Println(len(input)-2-i)
				break
			}
		}
  }
}