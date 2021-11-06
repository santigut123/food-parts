package fpStructs

import "fmt"


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
func iterateNutriPercentages(category string,actual *RDA,ref map[string]*Nutrient){
	fmt.Printf("%s: ",category)
	var percent float32
	for k,referenceVal := range ref{
		if actualVal,present := actual.RDANutri.Vitamins[k];present{
			percent = actualVal.Mass /referenceVal.Mass
		}else{
			percent=0;
		}
		fmt.Printf("%s: %f.3 ",k,percent)
	}
}
func(rda *RDA) PrintRDAPercentages (actual *RDA){
	iterateNutriPercentages("Vitamins",actual,rda.RDANutri.Vitamins)
	iterateNutriPercentages("Minerals",actual,rda.RDANutri.Minerals)
	iterateNutriPercentages("Amino Acids",actual,rda.RDANutri.AminoAcids)
	iterateNutriPercentages("Macros",actual,rda.RDANutri.Macros)
}
