package fpStructs

type RDA struct{
	Vitamins map[string]float32
	Minerals map[string]float32
	AminoAcids map[string]float32
	Macros map[string]float32
}
// The following 3 functions are for changing macros
// usually these are the things that need the most variation in a person's diet.
func(r *RDA) SetMacroProtein(weight float32){
	r.Macros["protein"]=weight
}
func(r *RDA) SetMarcoFat(weight float32){
	r.Macros["fat"]=weight
}
func(r *RDA) SetMacroCarbs(weight float32){
	r.Macros["carbs"]=weight
}
func(r *RDA) AddNutrient(nType string, n Nutrient) {
	if nType=="v"{
		_,present:=r.Vitamins[n.Name]
		if !present{
			r.Vitamins[n.Name]=n.Mass
		} else{
			r.Vitamins[n.Name]+=n.Mass
		}
	} else if nType=="m"{
		_,present:=r.Minerals[n.Name]
		if !present{
			r.Minerals[n.Name]=n.Mass
		} else{
			r.Minerals[n.Name]+=n.Mass
		}
	} else if nType=="a"{
		_,present:=r.AminoAcids[n.Name]
		if !present{
			r.AminoAcids[n.Name]=n.Mass
		} else{
			r.AminoAcids[n.Name]+=n.Mass
		}
	} else{
		_,present:=r.Macros[n.Name]
		if !present{
			r.Macros[n.Name]=n.Mass
		} else{
			r.Macros[n.Name]+=n.Mass
		}
	}
}
