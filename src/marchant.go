package main

import "fmt"

<<<<<<< HEAD
type Item struct {
	Nom      string
	Quantite int
}

type Character struct {
	Classe     string
	Or         int
	Inventaire []Item
	Attaque1   int
	// Tu peux ajouter d'autres champs nécessaires ici
}

func (player *Character) Merchant() {
	fmt.Println("\n=== Marchand ===")
	fmt.Println("Bonjour voyageur ! La première potion est gratuite pour vous aider dans votre aventure.")
	player.Inventaire = append(player.Inventaire, Item{"Potion de soin", 1})
	fmt.Println("Vous avez reçu une potion de soin gratuitement !")

	for {
		fmt.Println("Que voulez-vous faire ?")
		fmt.Println("1 - Acheter une potion de soin (15 or)")
		fmt.Println("2 - Acheter une potion de mana (20 or) [Uniquement Mage]")
		fmt.Println("3 - Acheter une Épée (30 or)")
		fmt.Println("4 - Acheter une Lance (80 or)")
		fmt.Println("5 - Acheter une Hache de guerre (150 or)")
		fmt.Println("6 - Acheter le sabre Kiketsu (35 or)")
		fmt.Println("7 - Acheter le sabre Shisui (85 or)")
		fmt.Println("8 - Acheter le katana Enme (150 or)")
		fmt.Println("9 - Acheter Panoplie de Chavrot (50 or)")
		fmt.Println("10 - Acheter Panoplie de Moonlight (80 or)")
		fmt.Println("11 - Acheter Panoplie de Ainzolgone (100 or)")
		fmt.Println("12 - Acheter Panoplie de Bahamut (125 or)")
		fmt.Println("13 - Acheter Panoplie de Negal (150 or)")
		fmt.Println("14 - Acheter Panoplie d'Achlys (300 or)")
		if player.Classe == "Mage" {
			fmt.Println("15 - Acheter le sort Éboulement (50 or)")
			fmt.Println("16 - Acheter le sort Éclair pourfandeur (80 or)")
			fmt.Println("17 - Acheter le sort Horde de corbeaux (120 or)")
			fmt.Println("18 - Acheter le sort Météore géant (150 or)")
		}
		fmt.Println("0 - Quitter le marchand")

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			if player.Or >= 15 {
				player.Or -= 15
				player.Inventaire = append(player.Inventaire, Item{"Potion de soin", 1})
				fmt.Println("Vous avez acheté une potion de soin !")
			} else {
				fmt.Println("Pas assez d'or pour acheter cette potion.")
			}
		case 2:
			if player.Classe != "Mage" {
				fmt.Println("Option invalide pour votre classe.")
				continue
			}
			if player.Or >= 20 {
				player.Or -= 20
				player.Inventaire = append(player.Inventaire, Item{"Potion de mana", 1})
				fmt.Println("Vous avez acheté une potion de mana !")
			} else {
				fmt.Println("Pas assez d'or pour acheter cette potion de mana.")
			}
		case 3:
			if player.Or >= 30 {
				player.Or -= 30
				player.Inventaire = append(player.Inventaire, Item{"Épée", 1})
				fmt.Println("Vous avez acheté une Épée !")
			} else {
				fmt.Println("Pas assez d'or pour acheter cette Épée.")
			}
		case 4:
			if player.Or >= 80 {
				player.Or -= 80
				player.Inventaire = append(player.Inventaire, Item{"Lance", 1})
				fmt.Println("Vous avez acheté une Lance !")
			} else {
				fmt.Println("Pas assez d'or pour acheter cette Lance.")
			}
		case 5:
			if player.Or >= 150 {
				player.Or -= 150
				player.Inventaire = append(player.Inventaire, Item{"Hache de guerre", 1})
				fmt.Println("Vous avez acheté une Hache de guerre !")
			} else {
				fmt.Println("Pas assez d'or pour acheter cette Hache de guerre.")
			}
		case 6:
			if player.Or >= 35 {
				player.Or -= 35
				player.Inventaire = append(player.Inventaire, Item{"Sabre Kiketsu", 1})
				fmt.Println("Vous avez acheté le sabre Kiketsu !")
			} else {
				fmt.Println("Pas assez d'or pour acheter ce sabre Kiketsu.")
			}
		case 7:
			if player.Or >= 85 {
				player.Or -= 85
				player.Inventaire = append(player.Inventaire, Item{"Sabre Shisui", 1})
				fmt.Println("Vous avez acheté le sabre Shisui !")
			} else {
				fmt.Println("Pas assez d'or pour acheter ce sabre Shisui.")
			}
		case 8:
			if player.Or >= 150 {
				player.Or -= 150
				player.Inventaire = append(player.Inventaire, Item{"Katana Enme", 1})
				fmt.Println("Vous avez acheté le katana Enme !")
			} else {
				fmt.Println("Pas assez d'or pour acheter ce katana Enme.")
			}
		case 9:
			if player.Or >= 50 {
				player.Or -= 50
				player.Inventaire = append(player.Inventaire, Item{"Panoplie de Chavrot", 1})
				fmt.Println("Vous avez acheté la Panoplie de Chavrot !")
			} else {
				fmt.Println("Pas assez d'or pour acheter cette panoplie.")
			}
		case 10:
			if player.Or >= 80 {
				player.Or -= 80
				player.Inventaire = append(player.Inventaire, Item{"Panoplie de Moonlight", 1})
				fmt.Println("Vous avez acheté la Panoplie de Moonlight !")
			} else {
				fmt.Println("Pas assez d'or pour acheter cette panoplie.")
			}
		case 11:
			if player.Or >= 100 {
				player.Or -= 100
				player.Inventaire = append(player.Inventaire, Item{"Panoplie de Ainzolgone", 1})
				fmt.Println("Vous avez acheté la Panoplie de Ainzolgone !")
			} else {
				fmt.Println("Pas assez d'or pour acheter cette panoplie.")
			}
		case 12:
			if player.Or >= 125 {
				player.Or -= 125
				player.Inventaire = append(player.Inventaire, Item{"Panoplie de Bahamut", 1})
				fmt.Println("Vous avez acheté la Panoplie de Bahamut !")
			} else {
				fmt.Println("Pas assez d'or pour acheter cette panoplie.")
			}
		case 13:
			if player.Or >= 150 {
				player.Or -= 150
				player.Inventaire = append(player.Inventaire, Item{"Panoplie de Negal", 1})
				fmt.Println("Vous avez acheté la Panoplie de Negal !")
			} else {
				fmt.Println("Pas assez d'or pour acheter cette panoplie.")
			}
		case 14:
			if player.Or >= 300 {
				player.Or -= 300
				player.Inventaire = append(player.Inventaire, Item{"Panoplie d'Achlys", 1})
				fmt.Println("Vous avez acheté la Panoplie d'Achlys !")
			} else {
				fmt.Println("Pas assez d'or pour acheter cette panoplie.")
			}
		case 15:
			if player.Classe != "Mage" {
				fmt.Println("Option invalide pour votre classe.")
				continue
			}
			if player.Or >= 50 {
				player.Or -= 50
				fmt.Println("Vous avez appris le sort Éboulement !")
				// Ajouter ici le code pour apprendre le sort (ex: ajouter à une liste de sorts)
			} else {
				fmt.Println("Pas assez d'or pour acheter ce sort.")
			}
		case 16:
			if player.Classe != "Mage" {
				fmt.Println("Option invalide pour votre classe.")
				continue
			}
			if player.Or >= 80 {
				player.Or -= 80
				fmt.Println("Vous avez appris le sort Éclair pourfandeur !")
			} else {
				fmt.Println("Pas assez d'or pour acheter ce sort.")
			}
		case 17:
			if player.Classe != "Mage" {
				fmt.Println("Option invalide pour votre classe.")
				continue
			}
			if player.Or >= 120 {
				player.Or -= 120
				fmt.Println("Vous avez appris le sort Horde de corbeaux !")
			} else {
				fmt.Println("Pas assez d'or pour acheter ce sort.")
			}
		case 18:
			if player.Classe != "Mage" {
				fmt.Println("Option invalide pour votre classe.")
				continue
			}
			if player.Or >= 150 {
				player.Or -= 150
				fmt.Println("Vous avez appris le sort Météore géant !")
			} else {
				fmt.Println("Pas assez d'or pour acheter ce sort.")
			}
		case 0:
			fmt.Println("Merci et bonne aventure !")
			return
		default:
			fmt.Println("Choix invalide, réessayez.")
		}
		fmt.Printf("Or restant : %d\n\n", player.Or)
	}
