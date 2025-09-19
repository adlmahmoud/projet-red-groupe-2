package main

import (
	"bufio"
	"fmt"
	"os"
)

func AccessInventory(player *Character) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n\t=== INVENTAIRE ===")
		fmt.Printf("\tCapacité: %d/%d\n", len(player.Inventory), player.MaxInventory)
		fmt.Printf("\tOr: %d pièces\n", player.Gold)

		if len(player.Inventory) == 0 {
			fmt.Println("\tVotre inventaire est vide.")
		} else {
			for i, item := range player.Inventory {
				fmt.Printf("\t%d. %s", i+1, item.Name)
				if item.Type == "potion" {
					fmt.Printf(" (Restaure %d PV)", item.Value)
				} else if item.Type == "mana_potion" {
					fmt.Printf(" (Restaure %d PM)", item.Value)
				} else if item.Type == "weapon" {
					fmt.Printf(" (Dégâts: %d)", item.Value)
				} else if item.Type == "material" {
					fmt.Printf(" (Matériau)")
				} else if item.Type == "book" {
					fmt.Printf(" (Livre de sort)")
				}
				fmt.Println()
			}
		}

		fmt.Println("\n\tOptions:")
		fmt.Println("\t1. Utiliser un objet")
		fmt.Println("\t2. Équiper un objet")
		fmt.Println("\t0. Retour au menu principal")

		fmt.Print("\tVotre choix: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			useItem(player)
		case "2":
			EquipItemMenu(player)
		case "0":
			return
		default:
			fmt.Println("\tChoix invalide. Veuillez choisir une option valide.")
		}
	}
}

func useItem(player *Character) {
	if len(player.Inventory) == 0 {
		fmt.Println("\tVotre inventaire est vide.")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\n\tQuel objet voulez-vous utiliser?")
	for i, item := range player.Inventory {
		fmt.Printf("\t%d. %s", i+1, item.Name)
		if item.Type == "potion" {
			fmt.Printf(" (Restaure %d PV)", item.Value)
		} else if item.Type == "mana_potion" {
			fmt.Printf(" (Restaure %d PM)", item.Value)
		} else if item.Type == "book" {
			fmt.Printf(" (Livre de sort: %s)", item.Description)
		}
		fmt.Println()
	}
	fmt.Printf("\t%d. Retour\n", len(player.Inventory)+1)

	fmt.Print("\tVotre choix: ")
	scanner.Scan()
	choice := scanner.Text()

	var index int
	fmt.Sscanf(choice, "%d", &index)

	if index == len(player.Inventory)+1 {
		return
	}

	if index < 1 || index > len(player.Inventory) {
		fmt.Println("\tChoix invalide.")
		return
	}

	item := player.Inventory[index-1]

	switch item.Type {
	case "potion":
		player.CurrentHealth += item.Value
		if player.CurrentHealth > player.MaxHealth {
			player.CurrentHealth = player.MaxHealth
		}
		fmt.Printf("\tVous utilisez une potion et récupérez %d PV. PV actuels: %d/%d\n", item.Value, player.CurrentHealth, player.MaxHealth)
		player.removeItem(item.Name)
	case "mana_potion":
		if player.Class == "Mage" {
			player.CurrentMana += item.Value
			if player.CurrentMana > player.MaxMana {
				player.CurrentMana = player.MaxMana
			}
			fmt.Printf("\tVous utilisez une potion de mana et récupérez %d PM. PM actuels: %d/%d\n", item.Value, player.CurrentMana, player.MaxMana)
			player.removeItem(item.Name)
		} else {
			fmt.Println("\tSeuls les mages peuvent utiliser des potions de mana.")
		}
	case "book":
		player.learnSpell(item.Description)
		player.removeItem(item.Name)
	default:
		fmt.Println("\tCet objet ne peut pas être utilisé directement.")
	}
}

func EquipItemMenu(player *Character) {
	if len(player.Inventory) == 0 {
		fmt.Println("\tVotre inventaire est vide.")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\n\tQuel objet voulez-vous équiper?")
	for i, item := range player.Inventory {
		if item.Type == "weapon" || item.Type == "armor" {
			fmt.Printf("\t%d. %s", i+1, item.Name)
			if item.Type == "weapon" {
				fmt.Printf(" (Dégâts: %d)", item.Value)
			} else if item.Type == "armor" {
				fmt.Printf(" (Bonus: +%d PV)", item.Value)
			}
			fmt.Println()
		}
	}
	fmt.Printf("\t%d. Retour\n", len(player.Inventory)+1)

	fmt.Print("\tVotre choix: ")
	scanner.Scan()
	choice := scanner.Text()

	var index int
	fmt.Sscanf(choice, "%d", &index)

	if index == len(player.Inventory)+1 {
		return
	}

	if index < 1 || index > len(player.Inventory) {
		fmt.Println("\tChoix invalide.")
		return
	}

	item := player.Inventory[index-1]

	if item.Type == "weapon" || item.Type == "armor" {
		player.equipItem(item.Name)
	} else {
		fmt.Println("\tCet objet ne peut pas être équipé.")
	}
}

func openChest(player *Character) {
	fmt.Println("\tVous ouvrez le coffre et y trouvez un équipement!")

	var item Item

	switch player.Class {
	case "Humain":
		if player.hasItem("Gants de combat") {
			fmt.Println("\tVous avez déjà des gants de combat. Vous recevez de l'or à la place.")
			player.Gold += 100
		} else {
			item = Item{Name: "Gants de combat", Description: "Augmente les dégâts de mêlée", Type: "weapon", Value: 80, Price: 150}
			if player.addItem(item) {
				fmt.Println("\tVous obtenez des Gants de combat!")
			}
		}
	case "Mage":
		if player.hasItem("Bâton magique") {
			fmt.Println("\tVous avez déjà un bâton magique. Vous recevez de l'or à la place.")
			player.Gold += 120
		} else {
			item = Item{Name: "Bâton magique", Description: "Augmente la puissance magique", Type: "weapon", Value: 60, Price: 180}
			if player.addItem(item) {
				fmt.Println("\tVous obtenez un Bâton magique!")
			}
		}
	case "Samouraï":
		if player.hasItem("Katana") {
			fmt.Println("\tVous avez déjà un katana. Vous recevez de l'or à la place.")
			player.Gold += 200
		} else {
			item = Item{Name: "Katana", Description: "Arme tranchante de samouraï", Type: "weapon", Value: 100, Price: 250}
			if player.addItem(item) {
				fmt.Println("\tVous obtenez un Katana!")
			}
		}
	}
	fmt.Println("\tAppuyez sur Entrée pour continuer...")
	fmt.Scanln()
}
