package fpStructs

import ("strings")

// Food, composed of Nutrients and has a name
// Nurtients the basic sub-unit of food
// RDA, recommended Daily Allowance, how much of something you should eat a day
// Food, made of a mix of Nutrients and has a mass and a name
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
type Nutrients struct{
	Vitamins map[string]Nutrient
	Minerals map[string]Nutrient
	Amino_acids map[string]Nutrient
	Macros map[string]Nutrient
}
func(ns *Nutrients) addVitamin(n Nutrient) {
	ns.Vitamins[n.Name]=n
}

func(ns *Nutrients) addMineral(n Nutrient) {
	ns.Minerals[n.Name]=n
}

func(ns *Nutrients) addAminoAcid(n Nutrient) {
	ns.Amino_acids[n.Name]=n
}
func(ns *Nutrients) PrintNutrients(){
	for _,v:=range ns.Vitamins{
		v.PrintNutrient()
	}
	for _,v:=range ns.Minerals{
		v.PrintNutrient()
	}
	for _,v:=range ns.Amino_acids{
		v.PrintNutrient()
	}
	for _,v:=range ns.Macros{
		v.PrintNutrient()
	}

}

type Food struct{
	FoodID FoodID
	Macros Macros
	Nutrients Nutrients
}
// adds nutrient to the correct map
func(f *Food) AddNutrient(n Nutrient){
	var nType rune = n.NType
	var name string =n.Name
	if(nType=='v'){
		f.Nutrients.Vitamins[name]=n
	}else if(nType=='m'){
		f.Nutrients.Minerals[name]=n
	}else if(nType=='a'){
		f.Nutrients.Amino_acids[name]=n;
	}else{
		f.Nutrients.Macros[name]=n;
	}
}
func (f *Food) PrintNutrients(){
	f.Nutrients.PrintNutrients()

}

func NewFood(foodID FoodID,macros Macros) *Food{
	newFood := Food{
		FoodID:    foodID,
		Macros:    macros,
		Nutrients: Nutrients{
			Vitamins:    make(map[string]Nutrient),
			Minerals:    make(map[string]Nutrient),
			Amino_acids: make(map[string]Nutrient),
			Macros:      make(map[string]Nutrient),
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
	if val,present :=f.Nutrients.Macros[name];present{
		return &val;
	}else if val,present :=f.Nutrients.Amino_acids[name];present{
		return &val;
	}else if val,present :=f.Nutrients.Minerals[name];present{
		return &val;
	}else if val,present :=f.Nutrients.Vitamins[name];present{
		return &val;
	}
	// if the program doesn't find the nutrient, return null
	return nil
}
