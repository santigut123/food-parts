package fpStructs


type RDA struct{
	RDANutri Nutrients
}

func(rda *RDA) MakeRDA() *RDA{
	newRDA := RDA{
		RDANutri: Nutrients{
			Vitamins:   map[string]*Nutrient{},
			Minerals:   map[string]*Nutrient{},
			AminoAcids: map[string]*Nutrient{},
			Macros:     map[string]*Nutrient{},
		},
	}
	return &newRDA
}
func(rda *RDA) UpdateNutrient(n Nutrient) {
	nutrients:= &rda.RDANutri

	switch n.NType{
		case 'v':
			if val, present:=nutrients.Vitamins[n.Name];present{
				val.Mass+=n.Mass
			}else{
				newNutri := n
				nutrients.Vitamins[n.Name]=&newNutri
			}
		case 'm':
			if val, present:=nutrients.Vitamins[n.Name];present{
				val.Mass+=n.Mass
			}else{
				newNutri := n
				nutrients.Vitamins[n.Name]=&newNutri
			}

		case 'a':
			if val, present:=nutrients.Vitamins[n.Name];present{
				val.Mass+=n.Mass
			}else{
				newNutri := n
				nutrients.Vitamins[n.Name]=&newNutri
			}

		case 'M':
			if val, present:=nutrients.Vitamins[n.Name];present{
				val.Mass+=n.Mass
			}else{
				newNutri := n
				nutrients.Vitamins[n.Name]=&newNutri
			}
	}
}
func (rda *RDA) printNutriPercentages(actual *RDA){

}
func(rda *RDA) PrintRDAPercentages (actual *RDA){

	for k,v := range rda.RDANutri.Vitamins{


	}
	for k,v := range rda.RDANutri.Minerals{

	}
	for k,v := range rda.RDANutri.AminoAcids{

	}
	for k,v := range rda.RDANutri.Macros{

	}
}
