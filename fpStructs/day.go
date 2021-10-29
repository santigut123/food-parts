package fpStructs

type NutrientCounter struct{
	rda *RDA
	day *Day
}
type Day struct{
	foods []Food
	int Protein
}
func (d *Day) GetNutrientAmount()
