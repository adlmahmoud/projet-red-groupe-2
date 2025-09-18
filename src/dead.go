package main

import (
	"fmt"
	"os"
)

var p Character

func Isdead(player Character) {
	if p.Pv <= 0 {
		fmt.Println("bwaaa NULLLLL, recommence si t'as une dignité !")
		os.Exit(0)
	}
}
