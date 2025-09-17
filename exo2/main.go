package main

import "fmt"

type Waste struct {
	Name     string
	Type     string
	Quantity int
}

type Dustbin struct {
	Name     string
	Type     string
	Contents []Waste
}

func (bin Dustbin) displayInfo() {
	fmt.Printf("=== %s ===\n", bin.Name)
	fmt.Printf("Contient des décchets de type %s :\n", bin.Type)
	if len(bin.Contents) < 1 {
		fmt.Println("\tpoublle vide\n")
	} else {
		for _, item := range bin.Contents {
			fmt.Printf("\t- %s x %d\n", item.Name, item.Quantity)
		}
	}
}

func (bin *Dustbin) addWaste(waste Waste) {
	for index := range bin.Contents {
		if bin.Contents[index].Name == waste.Name {
			bin.Contents[index].Quantity += waste.Quantity
			fmt.Println("Quantité modifiée")
			return
		}
	}
	bin.Contents = append(bin.Contents, waste)
	fmt.Println("ajout du nouveau déchet")

}

func main() {
	dustbin01 := Dustbin{"jaune", "plastic", []Waste{}}
	fmt.Println(dustbin01)
	dustbin02 := Dustbin{"vert", "verre", []Waste{}}
	fmt.Println(dustbin02)
	dustbin03 := Dustbin{"blanc", "papier", []Waste{}}
	fmt.Println(dustbin03)
	dustbin04 := Dustbin{"noir", "autre", []Waste{}}
	fmt.Println(dustbin04)
	dustbin01.addWaste(Waste{"bouteille", "plastique", 2})
	dustbin01.displayInfo()
	dustbin02.displayInfo()
	dustbin03.displayInfo()
	dustbin04.displayInfo()
}
