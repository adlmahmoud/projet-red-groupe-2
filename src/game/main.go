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

func main() {
	Degat := 0
	for i := 0; i <= Character.Pv; i++ {
		Degat := Character.Pv - i
	}
	fmt.Println(Degat)
}
