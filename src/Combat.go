package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Monster struct {
	Name          string
	MaxHealth     int
	CurrentHealth int
	Damage        int
	SkillName     string
	Initiative    int
}

func CreateMonster(name string, health, damage int, skill string) *Monster {
	return &Monster{
		Name:          name,
		MaxHealth:     health,
		CurrentHealth: health,
		Damage:        damage,
		SkillName:     skill,
		Initiative:    rand.Intn(10) + 1,
	}
}

func Combat(player *Character, monster *Monster, zoneName string, goldReward, xpReward int) {
	fmt.Printf("=== COMBAT CONTRE %s ===\n", monster.Name)
	fmt.Printf("Zone: %s\n", zoneName)

	// Déterminer qui commence en fonction de l'initiative
	playerTurn := player.Initiative >= monster.Initiative

	if playerTurn {
		fmt.Printf("%s a l'initiative et commence le combat!\n", player.Name)
	} else {
		fmt.Printf("%s a l'initiative et commence le combat!\n", monster.Name)
	}

	round := 1
	scanner := bufio.NewScanner(os.Stdin)

	for player.CurrentHealth > 0 && monster.CurrentHealth > 0 {
		fmt.Printf("\n--- Tour %d ---\n", round)

		if playerTurn {
			// Tour du joueur
			fmt.Printf("PV de %s: %d/%d\n", player.Name, player.CurrentHealth, player.MaxHealth)
			if player.Class == "Mage" {
				fmt.Printf("PM de %s: %d/%d\n", player.Name, player.CurrentMana, player.MaxMana)
			}
			fmt.Printf("PV de %s: %d/%d\n", monster.Name, monster.CurrentHealth, monster.MaxHealth)

			fmt.Println("Options:")
			fmt.Println("1. Attaquer")
			fmt.Println("2. Utiliser une compétence")
			fmt.Println("3. Utiliser un objet")
			fmt.Println("4. Fuir")

			fmt.Print("Votre choix: ")
			scanner.Scan()
			choice := scanner.Text()

			switch choice {
			case "1":
				// Attaque basique
				damage := player.BaseDamage
				monster.CurrentHealth -= damage
				fmt.Printf("%s inflige %d dégâts à %s avec une attaque basique!\n", player.Name, damage, monster.Name)
			case "2":
				// Utiliser une compétence
				if len(player.Skills) == 0 {
					fmt.Println("Vous n'avez aucune compétence!")
					continue
				}

				fmt.Println("Compétences disponibles:")
				for i, skill := range player.Skills {
					fmt.Printf("%d. %s (Dégâts: %d, Coût en mana: %d)\n", i+1, skill.Name, skill.Damage, skill.ManaCost)
				}
				fmt.Print("Votre choix: ")
				scanner.Scan()
				var skillChoice int
				fmt.Sscanf(scanner.Text(), "%d", &skillChoice)

				if skillChoice < 1 || skillChoice > len(player.Skills) {
					fmt.Println("Choix invalide.")
					continue
				}

				selectedSkill := player.Skills[skillChoice-1]

				if player.Class == "Mage" && player.CurrentMana < selectedSkill.ManaCost {
					fmt.Println("Vous n'avez pas assez de mana pour utiliser cette compétence!")
					continue
				}

				if player.Class == "Mage" {
					player.CurrentMana -= selectedSkill.ManaCost
				}

				damage := selectedSkill.Damage
				monster.CurrentHealth -= damage
				fmt.Printf("%s inflige %d dégâts à %s avec %s!\n", player.Name, damage, monster.Name, selectedSkill.Name)
			case "3":
				// Utiliser un objet
				UseItemInCombat(player)
				continue // Le tour n'est pas terminé
			case "4":
				// Fuir
				if rand.Intn(100) < 30 { // 30% de chance de fuite
					fmt.Println("Vous réussissez à fuir le combat!")
					return
				} else {
					fmt.Println("Vous ne parvenez pas à fuir!")
				}
			default:
				fmt.Println("Choix invalide.")
				continue
			}
		} else {
			// Tour du monstre
			damage := monster.Damage
			player.CurrentHealth -= damage
			fmt.Printf("%s inflige %d dégâts à %s avec %s!\n", monster.Name, damage, player.Name, monster.SkillName)
		}

		// Vérifier si le combat est terminé
		if player.CurrentHealth <= 0 {
			fmt.Printf("%s a été vaincu!\n", player.Name)
			player.revive()
			break
		}

		if monster.CurrentHealth <= 0 {
			fmt.Printf("%s a été vaincu!\n", monster.Name)
			player.Gold += goldReward
			player.AddXP(xpReward * player.Level)
			fmt.Printf("Vous obtenez %d pièces d'or et %d points d'expérience!\n", goldReward, xpReward*player.Level)
			break
		}

		// Changer de tour
		playerTurn = !playerTurn
		round++

		fmt.Println("Appuyez sur Entrée pour continuer...")
		fmt.Scanln()
	}
}

