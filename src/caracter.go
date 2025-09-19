package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Character struct {
	Name          string
	Class         string
	Level         int
	MaxHealth     int
	CurrentHealth int
	MaxMana       int
	CurrentMana   int
	Gold          int
	CurrentXP     int
	RequiredXP    int
	BaseDamage    int
	Inventory     []Item
	Skills        []Skill
	Equipment     Equipment
	MaxInventory  int
	Initiative    int
}

type Equipment struct {
	Head   Item
	Body   Item
	Feet   Item
	Weapon Item
}

type Item struct {
	Name        string
	Description string
	Type        string
	Value       int
	ManaCost    int
	Price       int
}

type Skill struct {
	Name     string
	Damage   int
	ManaCost int
}

func CreateCharacter() *Character {
	scanner := bufio.NewScanner(os.Stdin)
	var player Character

	fmt.Println("\n=== CRÉATION DU PERSONNAGE ===")

	for {
		fmt.Print("Entrez le nom de votre personnage: ")
		scanner.Scan()
		name := strings.TrimSpace(scanner.Text())

		if IsValidName(name) {
			player.Name = FormatName(name)
			break
		} else {
			fmt.Println("Nom invalide. Utilisez uniquement des lettres.")
		}
	}

	for {
		fmt.Println("Choisissez votre classe:")
		fmt.Println("1. Humain")
		fmt.Println("2. Mage")
		fmt.Println("3. Samouraï")
		fmt.Print("Votre choix: ")
		scanner.Scan()
		classChoice := scanner.Text()

		switch classChoice {
		case "1":
			player.Class = "Humain"
			player.MaxHealth = 200
			player.CurrentHealth = 30
			player.Inventory = []Item{
				{Name: "Potion", Description: "Restaure 50 PV", Type: "potion", Value: 50, Price: 10},
				{Name: "Potion", Description: "Restaure 50 PV", Type: "potion", Value: 50, Price: 10},
			}
			player.Skills = []Skill{
				{Name: "Coups de poing", Damage: 20, ManaCost: 0},
				{Name: "Coups de pieds", Damage: 20, ManaCost: 0},
			}
			player.Gold = 30
			player.BaseDamage = 20
			player.Initiative = 5
		case "2":
			player.Class = "Mage"
			player.MaxHealth = 250
			player.CurrentHealth = 35
			player.MaxMana = 200
			player.CurrentMana = 30
			player.Inventory = []Item{
				{Name: "Potion", Description: "Restaure 50 PV", Type: "potion", Value: 50, Price: 10},
				{Name: "Potion", Description: "Restaure 50 PV", Type: "potion", Value: 50, Price: 10},
				{Name: "Potion de mana", Description: "Restaure 50 PM", Type: "mana_potion", Value: 50, Price: 20},
				{Name: "Bâton magique", Description: "Arme de mage", Type: "weapon", Value: 30, Price: 50},
			}
			player.Skills = []Skill{
				{Name: "Luciole de feu", Damage: 25, ManaCost: 5},
				{Name: "Éclair", Damage: 25, ManaCost: 10},
			}
			player.Gold = 20
			player.BaseDamage = 15
			player.Initiative = 3
		case "3":
			player.Class = "Samouraï"
			player.MaxHealth = 300
			player.CurrentHealth = 40
			player.Inventory = []Item{
				{Name: "Potion", Description: "Restaure 50 PV", Type: "potion", Value: 50, Price: 10},
				{Name: "Katana", Description: "Arme de samouraï", Type: "weapon", Value: 45, Price: 60},
			}
			player.Skills = []Skill{
				{Name: "Coups de poing", Damage: 25, ManaCost: 0},
				{Name: "Coups de pieds", Damage: 25, ManaCost: 0},
			}
			player.Gold = 20
			player.BaseDamage = 25
			player.Initiative = 7
		default:
			fmt.Println("Choix invalide. Veuillez choisir 1, 2 ou 3.")
			continue
		}
		break
	}

	player.Level = 1
	player.MaxInventory = 10
	player.CurrentXP = 0
	player.RequiredXP = 100

	fmt.Printf("Félicitations! %s le %s est créé!\n", player.Name, player.Class)
	player.DisplayInfo()

	return &player
}

func IsValidName(name string) bool {
	if len(name) == 0 {
		return false
	}

	for _, char := range name {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || char == ' ') {
			return false
		}
	}

	return true
}

func FormatName(name string) string {
	name = strings.ToLower(name)
	if len(name) > 0 {
		name = strings.ToUpper(string(name[0])) + name[1:]
	}
	return name
}

