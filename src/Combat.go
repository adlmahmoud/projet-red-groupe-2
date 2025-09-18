package main

import (
	"fmt"
	"time"
)

// ----------------------------
// Types pour ennemis (MOB/Boss)
// ----------------------------
type MOB struct {
	Nom        string
	Pv         int
	PvMax      int
	Degat      int
	AttaqueNom string
}

type Boss struct {
	Nom        string
	Pv         int
	PvMax      int
	Degat      int
	AttaqueNom string
}

// ----------------------------
// Inits mobs (16)
// ordre utilisé par les zones :
// 0 Brigand, 1 Voleur, 2 Rhinoppotame, 3 Chien zombie,
// 4 Squelette Mage, 5 Mort-Vivant,
// 6 Scorpion-Araignée, 7 Crabe géant, 8 Basilic,
// 9 Troll, 10 Chauve-souris, 11 Faucon de combat,
// 12 Marionnette soldat, 13 Cavalier sans tête, 14 Chevalier noir,
// 15 Gobelin (entraînement)
// ----------------------------
func (mob *MOB) InitBrigand()       { *mob = MOB{"Brigand", 60, 60, 10, "Plaquage"} }
func (mob *MOB) InitVoleur()        { *mob = MOB{"Voleur", 60, 60, 20, "Lancer de couteau"} }
func (mob *MOB) InitRhinoppotame()  { *mob = MOB{"Rhinoppotame", 105, 105, 25, "Charge"} }
func (mob *MOB) InitChienzombie()   { *mob = MOB{"Chien zombie", 105, 105, 30, "Griffe de fer"} }
func (mob *MOB) InitSquelettemage() { *mob = MOB{"Squelette Mage", 210, 210, 30, "Tarpin Nul"} }
func (mob *MOB) InitMortvivant()    { *mob = MOB{"Mort-Vivant", 210, 210, 35, "Morsure"} }
func (mob *MOB) InitScorpionaraignee() {
	*mob = MOB{"Scorpion-Araignée", 280, 280, 40, "Piqûre empoisonnée"}
}
func (mob *MOB) InitCrabegeant()   { *mob = MOB{"Crabe géant", 310, 310, 40, "Pince Masse"} }
func (mob *MOB) InitBasilic()      { *mob = MOB{"Basilic", 350, 350, 45, "Coup de langue"} }
func (mob *MOB) InitTroll()        { *mob = MOB{"Troll", 600, 600, 60, "Coup de massue"} }
func (mob *MOB) InitChauvesouris() { *mob = MOB{"Chauve-souris géante", 630, 630, 60, "Ultrason"} }
func (mob *MOB) InitFaucondecombat() {
	*mob = MOB{"Faucon de combat", 700, 700, 75, "Piqué fulgurant"}
}
func (mob *MOB) InitMarionnettesoldat() {
	*mob = MOB{"Marionnette soldat", 805, 805, 70, "Coup d'épée"}
}
func (mob *MOB) InitCavaliersanstete() {
	*mob = MOB{"Cavalier sans tête", 835, 835, 85, "Lance perforante"}
}
func (mob *MOB) InitChevaliernoir() { *mob = MOB{"Chevalier noir", 950, 950, 85, "Excalibur"} }
func (mob *MOB) InitGobelin()       { *mob = MOB{"Gobelin d'entraînement", 40, 40, 5, "Coups maladroits"} }

// ----------------------------
// Inits bosses (6)
// ----------------------------
func (b *Boss) InitChavrot()    { *b = Boss{"Chavrot", 100, 100, 25, "Coups de vieux"} }
func (b *Boss) InitMoonlight()  { *b = Boss{"Moonlight", 200, 200, 35, "Éclipse"} }
func (b *Boss) InitAinzolgone() { *b = Boss{"Ainzolgone", 250, 250, 45, "Avada Kedavra"} }
func (b *Boss) InitBahamut()    { *b = Boss{"Bahamut", 650, 650, 55, "Morsure venimeuse"} }
func (b *Boss) InitNegal()      { *b = Boss{"Negal", 800, 800, 85, "Crocs acérés du lion"} }
func (b *Boss) InitAchlys()     { *b = Boss{"Achlys", 1200, 1200, 150, "Ynov Suprêmatique"} }

