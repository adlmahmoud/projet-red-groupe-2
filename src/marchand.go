package main

import "fmt"

func (player *Character) Merchant() {
	fmt.Println("\n=== Marchand ===")
	fmt.Println("Bonjour voyageur ! La première potion est gratuite pour vous aider dans votre aventure.")
	player.Inventaire = append(player.Inventaire, Item{"Potion", 1})
	fmt.Println("Vous avez reçu une potion de soin gratuitement !")

	for {
		fmt.Println("Que voulez-vous faire ?")
		fmt.Println("1 - Acheter une potion de soin (15 or)")
		if player.Classe == "Mage" {
			fmt.Println("2 - Acheter une potion de mana (20 or)")
			fmt.Println("3 - Acheter un nouveau sort (50 or)")
		}
		fmt.Println("0 - Quitter le marchand")

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			if player.Or >= 15 {
				player.Or -= 15
				player.Inventaire = append(player.Inventaire, Item{"Potion", 1})
				fmt.Println("Vous avez acheté une potion de soin !")
			} else {
				fmt.Println("Pas assez d'or pour acheter cette potion.")
			}

		case 2:
			if player.Classe != "Mage" {
				fmt.Println("Option invalide.")
				continue
			}
			if player.Or >= 20 {
				player.Or -= 20
				player.Inventaire = append(player.Inventaire, Item{"Potion de Mana", 1})
				fmt.Println("Vous avez acheté une potion de mana !")
			} else {
				fmt.Println("Pas assez d'or pour acheter cette potion de mana.")
			}

		case 3:
			if player.Classe != "Mage" {
				fmt.Println("Option invalide.")
				continue
			}
			if player.Or >= 50 {
				player.Or -= 50
				fmt.Println("Vous avez appris un nouveau sort puissant !")
				player.Attaque1 += 10
			} else {
				fmt.Println("Pas assez d'or pour acheter ce sort.")
			}

		case 0:
			fmt.Println("Merci et bonne aventure !")
			return

		default:
			fmt.Println("Choix invalide, réessayez.")
		}
		fmt.Printf("Or restant : %d\n\n", player.Or)
	}
}
