package fpStructs

type RDA struct{
	vitamins map[string]float32
	minerals map[string]float32
	amino_acids map[string]float32
	macros map[string]float32
}
// The following 3 functions are for changing macros
// usually these are the things that need the most variation in a person's diet.
func(r *RDA) SetMacroProtein(grams int){
	r.macros["protein"]=float32(grams)
}

func(r *RDA) SetMarcoFat(grams int){
	r.macros["protein"]=float32(grams)
}

func(r *RDA) SetMacroCarbs(grams int){
	r.macros["protein"]=float32(grams)
}
