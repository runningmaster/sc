package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/runningmaster/sc/internal/calc"
)

func main() {
	v, err := calc.Execute(strings.Join(os.Args[1:], " "))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(v)
}
