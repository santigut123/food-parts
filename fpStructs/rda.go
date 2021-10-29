package fpStructs

type RDA struct{
	Vitamins map[string]float32
	Minerals map[string]float32
	Amino_acids map[string]float32
	Macros map[string]float32
}
// The following 3 functions are for changing macros
// usually these are the things that need the most variation in a person's diet.
func(r *RDA) SetMacroProtein(grams int){
	r.Macros["protein"]=float32(grams)
}

func(r *RDA) SetMarcoFat(grams int){
	r.Macros["protein"]=float32(grams)
}

func(r *RDA) SetMacroCarbs(grams int){
	r.Macros["protein"]=float32(grams)
}
