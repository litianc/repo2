package main 

import "fmt"
import "github.com/litianc/repo/tempconv"

func main() {
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.CToF(tempconv.BoilingC)) // "212°F"
}
