package main

import "fmt"

type Character struct {
	Nom        string
	Classe     string
	Niveau     int
	Pv         int
	PvMax      int
	Mana       int
	Équipement string
	Inventaire []Item
}

type Item struct {
	Nom      string
	Quantite int
}

type Equipement struct {
	Nom      string
	Parti    string
	Quantite int
}

func (player *Character) InitCharacterHumain() {
	*player = Character{
		Nom:    "Chris",
		Classe: "Humain",
		Niveau: 1,
		Pv:     30,
		PvMax:  200,
		Inventaire: []Item{
			{"Potion de vie", 2},
			{"Pièces d'or", 30},
		},
	}
}

func (player *Character) InitCharacterMage() {
	*player = Character{
		Nom:    "Celeste",
		Classe: "Mage",
		Niveau: 1,
		Pv:     35,
		PvMax:  250,
		Mana:   30,
		Inventaire: []Item{
			{"Potion de vie", 2},
			{"Potion de mana", 1},
			{"Pièces d'or", 20},
		},
	}
}

func (player *Character) InitCharacterSamouraï() {
	*player = Character{
		Nom:    "Genzo",
		Classe: "Samouraï",
		Niveau: 1,
		Pv:     40,
		PvMax:  300,
		Inventaire: []Item{
			{"Potion de vie", 1},
			{"Pièces d'or", 20},
		},
	}
}

func (equipement *Équipement) InitEquipement() {
	*equipement = Équipement{""}

}

func main() {
	var players [3]Character

	// Initialisation de chaque personnage
	players[0].InitCharacterHumain()
	players[1].InitCharacterMage()
	players[2].InitCharacterSamouraï()

	// Affichage des informations des personnages
	fmt.Println("Personnage 1:", players[0])
	fmt.Println("Personnage 2:", players[1])
	fmt.Println("Personnage 3:", players[2])
}
