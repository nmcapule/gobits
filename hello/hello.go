package main

import (
	"fmt"
	"os"
	"strings"
)

func greet(name string) {
	up := strings.ToUpper(name)
	fmt.Println("Hey", up, "what's up?")
}

func main() {
	who := "dude"
	if len(os.Args) > 1 {
		who = os.Args[1]
	}

	greet(who)
}
