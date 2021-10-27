package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "G# r#cks!"
	replace := strings.NewReplacer("#", "o")
	fixed := replace.Replace(broken)
	fmt.Println(fixed)

}
