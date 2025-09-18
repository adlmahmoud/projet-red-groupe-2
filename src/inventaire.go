package main

import (
	"fmt"
)

<<<<<<< HEAD
func (player Character) AccessInventory() {
=======
func (player Character) accessInventory() {
>>>>>>> 25c08f92a67127ffc558f97092a48889d967b1a2
	fmt.Println("=== Inventaire du personnage ===")
	if len(player.Inventaire) == 0 {
		fmt.Println("\tInventaire vide")
	} else {
		for _, item := range player.Inventaire {
			fmt.Printf("\t- %s x %d\n", item.Nom, item.Quantite)
		}
	}
}

func (player *Character) MenuInventory() {
<<<<<<< HEAD
	for {
		player.AccessInventory()
=======
	for true {
		player.accessInventory()
>>>>>>> 25c08f92a67127ffc558f97092a48889d967b1a2
		fmt.Println("=== Menu inventaire ===")
		fmt.Printf("\t1 - Utiliser une potion de vie\n")
		fmt.Printf("\t0 - Retour\n")
		fmt.Println("Sélectionner un choix (1 ou 0) :")
		var userChose int
		fmt.Scan(&userChose)

		switch userChose {
		case 1:
<<<<<<< HEAD
			player.takePot()
		case 0:
=======
			player.takePotLife()
		case 2:
>>>>>>> 25c08f92a67127ffc558f97092a48889d967b1a2
			return
		default:
			fmt.Println("Erreur : choix non valide")
		}
	}
}
<<<<<<< HEAD

var items []string

func AjoutItem(item string) {
	if len(items) < 10 {
		items = append(items, item)
		fmt.Println("Item ajouté:", item)
	} else {
		fmt.Println("Nombre d'items autorisé dépassé, impossible d'ajouter:", item)
	}
}

func AfficherInventaire() {
	fmt.Println("=== Inventaire ===")
	for index, valeur := range items {
		fmt.Printf("%d: %s\n", index+1, valeur)
	}
	fmt.Println()
}
=======
>>>>>>> 25c08f92a67127ffc558f97092a48889d967b1a2
