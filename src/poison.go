package main

import (
	"fmt"
)

func (player *Character) takePotDegat() {
	for index := range player.Inventaire {
		if player.Inventaire[index].Nom == "Potion de vie" && player.Inventaire[index].Quantite > 0 {
			player.Pv -= 10
			if player.Pv <= 0 {
				fmt.Println("t'es mort comme un con")

			}
			fmt.Println("Potion de Poison utilisée (quantité -1)")
			fmt.Printf("Nouveau Pv : %d\n", player.Pv)
			player.Inventaire[index].Quantite -= 10
			if player.Inventaire[index].Quantite <= 0 {
				player.Inventaire = append(player.Inventaire[:index], player.Inventaire[index+1:]...)
			}
			return
		}
	}
	fmt.Println("Utilisation impossible : potion de Poison manquante")
}