// ----------------------------
// Combat (tour par tour)
// Utilise ta structure Character (déjà déclarée ailleurs)
// ----------------------------
func Combat(player *Character, enemyName string, enemyPV, enemyDegat int, enemyAttack string) bool {
	fmt.Printf("\n⚔️  Combat contre %s !\n", enemyName)
	for player.Pv > 0 && enemyPV > 0 {
		// État
		fmt.Printf("\n%s [PV: %d/%d]  —  %s [PV: %d]\n", player.Nom, player.Pv, player.PVMax, enemyName, enemyPV)

		// Tour du joueur (attaque simple)
		fmt.Printf("➡️ %s attaque (%d dégâts)\n", player.Nom, player.Attaque1)
		enemyPV -= player.Attaque1
		time.Sleep(500 * time.Millisecond)

		if enemyPV <= 0 {
			fmt.Printf("💥 %s est vaincu !\n", enemyName)
			// récompense simple
			player.Or += 20
			player.Niveau += 1
			fmt.Printf("🎁 +20 or — Niveau %d (Or total : %d)\n", player.Niveau, player.Or)
			return true
		}

		// Tour ennemi
		fmt.Printf("⚡ %s utilise '%s' et inflige %d dégâts !\n", enemyName, enemyAttack, enemyDegat)
		player.Pv -= enemyDegat
		if player.Pv < 0 {
			player.Pv = 0
		}
		time.Sleep(500 * time.Millisecond)

		if player.Pv <= 0 {
			fmt.Printf("\n☠️ %s est mort... GAME OVER\n", player.Nom)
			return false
		}
	}
	return false
}

// ----------------------------
// Soins entre zones (petit repos automatique)
// ----------------------------
func ReposEntreZones(player *Character) {
	soin := player.PVMax / 3 // 33% des PV max
	if soin < 1 {
		soin = 1
	}
	player.Pv += soin
	if player.Pv > player.PVMax {
		player.Pv = player.PVMax
	}
	// régénère aussi un peu de mana si présent
	if player.ManaMax > 0 {
		player.ManaActuel += player.ManaMax / 2
		if player.ManaActuel > player.ManaMax {
			player.ManaActuel = player.ManaMax
		}
	}
	fmt.Printf("\n⛺ Repos : vous récupérez %d PV (PV actuels : %d/%d) et un peu de mana.\n", soin, player.Pv, player.PVMax)
}

// ----------------------------
// ZONES (toutes)
// ----------------------------
func ZonePlaine(player *Character, mobs [16]MOB, bosses [6]Boss) bool {
	fmt.Println("\n🌼 [Zone 1] Plaine en fleurs — Brigand & Voleur — Boss: Chavrot")
	if !Combat(player, mobs[0].Nom, mobs[0].Pv, mobs[0].Degat, mobs[0].AttaqueNom) {
		return false
	}
	if !Combat(player, mobs[1].Nom, mobs[1].Pv, mobs[1].Degat, mobs[1].AttaqueNom) {
		return false
	}
	if !Combat(player, bosses[0].Nom, bosses[0].Pv, bosses[0].Degat, bosses[0].AttaqueNom) {
		return false
	}
	ReposEntreZones(player)
	return true
}

func ZoneTanner(player *Character, mobs [16]MOB, bosses [6]Boss) bool {
	fmt.Println("\n🐗 [Zone 2] Tanner des Alpha — Rhinoppotame & Chien zombie — Boss: Moonlight")
	if !Combat(player, mobs[2].Nom, mobs[2].Pv, mobs[2].Degat, mobs[2].AttaqueNom) {
		return false
	}
	if !Combat(player, mobs[3].Nom, mobs[3].Pv, mobs[3].Degat, mobs[3].AttaqueNom) {
		return false
	}
	if !Combat(player, bosses[1].Nom, bosses[1].Pv, bosses[1].Degat, bosses[1].AttaqueNom) {
		return false
	}
	ReposEntreZones(player)
	return true
}

func ZoneTour(player *Character, mobs [16]MOB, bosses [6]Boss) bool {
	fmt.Println("\n🏰 [Zone 3] Tour du Jugement — Squelette Mage & Mort-Vivant — Boss: Ainzolgone")
	if !Combat(player, mobs[4].Nom, mobs[4].Pv, mobs[4].Degat, mobs[4].AttaqueNom) {
		return false
	}
	if !Combat(player, mobs[5].Nom, mobs[5].Pv, mobs[5].Degat, mobs[5].AttaqueNom) {
		return false
	}
	if !Combat(player, bosses[2].Nom, bosses[2].Pv, bosses[2].Degat, bosses[2].AttaqueNom) {
		return false
	}
	ReposEntreZones(player)
	return true
}

