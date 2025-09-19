package main

import "math/rand"

func InitGoblin() *Monster {
	return &Monster{
		Name:          "Gobelin d'entraînement",
		MaxHealth:     40,
		CurrentHealth: 40,
		Damage:        5,
		SkillName:     "Coup de poing",
		Initiative:    rand.Intn(5) + 1,
	}
}

func CreateBoss(zoneName string) *Monster {
	switch zoneName {
	case "Plaine en fleur":
		return &Monster{
			Name:          "Chavrot",
			MaxHealth:     100,
			CurrentHealth: 100,
			Damage:        25,
			SkillName:     "Coups de vieux",
			Initiative:    rand.Intn(8) + 1,
		}
	case "Tanner des alpha":
		return &Monster{
			Name:          "Moonlight",
			MaxHealth:     200,
			CurrentHealth: 200,
			Damage:        35,
			SkillName:     "Eclipse",
			Initiative:    rand.Intn(9) + 1,
		}
	case "La tour du jugement":
		return &Monster{
			Name:          "Ainzolgone",
			MaxHealth:     250,
			CurrentHealth: 250,
			Damage:        45,
			SkillName:     "Avada kedabra",
			Initiative:    rand.Intn(7) + 1,
		}
	case "Désert aride du tombeau de nazarick":
		return &Monster{
			Name:          "Bahamut",
			MaxHealth:     650,
			CurrentHealth: 650,
			Damage:        55,
			SkillName:     "Morsure venimeuse",
			Initiative:    rand.Intn(6) + 1,
		}
	case "Crypte de l'arrière cour":
		return &Monster{
			Name:          "Negal",
			MaxHealth:     800,
			CurrentHealth: 800,
			Damage:        85,
			SkillName:     "Ragnarok",
			Initiative:    rand.Intn(10) + 1,
		}
	case "Château en ruine":
		return &Monster{
			Name:          "Achlys",
			MaxHealth:     1200,
			CurrentHealth: 1200,
			Damage:        150,
			SkillName:     "Ynov suprêmatique",
			Initiative:    rand.Intn(12) + 1,
		}
	default:
		return InitGoblin()
	}
}
