package main

import (
	"bufio"
	"fmt"
	"os"
)

func meetMerchant(player *Character) {
	fmt.Println("\n\tVous continuez sur le chemin menant aux portes du royaume. Au loin, une silhouette se dessine.")
	fmt.Println("\tEn vous rapprochant, vous découvrez un vieil homme. Il vous interpelle:")
	fmt.Println("\t<< Je suis marchand d'objets, de ressources, mais aussi d'informations... contre quelques pièces d'or.>>")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n\tQue faites-vous?")
		fmt.Println("\t1. Faire quelques courses")
		fmt.Println("\t2. Pêche aux informations")
		fmt.Println("\t0. Passer votre chemin et continuer")

		fmt.Print("\tVotre choix: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			VisitShop(player)
		case "2":
			BuyInformation(player)
		case "0":
			fmt.Println("\tVous décidez de passer votre chemin et de continuer votre route.")
			return
		default:
			fmt.Println("\tChoix invalide. Veuillez choisir 1, 2 ou 0.")
		}
	}
}

func VisitShop(player *Character) {
	scanner := bufio.NewScanner(os.Stdin)

	shopItems := []Item{
		{Name: "Potion de vie", Description: "Restaure 100 PV", Type: "potion", Value: 100, Price: 20},
		{Name: "Potion de mana", Description: "Restaure 100 PM", Type: "mana_potion", Value: 100, Price: 40},
		{Name: "Potion de poison", Description: "Inflige 20 dégâts par seconde pendant 3s", Type: "poison_potion", Value: 20, Price: 30},
		{Name: "Fourrure de Loup", Description: "Matériau de craft", Type: "material", Value: 0, Price: 8},
		{Name: "Peau de Troll", Description: "Matériau de craft", Type: "material", Value: 0, Price: 15},
		{Name: "Cuir de Sanglier", Description: "Matériau de craft", Type: "material", Value: 0, Price: 6},
		{Name: "Plume de Corbeau", Description: "Matériau de craft", Type: "material", Value: 0, Price: 3},
		{Name: "Épée", Description: "Arme de base", Type: "weapon", Value: 50, Price: 80},
		{Name: "Lance", Description: "Arme longue", Type: "weapon", Value: 80, Price: 150},
		{Name: "Hache de guerre", Description: "Arme lourde", Type: "weapon", Value: 120, Price: 250},
		{Name: "Livre de sort: Éboulement", Description: "Éboulement", Type: "book", Value: 0, Price: 100},
		{Name: "Livre de sort: Éclair pourfendeur", Description: "Éclair pourfendeur", Type: "book", Value: 0, Price: 200},
		{Name: "Livre de sort: Horde de corbeaux", Description: "Horde de corbeaux", Type: "book", Value: 0, Price: 350},
		{Name: "Sabre Kitetsu", Description: "Sabre tranchant", Type: "weapon", Value: 70, Price: 120},
		{Name: "Katana Enma", Description: "Katana légendaire", Type: "weapon", Value: 130, Price: 300},
		{Name: "Sabre Shisui", Description: "Sabre de maître", Type: "weapon", Value: 200, Price: 500},
		{Name: "Augmentation d'inventaire", Description: "+10 emplacements d'inventaire", Type: "upgrade", Value: 10, Price: 100},
	}

	for {
		fmt.Println("\n\t=== BOUTIQUE DU MARCHAND ===")
		fmt.Printf("\tVotre or: %d pièces\n", player.Gold)

		for i, item := range shopItems {
			fmt.Printf("\t%d. %s - %d pièces", i+1, item.Name, item.Price)
			if item.Type == "potion" {
				fmt.Printf(" (Restaure %d PV)", item.Value)
			} else if item.Type == "mana_potion" {
				fmt.Printf(" (Restaure %d PM)", item.Value)
			} else if item.Type == "weapon" {
				fmt.Printf(" (Dégâts: %d)", item.Value)
			} else if item.Type == "book" {
				fmt.Printf(" (Apprend le sort: %s)", item.Description)
			}
			fmt.Println()
		}
		fmt.Printf("\t0. Quitter la boutique\n")

		fmt.Print("\tVotre choix: ")
		scanner.Scan()
		choice := scanner.Text()

		if choice == "0" {
			fmt.Println("\tMerci pour votre visite!")
			return
		}

		var index int
		fmt.Sscanf(choice, "%d", &index)

		if index < 1 || index > len(shopItems) {
			fmt.Println("\tChoix invalide.")
			continue
		}

		selectedItem := shopItems[index-1]

		if player.Gold < selectedItem.Price {
			fmt.Println("\tVous n'avez pas assez d'or pour acheter cet objet.")
			continue
		}

		if !player.addItem(selectedItem) {
			fmt.Println("\tVotre inventaire est plein. Vous ne pouvez pas acheter cet objet.")
			continue
		}

		player.Gold -= selectedItem.Price
		fmt.Printf("\tVous achetez %s pour %d pièces d'or.\n", selectedItem.Name, selectedItem.Price)

		if selectedItem.Type == "upgrade" {
			player.upgradeInventory()
		}
	}
}

