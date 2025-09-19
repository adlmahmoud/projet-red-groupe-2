package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/fatih/color"
)

func Combat(player *Character, monster *Monster, zoneName string, goldReward, xpReward int) bool {
	color.Red("\t=== COMBAT CONTRE %s ===", monster.Name)
	color.Blue("\tZone: %s", zoneName)

	playerTurn := player.Initiative >= monster.Initiative

	if playerTurn {
		color.Green("\t%s a l'initiative et commence le combat!", player.Name)
	} else {
		color.Red("\t%s a l'initiative et commence le combat!", monster.Name)
	}

	round := 1
	scanner := bufio.NewScanner(os.Stdin)

	for player.CurrentHealth > 0 && monster.CurrentHealth > 0 {
		color.Cyan("\n\t--- Tour %d ---", round)

		if playerTurn {
			color.Green("\tPV de %s: %d/%d", player.Name, player.CurrentHealth, player.MaxHealth)
			if player.Class == "Mage" {
				color.Green("\tPM de %s: %d/%d", player.Name, player.CurrentMana, player.MaxMana)
			}
			color.Red("\tPV de %s: %d/%d", monster.Name, monster.CurrentHealth, monster.MaxHealth)

			fmt.Println("\tOptions:")
			fmt.Println("\t1. Attaquer")
			fmt.Println("\t2. Utiliser une compétence")
			fmt.Println("\t3. Utiliser un objet")
			fmt.Println("\t4. Fuir")

			fmt.Print("\tVotre choix: ")
			scanner.Scan()
			choice := scanner.Text()

			switch choice {
			case "1":
				damage := player.BaseDamage
				monster.CurrentHealth -= damage
				color.Green("\t%s inflige %d dégâts à %s avec une attaque basique!", player.Name, damage, monster.Name)
			case "2":
				if len(player.Skills) == 0 {
					color.Red("\tVous n'avez aucune compétence!")
					continue
				}

				fmt.Println("\tCompétences disponibles:")
				for i, skill := range player.Skills {
					fmt.Printf("\t%d. %s (Dégâts: %d, Coût en mana: %d)\n", i+1, skill.Name, skill.Damage, skill.ManaCost)
				}
				fmt.Print("\tVotre choix: ")
				scanner.Scan()
				var skillChoice int
				fmt.Sscanf(scanner.Text(), "%d", &skillChoice)

				if skillChoice < 1 || skillChoice > len(player.Skills) {
					color.Red("\tChoix invalide.")
					continue
				}

				selectedSkill := player.Skills[skillChoice-1]

				if player.Class == "Mage" && player.CurrentMana < selectedSkill.ManaCost {
					color.Red("\tVous n'avez pas assez de mana pour utiliser cette compétence!")
					continue
				}

				if player.Class == "Mage" {
					player.CurrentMana -= selectedSkill.ManaCost
				}

				damage := selectedSkill.Damage
				monster.CurrentHealth -= damage
				color.Green("\t%s inflige %d dégâts à %s avec %s!", player.Name, damage, monster.Name, selectedSkill.Name)
			case "3":
				useItemInCombat(player)
				continue
			case "4":
				if rand.Intn(100) < 50 {
					color.Yellow("\tVous réussissez à fuir le combat!")
					return true
				} else {
					color.Red("\tVous ne parvenez pas à fuir!")
				}
			default:
				color.Red("\tChoix invalide.")
				continue
			}
		} else {
			damage := monster.Damage
			player.CurrentHealth -= damage
			color.Red("\t%s inflige %d dégâts à %s avec %s!", monster.Name, damage, player.Name, monster.SkillName)
		}

		if player.CurrentHealth <= 0 {
			color.Red("\t%s a été vaincu!", player.Name)
			return false
		}

		if monster.CurrentHealth <= 0 {
			color.Green("\t%s a été vaincu!", monster.Name)
			player.Gold += goldReward
			player.addXP(xpReward)
			color.Yellow("\tVous obtenez %d pièces d'or et %d points d'expérience!", goldReward, xpReward)
			return true
		}

		playerTurn = !playerTurn
		round++

		fmt.Println("\tAppuyez sur Entrée pour continuer...")
		fmt.Scanln()
	}
	return true
}

func useItemInCombat(player *Character) {
	if len(player.Inventory) == 0 {
		color.Red("\tVotre inventaire est vide.")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\n\tQuel objet voulez-vous utiliser?")
	for i, item := range player.Inventory {
		if item.Type == "potion" || item.Type == "mana_potion" {
			fmt.Printf("\t%d. %s", i+1, item.Name)
			if item.Type == "potion" {
				fmt.Printf(" (Restaure %d PV)", item.Value)
			} else if item.Type == "mana_potion" {
				fmt.Printf(" (Restaure %d PM)", item.Value)
			}
			fmt.Println()
		}
	}
	fmt.Printf("\t%d. Retour\n", len(player.Inventory)+1)

	fmt.Print("\tVotre choix: ")
	scanner.Scan()
	choice := scanner.Text()

	var index int
	fmt.Sscanf(choice, "%d", &index)

	if index == len(player.Inventory)+1 {
		return
	}

	if index < 1 || index > len(player.Inventory) {
		color.Red("\tChoix invalide.")
		return
	}

	item := player.Inventory[index-1]

	switch item.Type {
	case "potion":
		player.CurrentHealth += item.Value
		if player.CurrentHealth > player.MaxHealth {
			player.CurrentHealth = player.MaxHealth
		}
		color.Green("\tVous utilisez une potion et récupérez %d PV. PV actuels: %d/%d", item.Value, player.CurrentHealth, player.MaxHealth)
		player.removeItem(item.Name)
	case "mana_potion":
		if player.Class == "Mage" {
			player.CurrentMana += item.Value
			if player.CurrentMana > player.MaxMana {
				player.CurrentMana = player.MaxMana
			}
			color.Green("\tVous utilisez une potion de mana et récupérez %d PM. PM actuels: %d/%d", item.Value, player.CurrentMana, player.MaxMana)
			player.removeItem(item.Name)
		} else {
			color.Red("\tSeuls les mages peuvent utiliser des potions de mana.")
		}
	default:
		color.Red("\tCet objet ne peut pas être utilisé en combat.")
	}
}

func StartTrainingFight(player *Character) {
	color.Yellow("\tVous entrez dans une zone d'entraînement...")
	goblin := initGoblin()
	if Combat(player, goblin, "Zone d'entraînement", 50, 250) {
		color.Green("\tVous avez gagné le combat d'entraînement!")
	}
}

func (c *Character) usePoisonPotion() {
	for i, item := range c.Inventory {
		if item.Type == "poison_potion" {
			color.Red("\tVous utilisez une potion de poison!")

			for j := 0; j < 3; j++ {
				c.CurrentHealth -= item.Value
				if c.CurrentHealth < 0 {
					c.CurrentHealth = 0
				}
				color.Red("\tPoison! Vous perdez %d PV. PV actuels: %d/%d", item.Value, c.CurrentHealth, c.MaxHealth)

				if c.CurrentHealth <= 0 {
					color.Red("\tVous avez été vaincu par le poison!")
					return
				}

				time.Sleep(1 * time.Second)
			}

			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return
		}
	}
	color.Red("\tVous n'avez pas de potion de poison dans votre inventaire.")
}
