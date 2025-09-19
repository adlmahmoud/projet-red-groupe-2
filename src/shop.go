package main

import (
	"bufio"
	"fmt"
	"os"
)

func MeetMerchant(player *Character) {
	fmt.Println("\nVous continuez sur le chemin menant aux portes du royaume. Au loin, une silhouette se dessine.")
	fmt.Println("En vous rapprochant, vous découvrez un vieil homme. Il vous interpelle:")
	fmt.Println("<< Je suis marchand d'objets, de ressources, mais aussi d'informations... contre quelques pièces d'or.>>")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nQue faites-vous?")
		fmt.Println("1. Faire quelques courses")
		fmt.Println("2. Pêche aux informations")
		fmt.Println("0. Passer votre chemin et continuer")

		fmt.Print("Votre choix: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			VisitShop(player)
		case "2":
			BuyInformation(player)
		case "0":
			fmt.Println("Vous décidez de passer votre chemin et de continuer votre route.")
			return
		default:
			fmt.Println("Choix invalide. Veuillez choisir 1, 2 ou 0.")
		}
	}
}

func VisitShop(player *Character) {
	scanner := bufio.NewScanner(os.Stdin)

	shopItems := []Item{
		{Name: "Potion de vie", Description: "Restaure 50 PV", Type: "potion", Value: 50, Price: 10},
		{Name: "Potion de mana", Description: "Restaure 50 PM", Type: "mana_potion", Value: 50, Price: 20},
		{Name: "Potion de poison", Description: "Inflige 10 dégâts par seconde pendant 3s", Type: "poison_potion", Value: 10, Price: 15},
		{Name: "Fourrure de Loup", Description: "Matériau de craft", Type: "material", Value: 0, Price: 4},
		{Name: "Peau de Troll", Description: "Matériau de craft", Type: "material", Value: 0, Price: 7},
		{Name: "Cuir de Sanglier", Description: "Matériau de craft", Type: "material", Value: 0, Price: 3},
		{Name: "Plume de Corbeau", Description: "Matériau de craft", Type: "material", Value: 0, Price: 1},
		{Name: "Épée", Description: "Arme de base", Type: "weapon", Value: 40, Price: 30},
		{Name: "Lance", Description: "Arme longue", Type: "weapon", Value: 70, Price: 80},
		{Name: "Hache de guerre", Description: "Arme lourde", Type: "weapon", Value: 120, Price: 150},
		{Name: "Livre de sort: Éboulement", Description: "Éboulement", Type: "book", Value: 0, Price: 35},
		{Name: "Livre de sort: Éclair pourfendeur", Description: "Éclair pourfendeur", Type: "book", Value: 0, Price: 85},
		{Name: "Livre de sort: Horde de corbeaux", Description: "Horde de corbeaux", Type: "book", Value: 0, Price: 160},
		{Name: "Sabre Kitetsu", Description: "Sabre tranchant", Type: "weapon", Value: 50, Price: 35},
		{Name: "Katana Enma", Description: "Katana légendaire", Type: "weapon", Value: 85, Price: 90},
		{Name: "Sabre Shisui", Description: "Sabre de maître", Type: "weapon", Value: 150, Price: 170},
		{Name: "Augmentation d'inventaire", Description: "+10 emplacements d'inventaire", Type: "upgrade", Value: 10, Price: 30},
	}

	for {
		fmt.Println("\n=== BOUTIQUE DU MARCHAND ===")
		fmt.Printf("Votre or: %d pièces\n", player.Gold)

		for i, item := range shopItems {
			fmt.Printf("%d. %s - %d pièces", i+1, item.Name, item.Price)
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
		fmt.Printf("0. Quitter la boutique\n")

		fmt.Print("Votre choix: ")
		scanner.Scan()
		choice := scanner.Text()

		if choice == "0" {
			fmt.Println("Merci pour votre visite!")
			return
		}

		var index int
		fmt.Sscanf(choice, "%d", &index)

		if index < 1 || index > len(shopItems) {
			fmt.Println("Choix invalide.")
			continue
		}

		selectedItem := shopItems[index-1]

		if player.Gold < selectedItem.Price {
			fmt.Println("Vous n'avez pas assez d'or pour acheter cet objet.")
			continue
		}

		if !player.AddItem(selectedItem) {
			fmt.Println("Votre inventaire est plein. Vous ne pouvez pas acheter cet objet.")
			continue
		}

		player.Gold -= selectedItem.Price
		fmt.Printf("Vous achetez %s pour %d pièces d'or.\n", selectedItem.Name, selectedItem.Price)

		if selectedItem.Type == "upgrade" {
			player.UpgradeInventory()
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
		{"Chavrot", "Avant que Chavrot ne devienne le gardien des portes du royaume d'Ynov, il était connu sous un autre nom : Morgann, un enfant né sous une lune noire...", 10},
		{"Moonlight", "Autrefois, une prêtresse nommée Lunara chantait pour les étoiles. Ses mélodies guidaient les marins perdus, apaisaient les bêtes lunaires...", 15},
		{"Ainzolgone", "Autrefois, Ainzolgone était un mage humain nommé Satorin, membre d'une guilde légendaire dans le monde...", 20},
		{"Bahamut", "Avant que Nazarick ne devienne une guilde en ruine, il existait un gouffre oublié, scellé par trois chaînes d'or...", 25},
		{"Negal", "Il fut un temps où les rois ne régnaient pas par droit divin, mais par la force brute. Negal était un guerrier né dans les fosses de l'Altus Profond...", 30},
		{"Achlys", "Il fut un temps où le royaume d'Ynov brillait sous la lumière des lunes jumelles. Le roi Thalor, sage et aimé, régnait avec justice...", 35},
	}

	for {
		fmt.Println("\n=== INFORMATIONS DISPONIBLES ===")
		fmt.Printf("Votre or: %d pièces\n", player.Gold)

		for i, info := range informations {
			fmt.Printf("%d. Information sur %s - %d pièces\n", i+1, info.name, info.price)
		}
		fmt.Printf("0. Quitter\n")

		fmt.Print("Votre choix: ")
		scanner.Scan()
		choice := scanner.Text()

		if choice == "0" {
			return
		}

		var index int
		fmt.Sscanf(choice, "%d", &index)

		if index < 1 || index > len(informations) {
			fmt.Println("Choix invalide.")
			continue
		}

		selectedInfo := informations[index-1]

		if player.Gold < selectedInfo.price {
			fmt.Println("Vous n'avez pas assez d'or pour acheter cette information.")
			continue
		}

		player.Gold -= selectedInfo.price
		fmt.Printf("Vous achetez l'information sur %s pour %d pièces d'or.\n", selectedInfo.name, selectedInfo.price)
		fmt.Printf("Information: %s\n", selectedInfo.description)
		fmt.Println("Appuyez sur Entrée pour continuer...")
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
			"Bonus: +10 PV max",
			5,
			map[string]int{"Plume de Corbeau": 1, "Cuir de Sanglier": 1},
			Item{Name: "Chapeau de l'aventurier", Description: "Bonus de vie", Type: "armor", Value: 10, Price: 30},
		},
		{
			"Tunique de l'aventurier",
			"Bonus: +25 PV max",
			10,
			map[string]int{"Fourrure de Loup": 2, "Peau de Troll": 1},
			Item{Name: "Tunique de l'aventurier", Description: "Bonus de vie", Type: "armor", Value: 25, Price: 60},
		},
		{
			"Bottes de l'aventurier",
			"Bonus: +15 PV max",
			7,
			map[string]int{"Fourrure de Loup": 1, "Cuir de Sanglier": 1},
			Item{Name: "Bottes de l'aventurier", Description: "Bonus de vie", Type: "armor", Value: 15, Price: 45},
		},
		{
			"Épée améliorée",
			"Dégâts: 60",
			15,
			map[string]int{"Cuir de Sanglier": 2, "Fourrure de Loup": 1},
			Item{Name: "Épée améliorée", Description: "Arme améliorée", Type: "weapon", Value: 60, Price: 70},
		},
	}

	for {
		fmt.Println("\n=== ATELIER DU FORGERON ===")
		fmt.Printf("Votre or: %d pièces\n", player.Gold)

		for i, recipe := range recipes {
			fmt.Printf("%d. %s - %d pièces", i+1, recipe.name, recipe.price)
			fmt.Printf(" (Matériaux: ")
			for material, quantity := range recipe.materials {
				fmt.Printf("%d %s ", quantity, material)
			}
			fmt.Printf(")\n")
		}
		fmt.Printf("0. Retour\n")

		fmt.Print("Votre choix: ")
		scanner.Scan()
		choice := scanner.Text()

		if choice == "0" {
			return
		}

		var index int
		fmt.Sscanf(choice, "%d", &index)

		if index < 1 || index > len(recipes) {
			fmt.Println("Choix invalide.")
			continue
		}

		selectedRecipe := recipes[index-1]

		if player.Gold < selectedRecipe.price {
			fmt.Println("Vous n'avez pas assez d'or pour fabriquer cet objet.")
			continue
		}

		hasMaterials := true
		for material, quantity := range selectedRecipe.materials {
			if player.CountItem(material) < quantity {
				fmt.Printf("Vous n'avez pas assez de %s (nécessaire: %d).\n", material, quantity)
				hasMaterials = false
				break
			}
		}

		if !hasMaterials {
			continue
		}

		if !player.AddItem(selectedRecipe.item) {
			fmt.Println("Votre inventaire est plein. Vous ne pouvez pas fabriquer cet objet.")
			continue
		}

		player.Gold -= selectedRecipe.price
		for material, quantity := range selectedRecipe.materials {
			for i := 0; i < quantity; i++ {
				player.RemoveItem(material)
			}
		}

		fmt.Printf("Vous fabriquez %s pour %d pièces d'or!\n", selectedRecipe.name, selectedRecipe.price)
	}
}