func BuyInformation(player *Character) {
	scanner := bufio.NewScanner(os.Stdin)

	informations := []struct {
		name        string
		description string
		price       int
	}{
		{"Chavrot", "Avant que Chavrot ne devienne le gardien des portes du royaume d'Ynov, il était connu sous un autre nom : Morgann, un enfant né sous une lune noire...", 25},
		{"Moonlight", "Autrefois, une prêtresse nommée Lunara chantait pour les étoiles. Ses mélodies guidaient les marins perdus, apaisaient les bêtes lunaires...", 35},
		{"Ainzolgone", "Autrefois, Ainzolgone était un mage humain nommé Satorin, membre d'une guilde légendaire dans le monde...", 50},
		{"Bahamut", "Avant que Nazarick ne devienne une guilde en ruine, il existait un gouffre oublié, scellé par trois chaînes d'or...", 75},
		{"Negal", "Il fut un temps où les rois ne régnaient pas par droit divin, mais par la force brute. Negal était un guerrier né dans les fosses de l'Altus Profond...", 100},
		{"Achlys", "Il fut un temps où le royaume d'Ynov brillait sous la lumière des lunes jumelles. Le roi Thalor, sage et aimé, régnait avec justice...", 150},
	}

	for {
		fmt.Println("\n\t=== INFORMATIONS DISPONIBLES ===")
		fmt.Printf("\tVotre or: %d pièces\n", player.Gold)

		for i, info := range informations {
			fmt.Printf("\t%d. Information sur %s - %d pièces\n", i+1, info.name, info.price)
		}
		fmt.Printf("\t0. Quitter\n")

		fmt.Print("\tVotre choix: ")
		scanner.Scan()
		choice := scanner.Text()

		if choice == "0" {
			return
		}

		var index int
		fmt.Sscanf(choice, "%d", &index)

		if index < 1 || index > len(informations) {
			fmt.Println("\tChoix invalide.")
			continue
		}

		selectedInfo := informations[index-1]

		if player.Gold < selectedInfo.price {
			fmt.Println("\tVous n'avez pas assez d'or pour acheter cette information.")
			continue
		}

		player.Gold -= selectedInfo.price
		fmt.Printf("\tVous achetez l'information sur %s pour %d pièces d'or.\n", selectedInfo.name, selectedInfo.price)
		fmt.Printf("\tInformation: %s\n", selectedInfo.description)
		fmt.Println("\tAppuyez sur Entrée pour continuer...")
		fmt.Scanln()
	}
}

func VisitBlacksmith(player *Character) {
	scanner := bufio.NewScanner(os.Stdin)

	recipes := []struct {
		name        string
		description string
		price       int
		materials   map[string]int
		item        Item
	}{
		{
			"Chapeau de l'aventurier",
			"Bonus: +20 PV max",
			25,
			map[string]int{"Plume de Corbeau": 2, "Cuir de Sanglier": 3},
			Item{Name: "Chapeau de l'aventurier", Description: "Bonus de vie", Type: "armor", Value: 20, Price: 100},
		},
		{
			"Tunique de l'aventurier",
			"Bonus: +50 PV max",
			50,
			map[string]int{"Fourrure de Loup": 4, "Peau de Troll": 2},
			Item{Name: "Tunique de l'aventurier", Description: "Bonus de vie", Type: "armor", Value: 50, Price: 200},
		},
		{
			"Bottes de l'aventurier",
			"Bonus: +30 PV max",
			35,
			map[string]int{"Fourrure de Loup": 2, "Cuir de Sanglier": 3},
			Item{Name: "Bottes de l'aventurier", Description: "Bonus de vie", Type: "armor", Value: 30, Price: 150},
		},
		{
			"Épée améliorée",
			"Dégâts: 100",
			75,
			map[string]int{"Cuir de Sanglier": 4, "Fourrure de Loup": 3},
			Item{Name: "Épée améliorée", Description: "Arme améliorée", Type: "weapon", Value: 100, Price: 300},
		},
	}

	for {
		fmt.Println("\n\t=== ATELIER DU FORGERON ===")
		fmt.Printf("\tVotre or: %d pièces\n", player.Gold)

		for i, recipe := range recipes {
			fmt.Printf("\t%d. %s - %d pièces", i+1, recipe.name, recipe.price)
			fmt.Printf(" (Matériaux: ")
			for material, quantity := range recipe.materials {
				fmt.Printf("%d %s ", quantity, material)
			}
			fmt.Printf(")\n")
		}
		fmt.Printf("\t0. Retour\n")

		fmt.Print("\tVotre choix: ")
		scanner.Scan()
		choice := scanner.Text()

		if choice == "0" {
			return
		}

		var index int
		fmt.Sscanf(choice, "%d", &index)

		if index < 1 || index > len(recipes) {
			fmt.Println("\tChoix invalide.")
			continue
		}

		selectedRecipe := recipes[index-1]

		if player.Gold < selectedRecipe.price {
			fmt.Println("\tVous n'avez pas assez d'or pour fabriquer cet objet.")
			continue
		}

		hasMaterials := true
		for material, quantity := range selectedRecipe.materials {
			if player.countItem(material) < quantity {
				fmt.Printf("\tVous n'avez pas assez de %s (nécessaire: %d).\n", material, quantity)
				hasMaterials = false
				break
			}
		}

		if !hasMaterials {
			continue
		}

		if !player.addItem(selectedRecipe.item) {
			fmt.Println("\tVotre inventaire est plein. Vous ne pouvez pas fabriquer cet objet.")
			continue
		}

		player.Gold -= selectedRecipe.price
		for material, quantity := range selectedRecipe.materials {
			for i := 0; i < quantity; i++ {
				player.removeItem(material)
			}
		}

		fmt.Printf("\tVous fabriquez %s pour %d pièces d'or!\n", selectedRecipe.name, selectedRecipe.price)
	}
}
