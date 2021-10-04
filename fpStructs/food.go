package fpStructs

import ("strings")

// Food, composed of nutrients and has a name
// Nurtients the basic sub-unit of food
// RDA, recommended Daily Allowance, how much of something you should eat a day
// Food, made of a mix of nutrients and has a mass and a name
type FoodID struct{
	Id int
	Name string
	Mass float32
}
type Macros struct{
	Protein float32
	Fat float32
	Carbs float32
}
type nutrients struct{
	vitamins map[string]Nutrient
	minerals map[string]Nutrient
	amino_acids map[string]Nutrient
	macros map[string]Nutrient
}
func(ns *nutrients) addVitamin(n Nutrient) {
	ns.vitamins[n.name]=n
}

func(ns *nutrients) addMineral(n Nutrient) {
	ns.minerals[n.name]=n
}

func(ns *nutrients) addAminoAcid(n Nutrient) {
	ns.amino_acids[n.name]=n
}

// adds nutrient to the correct map
func(ns *nutrients) addNutrient(n Nutrient){
	var nType rune = n.nType
	var name string =n.name
	if(nType=='v'){
		ns.vitamins[name]=n
	}else if(nType=='m'){
		ns.minerals[name]=n
	}else if(nType=='a'){
		ns.amino_acids[name]=n;
	}else{
		ns.macros[name]=n;
	}
}

type Food struct{
	FoodID FoodID
	Macros Macros
	nutrients nutrients
}
func NewFood(foodID FoodID,macros Macros) *Food{
	newFood := Food{
		FoodID:    foodID,
		Macros:    macros,
		nutrients: nutrients{
			vitamins:    make(map[string]Nutrient),
			minerals:    make(map[string]Nutrient),
			amino_acids: make(map[string]Nutrient),
			macros:      make(map[string]Nutrient),
		},
	}

	return &newFood

}
func(f *Food) GetID() int{
	return f.FoodID.Id
}
func(f *Food) GetName() string{
	return f.FoodID.Name
}
func(f *Food) ChangeName(name string) {
	f.FoodID.Name = name
}
func(f *Food) GetMass() float32{
	return f.FoodID.Mass
}
func(f *Food) SetMass(mass float32) {
	f.FoodID.Mass=mass
}
// Looks through all maps and returns nutrient pointer
func(f *Food) GetNutrient(name string) *Nutrient{
	name=strings.ToLower(name)
	if val,present :=f.nutrients.macros[name];present{
		return &val;
	}else if val,present :=f.nutrients.amino_acids[name];present{
		return &val;
	}else if val,present :=f.nutrients.minerals[name];present{
		return &val;
	}else if val,present :=f.nutrients.vitamins[name];present{
		return &val;
	}
	// if the program doesn't find the nutrient, return null
	return nil
}
