package main

import "fmt"

type MOB struct {
	Nom   string
	Pv    int
	PvMax int
	Degat int
}

func (mob *MOB) InitGobelin() {
	*mob = MOB{"Gobelin d'entraînement", 40, 40, 5}
}
func (mob *MOB) InitBrigand() {
	*mob = MOB{"Brigand", 60, 60, 10}
}
func (mob *MOB) InitVoleur() {
	*mob = MOB{"Voleur", 60, 60, 20}
}
func (mob *MOB) InitRhinoppotame() {
	*mob = MOB{"Rhinoppotame", 105, 105, 25}
}
func (mob *MOB) InitChienzombie() {
	*mob = MOB{"Chien zombie", 105, 105, 30}
}
func (mob *MOB) InitSquelettemage() {
	*mob = MOB{"Squelette Mage", 210, 210, 30}
}
func (mob *MOB) InitMortvivant() {
	*mob = MOB{"Mort-Vivant", 210, 210, 40}
}
func (mob *MOB) InitScorpionaraignee() {
	*mob = MOB{"Scorpion-Araignée", 280, 280, 40}
}
func (mob *MOB) InitCrabegeant() {
	*mob = MOB{"Crabe géant", 310, 310, 50}
}
func (mob *MOB) InitBasilic() {
	*mob = MOB{"Basilic", 350, 350, 50}
}
func (mob *MOB) InitTroll() {
	*mob = MOB{"Troll", 600, 600, 60}
}
func (mob *MOB) InitChauvesouris() {
	*mob = MOB{"Chauve-souris", 630, 630, 60}
}
func (mob *MOB) InitFaucondecombat() {
	*mob = MOB{"Faucon de combat", 700, 700, 70}
}
func (mob *MOB) InitMarionnettesoldat() {
	*mob = MOB{"Marionnette soldat", 805, 805, 80}
}
func (mob *MOB) InitCavaliersanstete() {
	*mob = MOB{"Cavalier sans tête", 835, 835, 80}
}
func (mob *MOB) InitChevaliernoir() {
	*mob = MOB{"Chevalier noir", 950, 950, 90}
}

// =================================
// BOSS
// =================================

type Boss struct {
	Nom   string
	Pv    int
	PvMax int
	Degat int
}

func (b *Boss) InitChavrot() {
	*b = Boss{"Chavrot", 100, 100, 25}
}
func (b *Boss) InitMoonlight() {
	*b = Boss{"Moonlight", 200, 200, 35}
}
func (b *Boss) InitAinzolgone() {
	*b = Boss{"Ainzolgone", 250, 250, 45}
}
func (b *Boss) InitBahamut() {
	*b = Boss{"Bahamut", 650, 650, 105}
}
func (b *Boss) InitNegal() {
	*b = Boss{"Negal", 800, 800, 165}
}
func (b *Boss) InitAchlys() {
	*b = Boss{"Achlys", 1200, 1200, 240}
}

// =================================
// MAIN
// =================================

func main() {
	// Initialisation des 16 mobs
	var mobs [16]MOB
	mobs[0].InitGobelin()
	mobs[1].InitBrigand()
	mobs[2].InitVoleur()
	mobs[3].InitRhinoppotame()
	mobs[4].InitChienzombie()
	mobs[5].InitSquelettemage()
	mobs[6].InitMortvivant()
	mobs[7].InitScorpionaraignee()
	mobs[8].InitCrabegeant()
	mobs[9].InitBasilic()
	mobs[10].InitTroll()
	mobs[11].InitChauvesouris()
	mobs[12].InitFaucondecombat()
	mobs[13].InitMarionnettesoldat()
	mobs[14].InitCavaliersanstete()
	mobs[15].InitChevaliernoir()

	// Initialisation des 6 boss
	var bosses [6]Boss
	bosses[0].InitChavrot()
	bosses[1].InitMoonlight()
	bosses[2].InitAinzolgone()
	bosses[3].InitBahamut()
	bosses[4].InitNegal()
	bosses[5].InitAchlys()

	// Affichage de test
	fmt.Println("===== MOBS =====")
	for i, mob := range mobs {
		fmt.Printf("[%2d] %-25s | PV: %4d/%-4d | Dégâts: %3d\n", i+1, mob.Nom, mob.Pv, mob.PvMax, mob.Degat)
	}

	fmt.Println("\n===== BOSSES =====")
	for i, boss := range bosses {
		fmt.Printf("[%2d] %-25s | PV: %4d/%-4d | Dégâts: %3d\n", i+1, boss.Nom, boss.Pv, boss.PvMax, boss.Degat)
	}
}
