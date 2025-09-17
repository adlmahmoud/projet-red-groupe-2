package main

type Character struct {
	Nom        string
	Classe     string
	Niveau     int
	Pv         int
	PvMax      int
	Inventaire []Item
}

type Item struct {
	Nom      string
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

/* func (player *Character) InitCharacter() {
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

func (player *Character) InitCharacter() {
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
} */
