package main

import "fmt"

//Créer une structure correspondant à une voiture définie par les champs suivants :

type Car struct {
	Brand string
	Name  string
	Power int
	Color string
	Trunk []Object
}

type Object struct {
	Name     string
	Quantity int
}

func (car Car) displayInfo() {
	fmt.Println("=== Information du véhicule ===")
	fmt.Printf("\tModèle : %s\n ", car.Name)
	fmt.Printf("\tMarque : %s\n", car.Brand)
	fmt.Printf("\tPuissance : %d\n", car.Power)
	fmt.Printf("\tCouleur : %s\n", car.Color)
	fmt.Printf("=== Coffre du véhicule ===")
	if len(car.Trunk) == 0 {
		fmt.Println("\t Aucun objet...")
		return
	}
	for _, item := range car.Trunk {
		fmt.Printf("\t-\n %s x%d\n", item.Name, item.Quantity)
	}
}

//Créer une méthode qui permet de changer la couleur du véhicule, en prenant en paramètre
//une chaîne de caractères correspondant à la nouvelle couleur. Attention, pensée à afficher un
//message d’erreur si le véhicule est de la même couleur sinon indiquée le changement de
//couleur.

func (voiture *Car) newColor(newColor string) {
	if voiture.Color == newColor {
		fmt.Println("Oups changement impossible... Couleur identique")
		return
	}
	voiture.Color = newColor
	fmt.Println("La couleur du véhicule à changé : %s\n", voiture.Color)
}

//Écrivez une méthode qui permet d’ajouter un objet dans le coffre du véhicule. Attention si un
//objet identique (du même nom) est présent uniquement la quantité sera modifiée, il ne doit
//pas avoir d’objet en doublon. Si l’objet n’est pas présent il sera alors ajouté au slice (le coffre).

func (car *Car) addObjectTrunk(obj Object) {
	for index := range car.Trunk {
		if car.Trunk[index].Name == obj.Name {
			car.Trunk[index].Quantity += obj.Quantity
			fmt.Printf("Quantité modifiée : %s +%d", obj.Name, obj.Quantity)
			return
		}
	}
	car.Trunk = append(car.Trunk, obj)
	fmt.Printf("Nouvelle item ajouté dans le coffre : %s x%d\n", obj.Name, obj.Quantity)
}

//Initialiser votre voiture ainsi que quelques objets par exemple sac de course, triangle de
//signalisation…
func main() {
	car01 := Car{"Ynov", "Ynfo", 100, "red", []Object{}}
	fmt.Println(car01)
	obj01 := Object{"ordinateur portable", 1}
	obj02 := Object{"sac de course", 1}
	car01.Trunk = append(car01.Trunk, obj02)
	car01.displayInfo()
	car01.newColor("Red")
	car01.displayInfo()
	car01.addObjectTrunk(obj01)
}
