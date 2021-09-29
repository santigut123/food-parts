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
	vitamins map[string]float32
	minerals map[string]float32
	amino_acids map[string]float32
	macros map[string]float32
}
// The following 4 functions are for changing macros
// usually these are the things that need the most variation in a person's diet.
func(r *RDAInfo) initializeProtein(grams int){
	r.macros["protein"]=float32(grams)
}
// Adds nutrient to nutritional profile
func(np *NutritionalProfile) AddNurient(n Nutrient){


}
