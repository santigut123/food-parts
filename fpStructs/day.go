package fpStructs

import (
)

type FoodDay struct{
	Name string
	Foods []Food
	DayRDA  RDA
	RDAStandard RDA
}

func NewFoodDay (name string,rda RDA) *FoodDay{
	nfd := FoodDay{
		Name:        name,
		Foods:       []Food{},
		DayRDA:      RDA{},
		RDAStandard: rda,
	}
	return &nfd
}

func (fd *FoodDay) AddFood(newFood Food){
	fd.Foods= append(fd.Foods,newFood)
}
func (fd *FoodDay) AddFoodNutrients(food Food){
	for k,v := range food.Nutrients.Macros{
		fd.DayRDA.SetNutrient()
	}
	for k,v := range food.Nutrients.Vitamins{

	}
	for k,v := range food.Nutrients.Minerals{

	}
	for k,v := range food.Nutrients.AminoAcids{

	}

}
func (fd *FoodDay) CountNutrients() {

}
