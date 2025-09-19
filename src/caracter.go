package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
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

	fmt.Println("\n\t=== CRÉATION DU PERSONNAGE ===")

	for {
		fmt.Print("\tEntrez le nom de votre personnage: ")
		scanner.Scan()
		name := strings.TrimSpace(scanner.Text())

		if IsValidName(name) {
			player.Name = formatName(name)
			break
		} else {
			fmt.Println("\tNom invalide. Utilisez uniquement des lettres.")
		}
	}

	for {
		fmt.Println("\tChoisissez votre classe:")
		fmt.Println("\t1. Humain")
		fmt.Println("\t2. Mage")
		fmt.Println("\t3. Samouraï (CHEATÉ!)")
		fmt.Print("\tVotre choix: ")
		scanner.Scan()
		classChoice := scanner.Text()

		switch classChoice {
		case "1":
			player.Class = "Humain"
			player.MaxHealth = 200
			player.CurrentHealth = 100
			player.Inventory = []Item{
				{Name: "Potion", Description: "Restaure 100 PV", Type: "potion", Value: 100, Price: 10},
				{Name: "Potion", Description: "Restaure 100 PV", Type: "potion", Value: 100, Price: 10},
			}
			player.Skills = []Skill{
				{Name: "Coups de poing", Damage: 30, ManaCost: 0},
				{Name: "Coups de pieds", Damage: 30, ManaCost: 0},
			}
			player.Gold = 100
			player.BaseDamage = 30
			player.Initiative = 7
		case "2":
			player.Class = "Mage"
			player.MaxHealth = 180
			player.CurrentHealth = 90
			player.MaxMana = 300
			player.CurrentMana = 150
			player.Inventory = []Item{
				{Name: "Potion", Description: "Restaure 100 PV", Type: "potion", Value: 100, Price: 10},
				{Name: "Potion", Description: "Restaure 100 PV", Type: "potion", Value: 100, Price: 10},
				{Name: "Potion de mana", Description: "Restaure 100 PM", Type: "mana_potion", Value: 100, Price: 20},
				{Name: "Bâton magique", Description: "Arme de mage", Type: "weapon", Value: 40, Price: 80},
			}
			player.Skills = []Skill{
				{Name: "Luciole de feu", Damage: 40, ManaCost: 15},
				{Name: "Éclair", Damage: 50, ManaCost: 25},
			}
			player.Gold = 80
			player.BaseDamage = 20
			player.Initiative = 5
		case "3":
			player.Class = "Samouraï"
			player.MaxHealth = 400
			player.CurrentHealth = 400
			player.Inventory = []Item{
				{Name: "Potion", Description: "Restaure 200 PV", Type: "potion", Value: 200, Price: 10},
				{Name: "Potion", Description: "Restaure 200 PV", Type: "potion", Value: 200, Price: 10},
				{Name: "Katana", Description: "Arme de samouraï", Type: "weapon", Value: 80, Price: 100},
			}
			player.Skills = []Skill{
				{Name: "Coups de poing", Damage: 60, ManaCost: 0},
				{Name: "Coups de pieds", Damage: 60, ManaCost: 0},
				{Name: "Dans 3 pas tu meurs", Damage: 120, ManaCost: 0},
			}
			player.Gold = 200
			player.BaseDamage = 60
			player.Initiative = 10
			color.Red("\t⚔️  SAMOURAÏ CHEATÉ ACTIVÉ! ⚔️")
		default:
			fmt.Println("\tChoix invalide. Veuillez choisir 1, 2 ou 3.")
			continue
		}
		break
	}

	player.Level = 1
	player.MaxInventory = 15
	player.CurrentXP = 0
	player.RequiredXP = 100

	fmt.Printf("\tFélicitations! %s le %s est créé!\n", player.Name, player.Class)
	player.displayInfo()

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

func formatName(name string) string {
	name = strings.ToLower(name)
	if len(name) > 0 {
		name = strings.ToUpper(string(name[0])) + name[1:]
	}
	return name
}

