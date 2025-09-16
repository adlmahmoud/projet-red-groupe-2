package main

import "fmt"

var items []string

func ajoutItem(item string) {
	if len(items) < 10 {
		items = append(items, item)
		fmt.Println("Item ajouté:", item)
	} else {
		fmt.Println("Nombre d'items autorisé dépassé, impossible d'ajouter:", item)
	}
}

func afficherInventaire() {
	fmt.Println("=== Inventaire ===")
	for index, valeur := range items {
		fmt.Printf("%d: %s\n", index+1, valeur)
	}
	fmt.Println()
}

func main() {
	ajoutItem("épée")
	ajoutItem("bouclier")
	// Ajouter plus d'items pour tester la limite
	for i := 0; i < 12; i++ {
		ajoutItem(fmt.Sprintf("potion%d", i+1))
	}
	afficherInventaire()
}
