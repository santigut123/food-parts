package fpStructs


type FoodDay struct{
	Name string
	Foods []Food
	DayRDA  RDA
}
type FoodHolder interface{
	AddFood(curFood Food)
}
func NewFoodDay (name string) *FoodDay{
	nfd := FoodDay{
		Name:        name,
		Foods:       []Food{},
		DayRDA:      RDA{},
	}
	return &nfd
}

func (fd *FoodDay) AddFood(newFood Food){
	fd.Foods= append(fd.Foods,newFood)
}
func(fd *FoodDay) AddFoods(newFoods []Food){
	fd.Foods = append(fd.Foods,newFoods...)
}

func (fd *FoodDay)countFoodNutrients(nMap *map[string]*Nutrient,foodMass float32){
	for _,n := range *nMap{
		fd.DayRDA.UpdateNutrient(n.Name,n.NType,((n.GetDensityGrams()*foodMass)))
	}
}

func (fd *FoodDay) CountNutrients() {
	for _,f := range fd.Foods{
		fd.countFoodNutrients(&f.Nutrients.Vitamins,f.GetMass())
		fd.countFoodNutrients(&f.Nutrients.Minerals,f.GetMass())
		fd.countFoodNutrients(&f.Nutrients.AminoAcids,f.GetMass())
		fd.countFoodNutrients(&f.Nutrients.Macros,f.GetMass())
	}
}
