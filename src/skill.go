package main

import (
	"fmt"
)

<<<<<<< HEAD
type Character1 struct {
	Name       string
	Skills     []string
	Inventaire []Item
}

func InitCharacter1(name string) *Character {
=======
type Character struct {
	name      string
	skills    []string
	inventory []string
}

// Initialise un personnage avec ses compétences de base
func initCharacter(name string) *Character {
>>>>>>> 25c08f92a67127ffc558f97092a48889d967b1a2
	var baseSkills []string

	switch name {
	case "Humain":
		baseSkills = []string{"Coup de poing", "Coup de pieds"}
	case "Mage":
		baseSkills = []string{"Luciole de feu", "Boule de feu"}
	case "Samouraï":
		baseSkills = []string{"Coup de poing", "Coup de pieds"}
	default:
		baseSkills = []string{"Aucun skill"}
	}

	return &Character{
<<<<<<< HEAD
		Nom:        name,
		skills:     baseSkills,
		Inventaire: p.Inventaire,
	}
}

func (c *Character) SpellBook(spell string) {
	for _, s := range c.skills {
		if s == spell {
			fmt.Printf("%s connaît déjà le sort '%s'.\n", c.Nom, spell)
=======
		name:      name,
		skills:    baseSkills,
		inventory: []string{},
	}
}

func (c *Character) spellBook(spell string) {
	for _, s := range c.skills {
		if s == spell {
			fmt.Printf("%s connaît déjà le sort '%s'.\n", c.name, spell)
>>>>>>> 25c08f92a67127ffc558f97092a48889d967b1a2
			return
		}
	}
	c.skills = append(c.skills, spell)
<<<<<<< HEAD
	fmt.Printf("%s a appris un nouveau sort : '%s'.\n", c.Nom, spell)
}
func (c *Character) AddItem(item string) {
	if len(c.Inventaire) >= 10 {
		fmt.Printf("Inventaire plein pour %s, impossible d'ajouter : %s\n", c.Nom, item)
		return
	}
	c.Inventaire = append(c.Inventaire, Item{})
	fmt.Printf("Item ajouté à l'inventaire de %s : %s\n", c.Nom, item)
}

func (c *Character) UseItem(item string) {
	foundIndex := -1
	for i, it := range c.Inventaire {
		if it.Nom == item {
			foundIndex = i
			break
		}

		if foundIndex == -1 {
			fmt.Printf("Item '%s' non trouvé dans l'inventaire de %s.\n", item, c.Nom)
			return
		}

		if item == "Livre de Sort : Boule de Feu" {
			c.SpellBook()("Boule de feu")
			c.Inventaire = append(c.Inventaire[:foundIndex], c.Inventaire[foundIndex+1:]...)
			fmt.Printf("%s a utilisé le '%s' et l'a retiré de son inventaire.\n", c.Nom, item)
		} else {
			fmt.Printf("%s utilise l'item '%s'.\n", c.Nom, item)
		}
	}
}

func (c *Character) AfficherInventaire() {
	fmt.Printf("=== Inventaire de %s ===\n", c.Nom)
	for i, item := range c.Inventaire {
		fmt.Printf("%d: %s\n", i+1, item)
	}
}
func (c *Character) AfficherSkills() {
	fmt.Printf("=== Sorts appris par %s ===\n", c.Nom)
=======
	fmt.Printf("%s a appris un nouveau sort : '%s'.\n", c.name, spell)
}
func (c *Character) addItem(item string) {
	if len(c.inventory) >= 10 {
		fmt.Printf("Inventaire plein pour %s, impossible d'ajouter : %s\n", c.name, item)
		return
	}
	c.inventory = append(c.inventory, item)
	fmt.Printf("Item ajouté à l'inventaire de %s : %s\n", c.name, item)
}

func (c *Character) useItem(item string) {
	foundIndex := -1
	for i, it := range c.inventory {
		if it == item {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		fmt.Printf("Item '%s' non trouvé dans l'inventaire de %s.\n", item, c.name)
		return
	}

	if item == "Livre de Sort : Boule de Feu" {
		c.spellBook("Boule de feu")
		c.inventory = append(c.inventory[:foundIndex], c.inventory[foundIndex+1:]...)
		fmt.Printf("%s a utilisé le '%s' et l'a retiré de son inventaire.\n", c.name, item)
	} else {
		fmt.Printf("%s utilise l'item '%s'.\n", c.name, item)
	}
}

func (c *Character) afficherInventaire() {
	fmt.Printf("=== Inventaire de %s ===\n", c.name)
	for i, item := range c.inventory {
		fmt.Printf("%d: %s\n", i+1, item)
	}
}
func (c *Character) afficherSkills() {
	fmt.Printf("=== Sorts appris par %s ===\n", c.name)
>>>>>>> 25c08f92a67127ffc558f97092a48889d967b1a2
	for i, skill := range c.skills {
		fmt.Printf("%d: %s\n", i+1, skill)
	}
	fmt.Println()
}
