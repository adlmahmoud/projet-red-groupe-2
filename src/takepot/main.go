package takepot

import (
	"fmt"
	"projet-red/takepot"
)

func (player *takepot.Character) takePot() {
	for index := range player.Inventaire {
		if player.Inventaire[index].Nom == "Potion de vie" && player.Inventaire[index].Quantite > 0 {
			player.Pv += 50
			if player.Pv > player.PvMax {
				player.Pv = player.PvMax
			}
			fmt.Println("Potion de vie utilisée (quantité -1)")
			fmt.Printf("Nouveau Pv : %d\n", player.Pv)
			player.Inventaire[index].Quantite -= 1
			if player.Inventaire[index].Quantite <= 0 {
				player.Inventaire = append(player.Inventaire[:index], player.Inventaire[index+1:]...)
			}
			return
		}
	}
	fmt.Println("Utilisation impossible : potion de vie manquante")
}
