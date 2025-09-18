package main

import "fmt"

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
}