func (c *Character) DisplayInfo() {
	fmt.Println("\n=== INFORMATIONS DU PERSONNAGE ===")
	fmt.Printf("Nom: %s\n", c.Name)
	fmt.Printf("Classe: %s\n", c.Class)
	fmt.Printf("Niveau: %d\n", c.Level)
	fmt.Printf("PV: %d/%d\n", c.CurrentHealth, c.MaxHealth)
	if c.Class == "Mage" {
		fmt.Printf("PM: %d/%d\n", c.CurrentMana, c.MaxMana)
	}
	fmt.Printf("Or: %d pièces\n", c.Gold)
	fmt.Printf("XP: %d/%d\n", c.CurrentXP, c.RequiredXP)
	fmt.Printf("Dégâts de base: %d\n", c.BaseDamage)
	fmt.Printf("Initiative: %d\n", c.Initiative)

	if len(c.Skills) > 0 {
		fmt.Println("Compétences:")
		for _, skill := range c.Skills {
			fmt.Printf("  - %s (Dégâts: %d, Coût en mana: %d)\n", skill.Name, skill.Damage, skill.ManaCost)
		}
	}

	if c.Equipment.Head.Name != "" || c.Equipment.Body.Name != "" || c.Equipment.Feet.Name != "" || c.Equipment.Weapon.Name != "" {
		fmt.Println("Équipement:")
		if c.Equipment.Head.Name != "" {
			fmt.Printf("  Tête: %s\n", c.Equipment.Head.Name)
		}
		if c.Equipment.Body.Name != "" {
			fmt.Printf("  Corps: %s\n", c.Equipment.Body.Name)
		}
		if c.Equipment.Feet.Name != "" {
			fmt.Printf("  Pieds: %s\n", c.Equipment.Feet.Name)
		}
		if c.Equipment.Weapon.Name != "" {
			fmt.Printf("  Arme: %s\n", c.Equipment.Weapon.Name)
		}
	}
	fmt.Println("Appuyez sur Entrée pour continuer...")
	fmt.Scanln()
}

func (c *Character) UsePotion() {
	for i, item := range c.Inventory {
		if item.Type == "potion" {
			c.CurrentHealth += item.Value
			if c.CurrentHealth > c.MaxHealth {
				c.CurrentHealth = c.MaxHealth
			}
			fmt.Printf("Vous utilisez une potion et récupérez %d PV. PV actuels: %d/%d\n", item.Value, c.CurrentHealth, c.MaxHealth)

			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return
		}
	}
	fmt.Println("Vous n'avez pas de potion de vie dans votre inventaire.")
}

func (c *Character) UseManaPotion() {
	for i, item := range c.Inventory {
		if item.Type == "mana_potion" {
			c.CurrentMana += item.Value
			if c.CurrentMana > c.MaxMana {
				c.CurrentMana = c.MaxMana
			}
			fmt.Printf("Vous utilisez une potion de mana et récupérez %d PM. PM actuels: %d/%d\n", item.Value, c.CurrentMana, c.MaxMana)

			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return
		}
	}
	fmt.Println("Vous n'avez pas de potion de mana dans votre inventaire.")
}

func (c *Character) isDead() bool {
	return c.CurrentHealth <= 0
}

func (c *Character) revive() {
	if c.isDead() {
		c.CurrentHealth = c.MaxHealth / 2
		fmt.Printf("%s est ressuscité avec %d PV!\n", c.Name, c.CurrentHealth)
	}
}

func (c *Character) LearnSpell(spellName string) {
	for _, skill := range c.Skills {
		if skill.Name == spellName {
			fmt.Println("Vous connaissez déjà ce sort!")
			return
		}
	}

	switch spellName {
	case "Boule de feu":
		c.Skills = append(c.Skills, Skill{Name: "Boule de feu", Damage: 30, ManaCost: 10})
	case "Éclair pourfendeur":
		c.Skills = append(c.Skills, Skill{Name: "Éclair pourfendeur", Damage: 30, ManaCost: 10})
	case "Éboulement":
		c.Skills = append(c.Skills, Skill{Name: "Éboulement", Damage: 40, ManaCost: 20})
	case "Horde de corbeaux":
		c.Skills = append(c.Skills, Skill{Name: "Horde de corbeaux", Damage: 130, ManaCost: 50})
	}

	fmt.Printf("Vous avez appris le sort: %s!\n", spellName)
}

func (c *Character) AddItem(item Item) bool {
	if len(c.Inventory) >= c.MaxInventory {
		fmt.Println("Votre inventaire est plein!")
		return false
	}

	c.Inventory = append(c.Inventory, item)
	return true
}

