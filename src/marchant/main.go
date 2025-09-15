package marchant

import "fmt"

func (player Character) accessInventory() {
	fmt.Println("=== Inventaire du personnage ===")
	if len(player.Inventaire) == 0 {
		fmt.Println("\tInventaire vide")
	} else {
		for _, item := range player.Inventaire {
			fmt.Printf("\t- %s x %d\n", item.Nom, item.Quantite)
		}
	}
}

// Menu permettant d’interagir avec l’inventaire
func (player *Character) MenuInventory() {
	for true {
		player.accessInventory() // Affiche l’inventaire
		fmt.Println("=== Menu inventaire ===")
		fmt.Printf("\t1 - Marchant\n")
		fmt.Printf("\t0 - Retour\n")
		fmt.Println("Sélectionner un choix (1 ou 0) :")
		var userChose int
		fmt.Scan(&userChose)

		switch userChose {
		case 1:
			player.marchant()
		case 0:
			return
		default:
			fmt.Println("Erreur : choix non valide")
		}
	}
}
