package main

import "fmt"

type Character struct {
	Nom        string
	Classe     string
	PVMax      int
	Pv         int
	ManaMax    int
	ManaActuel int
	Attaque1   int
	Attaque2   int
	Inventaire []Item
	Niveau     int
	Or         int
}
type Item struct {
	Nom      string
	Quantite int
}

func InitCharacter() Character {
	var player Character
	fmt.Print("Donner un nom au personnage : ")
	fmt.Scanln(&player.Nom)

	fmt.Println("\tChoisissez une classe :")
	fmt.Println("\t1 - Mage")
	fmt.Println("\t2 - Samouraï")
	fmt.Println("\t3 - Humain")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		player.Classe = "Mage"
		player.PVMax = 250
		player.Pv = 250
		player.ManaMax = 200
		player.ManaActuel = 200
		player.Attaque1 = 12
		player.Inventaire = []Item{
			{"Potion", 2},
			{"Potion de Mana", 1},
			{"Bâton magique", 1},
		}
	case 2:
		player.Classe = "Samouraï"
		player.PVMax = 300
		player.Pv = 300
		player.ManaMax = 50
		player.ManaActuel = 50
		player.Attaque1 = 20
		player.Inventaire = []Item{
			{"Potion", 2},
			{"Katana", 1},
		}
	case 3:
		player.Classe = "Humain"
		player.PVMax = 200
		player.Pv = 200
		player.ManaMax = 100
		player.ManaActuel = 100
		player.Attaque1 = 15
		player.Inventaire = []Item{
			{"Potion", 2},
		}
	default:
		fmt.Println("Classe invalide, Humain par défaut")
		player.Classe = "Humain"
		player.PVMax = 200
		player.Pv = 200
		player.ManaMax = 100
		player.ManaActuel = 100
		player.Attaque1 = 15
		player.Inventaire = []Item{
			{"Potion", 2},
		}
	}
	player.Niveau = 1
	player.Or = 100
	return player
}
func NewCharacter(nom, classe string) Character {
	switch classe {
	case "Mage":
		return Character{
			Nom:        nom,
			Classe:     "Mage",
			PVMax:      250,
			Pv:         50,
			ManaMax:    200,
			ManaActuel: 200,
			Attaque1:   12,
			Inventaire: []Item{
				{"Potion", 2},
				{"Potion de Mana", 1},
				{"Bâton magique", 1},
			},
			Niveau: 1,
			Or:     100,
		}

	case "Samouraï":
		return Character{
			Nom:      nom,
			Classe:   "Samouraï",
			PVMax:    300,
			Pv:       300,
			Attaque1: 45,
			Attaque2: 25,
			Inventaire: []Item{
				{"Potion", 1},
				{"Katana", 1},
			},
			Niveau: 1,
			Or:     100,
		}

	case "Humain":
		return Character{
			Nom:      nom,
			Classe:   "Humain",
			PVMax:    200,
			Pv:       200,
			Attaque1: 20,
			Attaque2: 30,
			Inventaire: []Item{
				{"Potion", 2},
			},
			Niveau: 1,
			Or:     30,
		}
	default:
		return Character{}
	}
}
