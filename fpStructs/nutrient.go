package fpStructs

import ("strings")

type Nutrient struct{
	Mass float32
	Volume float32
	Name string
	Units string
	NType rune
}

func(n *Nutrient) GetMass() float32{
	return n.Mass
}

func(n *Nutrient) GetVolume() float32{
	return n.Volume
}

func(n *Nutrient) GetName() string{
	return n.Name
}

func(n *Nutrient) SetName(name string) {
	n.Name=strings.ToLower(name);
}
func(n *Nutrient) GetTypeString() string{
	var nType rune = n.NType
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
	return (n.Mass/n.Volume)
}
