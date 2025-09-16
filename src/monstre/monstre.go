package main

type MOB struct {
	Nom   string
	Pv    int
	PvMax int
	Degat int
}

func (mob *MOB) InitBrigand() {
	*mob = MOB{
		Nom:   "Brigand",
		Pv:    60,
		PvMax: 60,
		Degat: 10,
	}
}
func (mob *MOB) InitVoleur() {
	*mob = MOB{
		Nom:   "Voleur",
		Pv:    60,
		PvMax: 60,
		Degat: 20,
	}
}
func (mob *MOB) InitRhinoppotame() {
	*mob = MOB{
		Nom:   "Rhinoppotame",
		Pv:    105,
		PvMax: 105,
		Degat: 25,
	}
}
func (mob *MOB) InitChienzombie() {
	*mob = MOB{
		Nom:   "Chien zombie",
		Pv:    105,
		PvMax: 105,
		Degat: 30,
	}
}
func (mob *MOB) InitSquelettemage() {
	*mob = MOB{
		Nom:   "Squelette Mage",
		Pv:    210,
		PvMax: 210,
		Degat: 30,
	}
}
func (mob *MOB) InitMortvivant() {
	*mob = MOB{
		Nom:   "Mort-Vivant",
		Pv:    210,
		PvMax: 210,
		Degat: 40,
	}
}
func (mob *MOB) InitScorpionaraignée() {
	*mob = MOB{
		Nom:   "Scorpion-Araignée",
		Pv:    280,
		PvMax: 280,
		Degat: 40,
	}
}
func (mob *MOB) InitCrabegéant() {
	*mob = MOB{
		Nom:   "Crabe géant",
		Pv:    310,
		PvMax: 310,
		Degat: 50,
	}
}
func (mob *MOB) InitBasilic() {
	*mob = MOB{
		Nom:   "Basilc",
		Pv:    350,
		PvMax: 350,
		Degat: 50,
	}
}
func (mob *MOB) InitTroll() {
	*mob = MOB{
		Nom:   "Troll",
		Pv:    600,
		PvMax: 600,
		Degat: 60,
	}
}
func (mob *MOB) InitChauvesouris() {
	*mob = MOB{
		Nom:   "Chauvre-souris",
		Pv:    630,
		PvMax: 630,
		Degat: 60,
	}
}
func (mob *MOB) InitFaucondecombat() {
	*mob = MOB{
		Nom:   "Faucon de combat",
		Pv:    700,
		PvMax: 700,
		Degat: 70,
	}
}
func (mob *MOB) InitMarionnettesoldat() {
	*mob = MOB{
		Nom:   "Marionnette soldat",
		Pv:    805,
		PvMax: 805,
		Degat: 80,
	}
}
func (mob *MOB) InitCavaliersanstête() {
	*mob = MOB{
		Nom:   "Cavalier sans tête",
		Pv:    835,
		PvMax: 835,
		Degat: 80,
	}
}
func (mob *MOB) InitChevaliernoir() {
	*mob = MOB{
		Nom:   "Chevalier noir",
		Pv:    950,
		PvMax: 950,
		Degat: 90,
	}
}

type Boss struct {
	Nom   string
	Pv    int
	PvMax int
	Degat int
}

func (mob *Boss) initChavrot() {
	*mob = Boss{
		Nom:   "Chavrot",
		Pv:    100,
		PvMax: 100,
		Degat: 25,
	}
}
func (mob *Boss) initMoonlight() {
	*mob = Boss{
		Nom:   "Moonlight",
		Pv:    200,
		PvMax: 200,
		Degat: 35,
	}
}
func (mob *Boss) initAinzolgone() {
	*mob = Boss{
		Nom:   "Ainzolgone",
		Pv:    250,
		PvMax: 250,
		Degat: 45,
	}
}
func (mob *Boss) initBahamut() {
	*mob = Boss{
		Nom:   "Bahamut",
		Pv:    650,
		PvMax: 650,
		Degat: 105,
	}
}
func (mob *Boss) initNegal() {
	*mob = Boss{
		Nom:   "Negal",
		Pv:    800,
		PvMax: 800,
		Degat: 165,
	}
}
func (mob *Boss) initAchlys() {
	*mob = Boss{
		Nom:   "Achlys",
		Pv:    1200,
		PvMax: 1200,
		Degat: 240,
	}
}
