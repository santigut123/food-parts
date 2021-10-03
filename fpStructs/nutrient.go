package fpStructs

import ("strings")

type Nutrient struct{
	mass float32
	volume float32
	name string
	nType rune
}

func(n *Nutrient) GetMass() float32{
	return n.mass
}

func(n *Nutrient) GetVolume() float32{
	return n.volume
}

func(n *Nutrient) GetName() string{
	return n.name
}

func(n *Nutrient) SetName(name string) {
	n.name=strings.ToLower(name);
}

func(n *Nutrient) GetTypeString() string{
	var nType rune = n.nType
	var typeString string;
	if(nType == 'v'){
		typeString= "Vitamin"
	}else if(nType =='m'){
		typeString = "Mineral"
	}else if(nType =='a'){
		typeString = "Amino Acid"
	}else{
		typeString = "Macro"
	}
	return typeString
}

func(n *Nutrient) getDensity () float32{
	return (n.mass/n.volume)
}
