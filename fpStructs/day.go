package fpStructs


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

func (fd *FoodDay)countFoodNutrients(nMap *map[string]*Nutrient,foodMass float32){
	for _,n := range *nMap{
		fd.DayRDA.UpdateNutrient(n.Name,n.NType,((n.GetDensityGrams()*foodMass)))
	}
}

func (fd *FoodDay) CountNutrientsRDA() {
	for _,f := range fd.Foods{
		fd.countFoodNutrients(&f.Nutrients.Vitamins,f.GetMass())
		fd.countFoodNutrients(&f.Nutrients.Minerals,f.GetMass())
		fd.countFoodNutrients(&f.Nutrients.AminoAcids,f.GetMass())
		fd.countFoodNutrients(&f.Nutrients.Macros,f.GetMass())
	}
}
