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
func (fd *FoodDay) CountNutrients() {

}
