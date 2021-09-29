package food_components

import (
	"fmt"
	"hash/maphash"
)

// Everything information is stored in grams
type Nutrient struct{
	mass float32
	volume float32
	name string
	nType string
}
func(n *Nutrient) getDensity (density float32){
	density=(n.mass/n.volume)
}

type RDAInfo struct{
	age int
	weight int
	sex rune
	rdi map[string]float32
}
type
func(r *RDAInfo) initializeProtein(){

}
func(r *RDAInfo) InitializeRdi(){


}

// Adds nutrient to nutritional profile
func(np *NutritionalProfile) AddNurient(n Nutrient){


}