func (c *Character) displayInfo() {
	color.Cyan("\n\t=== INFORMATIONS DU PERSONNAGE ===")
	color.White("\tNom: %s", c.Name)
	color.White("\tClasse: %s", c.Class)
	color.White("\tNiveau: %d", c.Level)
	color.White("\tPV: %d/%d", c.CurrentHealth, c.MaxHealth)
	if c.Class == "Mage" {
		color.White("\tPM: %d/%d", c.CurrentMana, c.MaxMana)
	}
	color.Yellow("\tOr: %d pièces", c.Gold)
	color.Green("\tXP: %d/%d", c.CurrentXP, c.RequiredXP)
	color.White("\tDégâts de base: %d", c.BaseDamage)
	color.White("\tInitiative: %d", c.Initiative)

	if len(c.Skills) > 0 {
		color.Cyan("\tCompétences:")
		for _, skill := range c.Skills {
			color.White("\t  - %s (Dégâts: %d, Coût en mana: %d)", skill.Name, skill.Damage, skill.ManaCost)
		}
	}

	if c.Equipment.Head.Name != "" || c.Equipment.Body.Name != "" || c.Equipment.Feet.Name != "" || c.Equipment.Weapon.Name != "" {
		color.Cyan("\tÉquipement:")
		if c.Equipment.Head.Name != "" {
			color.White("\t  Tête: %s", c.Equipment.Head.Name)
		}
		if c.Equipment.Body.Name != "" {
			color.White("\t  Corps: %s", c.Equipment.Body.Name)
		}
		if c.Equipment.Feet.Name != "" {
			color.White("\t  Pieds: %s", c.Equipment.Feet.Name)
		}
		if c.Equipment.Weapon.Name != "" {
			color.White("\t  Arme: %s", c.Equipment.Weapon.Name)
		}
	}
	fmt.Println("\tAppuyez sur Entrée pour continuer...")
	fmt.Scanln()
}

func (c *Character) UsePotion() {
	for i, item := range c.Inventory {
		if item.Type == "potion" {
			c.CurrentHealth += item.Value
			if c.CurrentHealth > c.MaxHealth {
				c.CurrentHealth = c.MaxHealth
			}
			fmt.Printf("\tVous utilisez une potion et récupérez %d PV. PV actuels: %d/%d\n", item.Value, c.CurrentHealth, c.MaxHealth)
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return
		}
	}
	fmt.Println("\tVous n'avez pas de potion de vie dans votre inventaire.")
}

func (c *Character) useManaPotion() {
	for i, item := range c.Inventory {
		if item.Type == "mana_potion" {
			c.CurrentMana += item.Value
			if c.CurrentMana > c.MaxMana {
				c.CurrentMana = c.MaxMana
			}
			fmt.Printf("\tVous utilisez une potion de mana et récupérez %d PM. PM actuels: %d/%d\n", item.Value, c.CurrentMana, c.MaxMana)
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return
		}
	}
	fmt.Println("\tVous n'avez pas de potion de mana dans votre inventaire.")
}

func (c *Character) isDead() bool {
	return c.CurrentHealth <= 0
}

func (c *Character) revive() {
	c.CurrentHealth = c.MaxHealth / 2
	if c.Class == "Mage" {
		c.CurrentMana = c.MaxMana
	}
	fmt.Printf("\t%s est ressuscité avec %d PV!\n", c.Name, c.CurrentHealth)
}

func (c *Character) learnSpell(spellName string) {
	for _, skill := range c.Skills {
		if skill.Name == spellName {
			fmt.Println("\tVous connaissez déjà ce sort!")
			return
		}
	}

	switch spellName {
	case "Boule de feu":
		c.Skills = append(c.Skills, Skill{Name: "Boule de feu", Damage: 60, ManaCost: 20})
	case "Éclair pourfendeur":
		c.Skills = append(c.Skills, Skill{Name: "Éclair pourfendeur", Damage: 80, ManaCost: 30})
	case "Éboulement":
		c.Skills = append(c.Skills, Skill{Name: "Éboulement", Damage: 100, ManaCost: 40})
	case "Horde de corbeaux":
		c.Skills = append(c.Skills, Skill{Name: "Horde de corbeaux", Damage: 200, ManaCost: 80})
	}
	fmt.Printf("\tVous avez appris le sort: %s!\n", spellName)
}

func (c *Character) addItem(item Item) bool {
	if len(c.Inventory) >= c.MaxInventory {
		fmt.Println("\tVotre inventaire est plein!")
		return false
	}
	c.Inventory = append(c.Inventory, item)
	return true
}

func (c *Character) removeItem(itemName string) bool {
	for i, item := range c.Inventory {
		if item.Name == itemName {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return true
		}
	}
	return false
}

func (c *Character) hasItem(itemName string) bool {
	for _, item := range c.Inventory {
		if item.Name == itemName {
			return true
		}
	}
	return false
}

func (c *Character) countItem(itemName string) int {
	count := 0
	for _, item := range c.Inventory {
		if item.Name == itemName {
			count++
		}
	}
	return count
}

