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
func(rda *RDA) UpdateNutrient(nName string,nType rune, nMass float32) {
	nutrients:= &rda.RDANutri

	// Must change out specific references to vitamins/minerals
	switch nType{
		case 'v':
			if val, present:=nutrients.Vitamins[n.Name];present{
				val.Mass += nMass
			}else{
				newNutri := Nutrient{
					Mass:   nMass,
					Volume: 0,
					Name:   nName,
					Units:  "G",
					NType:  nType,
				}
				nutrients.Vitamins[nName]=&newNutri
			}
		case 'm':
			if val, present:=nutrients.Vitamins[n.Name];present{
				val.Mass+=nMass
			}else{
				newNutri := Nutrient{
					Mass:   nMass,
					Volume: 0,
					Name:   nName,
					Units:  "G",
					NType:  nType,
				}
				nutrients.Vitamins[nName]=&newNutri
			}

		case 'a':
			if val, present:=nutrients.Vitamins[n.Name];present{
				val.Mass+=nMass
			}else{
				newNutri := Nutrient{
					Mass:   nMass,
					Volume: 0,
					Name:   nName,
					Units:  "G",
					NType:  nType,
				}
				nutrients.Vitamins[nName]=&newNutri
			}

		case 'M':
			if val, present:=nutrients.Vitamins[n.Name];present{
				val.Mass+=nMass
			}else{
				newNutri := Nutrient{
					Mass:   nMass,
					Volume: 0,
					Name:   nName,
					Units:  "G",
					NType:  nType,
				}
				nutrients.Vitamins[nName]=&newNutri
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