func (c *Character) RemoveItem(itemName string) bool {
	for i, item := range c.Inventory {
		if item.Name == itemName {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return true
		}
	}
	return false
}

func (c *Character) HasItem(itemName string) bool {
	for _, item := range c.Inventory {
		if item.Name == itemName {
			return true
		}
	}
	return false
}

func (c *Character) CountItem(itemName string) int {
	count := 0
	for _, item := range c.Inventory {
		if item.Name == itemName {
			count++
		}
	}
	return count
}

func (c *Character) EquipItem(itemName string) {
	for i, item := range c.Inventory {
		if item.Name == itemName {
			switch item.Type {
			case "weapon":
				if c.Equipment.Weapon.Name != "" {
					c.AddItem(c.Equipment.Weapon)
				}

				c.Equipment.Weapon = item

				if c.Class == "Humain" && item.Name == "Gants de combat" {
					c.Skills = append(c.Skills, Skill{Name: "Patate de forain", Damage: 30, ManaCost: 0})
					fmt.Println("Vous équipez les gants de combat et apprenez 'Patate de forain'!")
				} else if c.Class == "Mage" && item.Name == "Bâton magique" {
					c.Skills = append(c.Skills, Skill{Name: "Boule de feu", Damage: 30, ManaCost: 10})
					fmt.Println("Vous équipez le bâton magique et apprenez 'Boule de feu'!")
				} else if c.Class == "Samouraï" && item.Name == "Katana" {
					c.Skills = append(c.Skills, Skill{Name: "Dans 3 pas tu meurs", Damage: 45, ManaCost: 0})
					fmt.Println("Vous équipez le katana et apprenez 'Dans 3 pas tu meurs'!")
				}

			case "armor":
				switch item.Name {
				case "Chapeau de l'aventurier":
					if c.Equipment.Head.Name != "" {
						c.AddItem(c.Equipment.Head)
						c.MaxHealth -= 10
					}
					c.Equipment.Head = item
					c.MaxHealth += 10
					fmt.Println("Vous équipez le Chapeau de l'aventurier! +10 PV max")
				case "Tunique de l'aventurier":
					if c.Equipment.Body.Name != "" {
						c.AddItem(c.Equipment.Body)
						c.MaxHealth -= 25
					}
					c.Equipment.Body = item
					c.MaxHealth += 25
					fmt.Println("Vous équipez la Tunique de l'aventurier! +25 PV max")
				case "Bottes de l'aventurier":
					if c.Equipment.Feet.Name != "" {
						c.AddItem(c.Equipment.Feet)
						c.MaxHealth -= 15
					}
					c.Equipment.Feet = item
					c.MaxHealth += 15
					fmt.Println("Vous équipez les Bottes de l'aventurier! +15 PV max")
				}
			}

			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return
		}
	}
	fmt.Printf("Vous n'avez pas d'objet '%s' dans votre inventaire.\n", itemName)
}

func (c *Character) UpgradeInventory() {
	if c.MaxInventory >= 40 {
		fmt.Println("Vous avez déjà atteint la capacité maximale d'inventaire!")
		return
	}

	if c.Gold >= 30 {
		c.MaxInventory += 10
		c.Gold -= 30
		fmt.Printf("Votre capacité d'inventaire est maintenant de %d emplacements!\n", c.MaxInventory)
	} else {
		fmt.Println("Vous n'avez pas assez d'or (30 pièces requis).")
	}
}

func (c *Character) AddXP(amount int) {
	c.CurrentXP += amount
	fmt.Printf("Vous gagnez %d points d'expérience!\n", amount)

	for c.CurrentXP >= c.RequiredXP {
		c.LevelUp()
	}
}

func (c *Character) LevelUp() {
	c.Level++
	c.CurrentXP -= c.RequiredXP
	c.RequiredXP = c.Level * 100

	c.MaxHealth += 20
	c.CurrentHealth = c.MaxHealth
	c.BaseDamage += 5

	if c.Class == "Mage" {
		c.MaxMana += 30
		c.CurrentMana = c.MaxMana
	}

	fmt.Printf("Félicitations! Vous êtes maintenant niveau %d!\n", c.Level)
	fmt.Printf("PV max: %d, Dégâts de base: %d\n", c.MaxHealth, c.BaseDamage)
	if c.Class == "Mage" {
		fmt.Printf("PM max: %d\n", c.MaxMana)
	}
}

func (c *Character) RemoveSkill(skillName string) {
	for i, skill := range c.Skills {
		if skill.Name == skillName {
			c.Skills = append(c.Skills[:i], c.Skills[i+1:]...)
			return
		}
	}
}
