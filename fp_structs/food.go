package fp_structs

import ("strings")

// Food, composed of nutrients and has a name
// Nurtients the basic sub-unit of food
// RDA, recommended Daily Allowance, how much of something you should eat a day
// Food, made of a mix of nutrients and has a mass and a name
type Food struct{
	name string
	mass float32
	vitamins map[string]Nutrient
	minerals map[string]Nutrient
	amino_acids map[string]Nutrient
	macros map[string]Nutrient
}

func(f *Food) GetName() string{
	return f.name
}

func(f *Food) GetMass() float32{
	return f.mass
}

func(f *Food) addVitamin(n Nutrient) {
	f.vitamins[n.name]=n
}

func(f *Food) addMineral(n Nutrient) {
	f.minerals[n.name]=n
}

func(f *Food) addAminoAcid(n Nutrient) {
	f.amino_acids[n.name]=n
}

func(f *Food) addMacro(n Nutrient) {
	f.macros[n.name]=n
}
// adds nutrient to the correct map
func(f *Food) AddNutrient(n Nutrient){
	var nType rune = n.nType
	var name string =n.name
	if(nType=='v'){
		f.vitamins[name]=n
	}else if(nType=='m'){
		f.minerals[name]=n
	}else if(nType=='a'){
		f.amino_acids[name]=n;
	}else{
		f.macros[name]=n;
	}
}
// Looks through all maps and returns nutrient pointer
func(f *Food) GetNutrient(name string) *Nutrient{
	name=strings.ToLower(name)
	if val,present :=f.macros[name];present{
		return &val;
	}else if val,present :=f.amino_acids[name];present{
		return &val;
	}else if val,present :=f.minerals[name];present{
		return &val;
	}else if val,present :=f.vitamins[name];present{
		return &val;
	}
	return nil
}