func (c *Character) equipItem(itemName string) {
	for i, item := range c.Inventory {
		if item.Name == itemName {
			switch item.Type {
			case "weapon":
				if c.Equipment.Weapon.Name != "" {
					c.addItem(c.Equipment.Weapon)
				}
				c.Equipment.Weapon = item
				if c.Class == "Humain" && item.Name == "Gants de combat" {
					c.Skills = append(c.Skills, Skill{Name: "Patate de forain", Damage: 80, ManaCost: 0})
					fmt.Println("\tVous équipez les gants de combat et apprenez 'Patate de forain'!")
				} else if c.Class == "Mage" && item.Name == "Bâton magique" {
					c.Skills = append(c.Skills, Skill{Name: "Boule de feu", Damage: 80, ManaCost: 30})
					fmt.Println("\tVous équipez le bâton magique et apprenez 'Boule de feu'!")
				} else if c.Class == "Samouraï" && item.Name == "Katana" {
					c.Skills = append(c.Skills, Skill{Name: "Dans 3 pas tu meurs", Damage: 200, ManaCost: 0})
					fmt.Println("\tVous équipez le katana et apprenez 'Dans 3 pas tu meurs'!")
				}
			case "armor":
				switch item.Name {
				case "Chapeau de l'aventurier":
					if c.Equipment.Head.Name != "" {
						c.addItem(c.Equipment.Head)
						c.MaxHealth -= 20
					}
					c.Equipment.Head = item
					c.MaxHealth += 20
					fmt.Println("\tVous équipez le Chapeau de l'aventurier! +20 PV max")
				case "Tunique de l'aventurier":
					if c.Equipment.Body.Name != "" {
						c.addItem(c.Equipment.Body)
						c.MaxHealth -= 50
					}
					c.Equipment.Body = item
					c.MaxHealth += 50
					fmt.Println("\tVous équipez la Tunique de l'aventurier! +50 PV max")
				case "Bottes de l'aventurier":
					if c.Equipment.Feet.Name != "" {
						c.addItem(c.Equipment.Feet)
						c.MaxHealth -= 30
					}
					c.Equipment.Feet = item
					c.MaxHealth += 30
					fmt.Println("\tVous équipez les Bottes de l'aventurier! +30 PV max")
				}
			}
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return
		}
	}
	fmt.Printf("\tVous n'avez pas d'objet '%s' dans votre inventaire.\n", itemName)
}

func (c *Character) upgradeInventory() {
	if c.MaxInventory >= 40 {
		fmt.Println("\tVous avez déjà atteint la capacité maximale d'inventaire!")
		return
	}
	if c.Gold >= 30 {
		c.MaxInventory += 10
		c.Gold -= 30
		fmt.Printf("\tVotre capacité d'inventaire est maintenant de %d emplacements!\n", c.MaxInventory)
	} else {
		fmt.Println("\tVous n'avez pas assez d'or (30 pièces requis).")
	}
}

func (c *Character) addXP(amount int) {
	c.CurrentXP += amount
	fmt.Printf("\tVous gagnez %d points d'expérience!\n", amount)

	levelsGained := 0
	for c.CurrentXP >= c.RequiredXP {
		c.levelUp()
		levelsGained++
		c.CurrentXP -= c.RequiredXP
		c.RequiredXP = c.Level * 50
	}

	if levelsGained > 0 {
		fmt.Printf("\tVous avez gagné %d niveau(s)!\n", levelsGained)
	}
}

func (c *Character) levelUp() {
	c.Level++

	if c.Class == "Samouraï" {
		c.MaxHealth += 80
		c.BaseDamage += 25
	} else if c.Class == "Humain" {
		c.MaxHealth += 40
		c.BaseDamage += 15
	} else if c.Class == "Mage" {
		c.MaxHealth += 30
		c.MaxMana += 80
		c.BaseDamage += 12
	}

	c.CurrentHealth = c.MaxHealth
	if c.Class == "Mage" {
		c.CurrentMana = c.MaxMana
	}

	fmt.Printf("\tFélicitations! Vous êtes maintenant niveau %d!\n", c.Level)
	fmt.Printf("\tPV max: %d, Dégâts de base: %d\n", c.MaxHealth, c.BaseDamage)
	if c.Class == "Mage" {
		fmt.Printf("\tPM max: %d\n", c.MaxMana)
	}
}

func (c *Character) removeSkill(skillName string) {
	for i, skill := range c.Skills {
		if skill.Name == skillName {
			c.Skills = append(c.Skills[:i], c.Skills[i+1:]...)
			return
		}
	}
}
