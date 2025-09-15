package main

import (
	"fmt"
	"os"
	Character "projet-red/character/main.go"
)

func Isdead(Character.Pv) {
	dead := []rune(Character.Pv)
	if dead == 0 {
		fmt.Println("bwaaa NULLLLL, recommence si ta une digniter")
		os.Exit(0)
	}
}
