package inventaire

import (
	"fmt"
	"projet-red/inventaire"
)

func (player inventaire.Character) accessInventory() {
	fmt.Println("=== Inventaire du personnage ===")
	if len(player.Inventaire) == 0 {
		fmt.Println("\tInventaire vide")
	} else {
		for _, item := range player.Inventaire {
			fmt.Printf("\t- %s x %d\n", item.Nom, item.Quantite)
		}
	}
}

func (player *inventaire.Character) MenuInventory() {
	for true {
		player.accessInventory()
		fmt.Println("=== Menu inventaire ===")
		fmt.Printf("\t1 - Utiliser une potion de vie\n")
		fmt.Printf("\t0 - Retour\n")
		fmt.Println("Sélectionner un choix (1 ou 0) :")
		var userChose int
		fmt.Scan(&userChose)

		switch userChose {
		case 1:
			player.takePot()
		case 2:
			return
		default:
			fmt.Println("Erreur : choix non valide")
		}
	}
}