=======
func main() {
	fmt.Println("=== Bienvenue dans YNOV KINGDOM ===")
	var nom string
	fmt.Print("Entrez le nom de votre héros : ")
	fmt.Scanln(&nom)

	player := Character{Nom: nom, Pv: 100, PVMax: 100, Attaque1: 20, Or: 0, Niveau: 1}

	// Initialisation des mobs
	var mobs [16]MOB
	mobs[0].InitBrigand()
	mobs[1].InitVoleur()
	mobs[2].InitRhinoppotame()
	mobs[3].InitChienzombie()
	mobs[4].InitSquelettemage()
	mobs[5].InitMortvivant()
	mobs[6].InitScorpionaraignee()
	mobs[7].InitCrabegeant()
	mobs[8].InitBasilic()
	mobs[9].InitTroll()
	mobs[10].InitChauvesouris()
	mobs[11].InitFaucondecombat()
	mobs[12].InitMarionnettesoldat()
	mobs[13].InitCavaliersanstete()
	mobs[14].InitChevaliernoir()

	// Initialisation des boss
	var bosses [6]Boss
	bosses[0].InitChavrot()
	bosses[1].InitMoonlight()
	bosses[2].InitAinzolgone()
	bosses[3].InitBahamut()
	bosses[4].InitNegal()
	bosses[5].InitAchlys()

	// Progression
	if !ZonePlaine(&player, mobs, bosses) {
		return
	}
	if !ZoneTanner(&player, mobs, bosses) {
		return
	}
	if !ZoneTour(&player, mobs, bosses) {
		return
	}
	if !ZoneDesert(&player, mobs, bosses) {
		return
	}
	if !ZoneCrypte(&player, mobs, bosses) {
		return
	}
	if !ZoneChateau(&player, mobs, bosses) {
		return
	}

	fmt.Println("\n🏆 Félicitations ! Vous avez vaincu Achlys et libéré YNOV KINGDOM ! 🏆")
>>>>>>> 80d0d88a2102472cfac0ba484f5483acf1e3eb29
}
