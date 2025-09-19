package main

import (
	"bufio"
	"fmt"
	"os"
)

func AccessInventory(player *Character) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n=== INVENTAIRE ===")
		fmt.Printf("Capacité: %d/%d\n", len(player.Inventory), player.MaxInventory)
		fmt.Printf("Or: %d pièces\n", player.Gold)

		if len(player.Inventory) == 0 {
			fmt.Println("Votre inventaire est vide.")
		} else {
			for i, item := range player.Inventory {
				fmt.Printf("%d. %s", i+1, item.Name)
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

		fmt.Println("\nOptions:")
		fmt.Println("1. Utiliser un objet")
		fmt.Println("2. Équiper un objet")
		fmt.Println("0. Retour au menu principal")

		fmt.Print("Votre choix: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			UseItem(player)
		case "2":
			EquipItemMenu(player)
		case "0":
			return
		default:
			fmt.Println("Choix invalide. Veuillez choisir une option valide.")
		}
	}
}

func UseItem(player *Character) {
	if len(player.Inventory) == 0 {
		fmt.Println("Votre inventaire est vide.")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\nQuel objet voulez-vous utiliser?")
	for i, item := range player.Inventory {
		fmt.Printf("%d. %s", i+1, item.Name)
		if item.Type == "potion" {
			fmt.Printf(" (Restaure %d PV)", item.Value)
		} else if item.Type == "mana_potion" {
			fmt.Printf(" (Restaure %d PM)", item.Value)
		} else if item.Type == "book" {
			fmt.Printf(" (Livre de sort: %s)", item.Description)
		}
		fmt.Println()
	}
	fmt.Printf("%d. Retour\n", len(player.Inventory)+1)

	fmt.Print("Votre choix: ")
	scanner.Scan()
	choice := scanner.Text()

	var index int
	fmt.Sscanf(choice, "%d", &index)

	if index == len(player.Inventory)+1 {
		return
	}

	if index < 1 || index > len(player.Inventory) {
		fmt.Println("Choix invalide.")
		return
	}

	item := player.Inventory[index-1]

	switch item.Type {
	case "potion":
		player.CurrentHealth += item.Value
		if player.CurrentHealth > player.MaxHealth {
			player.CurrentHealth = player.MaxHealth
		}
		fmt.Printf("Vous utilisez une potion et récupérez %d PV. PV actuels: %d/%d\n", item.Value, player.CurrentHealth, player.MaxHealth)
		player.RemoveItem(item.Name)
	case "mana_potion":
		if player.Class == "Mage" {
			player.CurrentMana += item.Value
			if player.CurrentMana > player.MaxMana {
				player.CurrentMana = player.MaxMana
			}
			fmt.Printf("Vous utilisez une potion de mana et récupérez %d PM. PM actuels: %d/%d\n", item.Value, player.CurrentMana, player.MaxMana)
			player.RemoveItem(item.Name)
		} else {
			fmt.Println("Seuls les mages peuvent utiliser des potions de mana.")
		}
	case "book":
		player.LearnSpell(item.Description)
		player.RemoveItem(item.Name)
	default:
		fmt.Println("Cet objet ne peut pas être utilisé directement.")
	}
}

func EquipItemMenu(player *Character) {
	if len(player.Inventory) == 0 {
		fmt.Println("Votre inventaire est vide.")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\nQuel objet voulez-vous équiper?")
	for i, item := range player.Inventory {
		if item.Type == "weapon" || item.Type == "armor" {
			fmt.Printf("%d. %s", i+1, item.Name)
			if item.Type == "weapon" {
				fmt.Printf(" (Dégâts: %d)", item.Value)
			} else if item.Type == "armor" {
				fmt.Printf(" (Bonus: +%d PV)", item.Value)
			}
			fmt.Println()
		}
	}
	fmt.Printf("%d. Retour\n", len(player.Inventory)+1)

	fmt.Print("Votre choix: ")
	scanner.Scan()
	choice := scanner.Text()

	var index int
	fmt.Sscanf(choice, "%d", &index)

	if index == len(player.Inventory)+1 {
		return
	}

	if index < 1 || index > len(player.Inventory) {
		fmt.Println("Choix invalide.")
		return
	}

	item := player.Inventory[index-1]

	if item.Type == "weapon" || item.Type == "armor" {
		player.EquipItem(item.Name)
	} else {
		fmt.Println("Cet objet ne peut pas être équipé.")
	}
}

func OpenChest(player *Character) {
	fmt.Println("Vous ouvrez le coffre et y trouvez un équipement!")

	var item Item

	switch player.Class {
	case "Humain":
		if player.HasItem("Gants de combat") {
			fmt.Println("Vous avez déjà des gants de combat. Vous recevez de l'or à la place.")
			player.Gold += 40
		} else {
			item = Item{Name: "Gants de combat", Description: "Augmente les dégâts de mêlée", Type: "weapon", Value: 30, Price: 40}
			if player.AddItem(item) {
				fmt.Println("Vous obtenez des Gants de combat!")
			}
		}
	case "Mage":
		if player.HasItem("Bâton magique") {
			fmt.Println("Vous avez déjà un bâton magique. Vous recevez de l'or à la place.")
			player.Gold += 50
		} else {
			item = Item{Name: "Bâton magique", Description: "Augmente la puissance magique", Type: "weapon", Value: 30, Price: 50}
			if player.AddItem(item) {
				fmt.Println("Vous obtenez un Bâton magique!")
			}
		}
	case "Samouraï":
		if player.HasItem("Katana") {
			fmt.Println("Vous avez déjà un katana. Vous recevez de l'or à la place.")
			player.Gold += 60
		} else {
			item = Item{Name: "Katana", Description: "Arme tranchante de samouraï", Type: "weapon", Value: 45, Price: 60}
			if player.AddItem(item) {
				fmt.Println("Vous obtenez un Katana!")
			}
		}
	}
	fmt.Println("Appuyez sur Entrée pour continuer...")
	fmt.Scanln()
}
