package main

import "math/rand"

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

func CreateBoss(name string, health, damage int, skill string) *Monster {
	return &Monster{
		Name:          name,
		MaxHealth:     health,
		CurrentHealth: health,
		Damage:        damage,
		SkillName:     skill,
		Initiative:    rand.Intn(8) + 3,
	}
}

func initGoblin() *Monster {
	return &Monster{
		Name:          "Gobelin d'entraînement",
		MaxHealth:     80,
		CurrentHealth: 80,
		Damage:        15,
		SkillName:     "Coup de poing",
		Initiative:    rand.Intn(5) + 1,
	}
}
