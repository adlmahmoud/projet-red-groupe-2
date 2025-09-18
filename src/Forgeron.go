package main

import "fmt"

func (player *Character) MenuForgeron() {
	for {
		fmt.Println("\n=== FORGERON ===")
		fmt.Printf("Votre or actuel : %d pièces\n", player.Or)
		fmt.Println("Équipements disponibles :")
		fmt.Println("1 - Épée : 30 pièces d'or (40 dégâts)")
		fmt.Println("2 - Lance : 80 pièces d'or (70 dégâts)")
		fmt.Println("3 - Hache de guerre : 150 pièces d'or (120 dégâts)")
		fmt.Println("0 - Retour")
		fmt.Print("Choisissez un équipement à fabriquer : ")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			player.FabriquerObjet("Épée", 30)
		case 2:
			player.FabriquerObjet("Lance", 80)
		case 3:
			player.FabriquerObjet("Hache de guerre", 150)
		case 0:
			return
		default:
			fmt.Println("Choix invalide, réessayez.")
		}
	}
}
func (player *Character) FabriquerObjet(nom string, prix int) {
	if player.Or < prix {
		fmt.Printf("Vous n'avez pas assez d'or pour fabriquer %s. Il vous faut %d pièces.\n", nom, prix)
		return
	}

	player.Or -= prix

	for i, item := range player.Inventaire {
		if item.Nom == nom {
			player.Inventaire[i].Quantite++
			fmt.Printf("Vous avez fabriqué un(e) %s ! Quantité : %d. Or restant : %d pièces.\n",
				nom, player.Inventaire[i].Quantite, player.Or)
			return
		}
	}

	player.Inventaire = append(player.Inventaire, Item{Nom: nom, Quantite: 1})
	fmt.Printf("Vous avez fabriqué : %s ! Or restant : %d pièces.\n", nom, player.Or)
}