func UseItemInCombat(player *Character) {
	if len(player.Inventory) == 0 {
		fmt.Println("Votre inventaire est vide.")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\nQuel objet voulez-vous utiliser?")
	for i, item := range player.Inventory {
		if item.Type == "potion" || item.Type == "mana_potion" {
			fmt.Printf("%d. %s", i+1, item.Name)
			if item.Type == "potion" {
				fmt.Printf(" (Restaure %d PV)", item.Value)
			} else if item.Type == "mana_potion" {
				fmt.Printf(" (Restaure %d PM)", item.Value)
			}
			fmt.Println()
		}
	}
	fmt.Printf("%d. Retour\n", len(player.Inventory)+1)

	fmt.Print("Votre choix: ")
	scanner.Scan()
	choice := scanner.Text()

	var index int
	fmt.Sscanf(choice, "%d", &index)

	if index == len(player.Inventory)+1 {
		return
	}

	if index < 1 || index > len(player.Inventory) {
		fmt.Println("Choix invalide.")
		return
	}

	item := player.Inventory[index-1]

	switch item.Type {
	case "potion":
		player.CurrentHealth += item.Value
		if player.CurrentHealth > player.MaxHealth {
			player.CurrentHealth = player.MaxHealth
		}
		fmt.Printf("Vous utilisez une potion et récupérez %d PV. PV actuels: %d/%d\n", item.Value, player.CurrentHealth, player.MaxHealth)
		player.RemoveItem(item.Name)
	case "mana_potion":
		if player.Class == "Mage" {
			player.CurrentMana += item.Value
			if player.CurrentMana > player.MaxMana {
				player.CurrentMana = player.MaxMana
			}
			fmt.Printf("Vous utilisez une potion de mana et récupérez %d PM. PM actuels: %d/%d\n", item.Value, player.CurrentMana, player.MaxMana)
			player.RemoveItem(item.Name)
		} else {
			fmt.Println("Seuls les mages peuvent utiliser des potions de mana.")
		}
	default:
		fmt.Println("Cet objet ne peut pas être utilisé en combat.")
	}
}

func StartTrainingFight(player *Character) {
	fmt.Println("Vous entrez dans une zone d'entraînement...")
	goblin := CreateMonster("Gobelin d'entraînement", 40, 5, "Coup de poing")
	Combat(player, goblin, "Zone d'entraînement", 5, 10)
}

func (c *Character) UsePoisonPotion() {
	for i, item := range c.Inventory {
		if item.Type == "poison_potion" {
			fmt.Println("Vous utilisez une potion de poison!")

			for j := 0; j < 3; j++ {
				c.CurrentHealth -= item.Value
				if c.CurrentHealth < 0 {
					c.CurrentHealth = 0
				}
				fmt.Printf("Poison! Vous perdez %d PV. PV actuels: %d/%d\n", item.Value, c.CurrentHealth, c.MaxHealth)

				if c.CurrentHealth <= 0 {
					fmt.Println("Vous avez été vaincu par le poison!")
					c.revive()
					break
				}

				time.Sleep(1 * time.Second)
			}

			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return
		}
	}
	fmt.Println("Vous n'avez pas de potion de poison dans votre inventaire.")
}