func ZoneDesert(player *Character, mobs [16]MOB, bosses [6]Boss) bool {
	fmt.Println("\n🏜️ [Zone 4] Désert du tombeau — Scorpion-Araignée, Crabe géant, Basilic — Boss: Bahamut")
	if !Combat(player, mobs[6].Nom, mobs[6].Pv, mobs[6].Degat, mobs[6].AttaqueNom) {
		return false
	}
	if !Combat(player, mobs[7].Nom, mobs[7].Pv, mobs[7].Degat, mobs[7].AttaqueNom) {
		return false
	}
	if !Combat(player, mobs[8].Nom, mobs[8].Pv, mobs[8].Degat, mobs[8].AttaqueNom) {
		return false
	}
	if !Combat(player, bosses[3].Nom, bosses[3].Pv, bosses[3].Degat, bosses[3].AttaqueNom) {
		return false
	}
	ReposEntreZones(player)
	return true
}

func ZoneCrypte(player *Character, mobs [16]MOB, bosses [6]Boss) bool {
	fmt.Println("\n🕳 [Zone 5] Crypte arrière-cour — Troll, Chauve-souris, Faucon de combat — Boss: Negal")
	if !Combat(player, mobs[9].Nom, mobs[9].Pv, mobs[9].Degat, mobs[9].AttaqueNom) {
		return false
	}
	if !Combat(player, mobs[10].Nom, mobs[10].Pv, mobs[10].Degat, mobs[10].AttaqueNom) {
		return false
	}
	if !Combat(player, mobs[11].Nom, mobs[11].Pv, mobs[11].Degat, mobs[11].AttaqueNom) {
		return false
	}
	if !Combat(player, bosses[4].Nom, bosses[4].Pv, bosses[4].Degat, bosses[4].AttaqueNom) {
		return false
	}
	ReposEntreZones(player)
	return true
}

func ZoneChateau(player *Character, mobs [16]MOB, bosses [6]Boss) bool {
	fmt.Println("\n🏯 [Zone 6] Château en ruine — Marionnette soldat, Cavalier sans tête, Chevalier noir — Boss: Achlys")
	if !Combat(player, mobs[12].Nom, mobs[12].Pv, mobs[12].Degat, mobs[12].AttaqueNom) {
		return false
	}
	if !Combat(player, mobs[13].Nom, mobs[13].Pv, mobs[13].Degat, mobs[13].AttaqueNom) {
		return false
	}
	if !Combat(player, mobs[14].Nom, mobs[14].Pv, mobs[14].Degat, mobs[14].AttaqueNom) {
		return false
	}
	if !Combat(player, bosses[5].Nom, bosses[5].Pv, bosses[5].Degat, bosses[5].AttaqueNom) {
		return false
	}
	ReposEntreZones(player)
	return true
}

// ----------------------------
// MAIN (ne redéclare pas Character / Item)
// Assure-toi que Character et Item existent dans le même package.
// ----------------------------
func main() {
	fmt.Println("=== Bienvenue dans YNOV KINGDOM ===")

	var nom string
	fmt.Print("Entrez le nom de votre héros : ")
	fmt.Scanln(&nom)

	// On utilise TA structure Character déjà déclarée ailleurs.
	player := Character{
		Nom:        nom,
		Classe:     "Humain",
		PVMax:      100,
		Pv:         100,
		ManaMax:    100,
		ManaActuel: 100,
		Attaque1:   25,
		Attaque2:   0,
		Inventaire: []Item{{Nom: "Potion", Quantite: 2}},
		Niveau:     1,
		Or:         0,
	}

	// Initialisation des mobs (tous les 16)
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
	mobs[15].InitGobelin()

	// Initialisation des bosses (6)
	var bosses [6]Boss
	bosses[0].InitChavrot()
	bosses[1].InitMoonlight()
	bosses[2].InitAinzolgone()
	bosses[3].InitBahamut()
	bosses[4].InitNegal()
	bosses[5].InitAchlys()

	// Progression complète par zones
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
}
