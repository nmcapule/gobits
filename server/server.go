package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func reverse(s string) string {
	sl := len(s)

	cs := []rune(s)
	for i := 0; i < sl/2; i++ {
		cs[i], cs[sl-i-1] = cs[sl-i-1], cs[i]
	}

	return string(cs)
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	who := r.FormValue("name")
	if who == "" {
		who = "someone"
	}
	dwho := reverse(who)

	fmt.Fprintf(w, "Hello %s! I'm your fiendly doppelganger %s!", strings.Title(who), strings.Title(dwho))
}

func main() {
	http.HandleFunc("/greet", greetHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
