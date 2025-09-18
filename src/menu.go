package main

import (
	"fmt"
)

var player Character

func (player *Character) MainMenu() {
	for {
		fmt.Println("\t=== Menu Principal ===")
		fmt.Println("\t1 - Afficher les informations du personnage")
		fmt.Println("\t2 - Accéder à l'inventaire")
		fmt.Println("\t3 - Aller voir le marchand")
		fmt.Println("\t4 - Aller voir le forgeron")
		fmt.Println("\t0 - Quitter")
		fmt.Println("\tSélectionner un choix (1, 2, 3, 4 ou 0) :")
		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			fmt.Printf("%s\n - Classe: %s\n - Niveau: %d\n - PV: %d/%d\n - Or: %d\n",
				player.Nom, player.Classe, player.Niveau, player.Pv, player.PVMax, player.Or)
			if player.Classe == "Mage" {
				fmt.Printf("Mana: %d/%d\n", player.ManaActuel, player.ManaMax)
			}
			fmt.Println()

		case 2:
			player.AccessInventory()
		case 3:
			player.Merchant()
		case 4:
			player.MenuForgeron()

		case 0:
			fmt.Println("Merci d'avoir joué !")
			return

		default:
			fmt.Println("Choix invalide, réessayez.")
		}
	}
}
