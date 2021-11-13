package fpStructs

import (
	"fmt"
	"strings"
)

type Nutrient struct{
	Mass float32
	Volume float32
	Name string
	Units string
	NType rune
}
func(n *Nutrient) PrintNutrient(){
	fmt.Println(n.Name," ",n.Mass,n.Units," Type:",n.NType)
}
func(n *Nutrient) GetMass() float32{
	return n.Mass
}
func(n *Nutrient) ConvertToGrams() float32{
	if(n.Units=="G"){
		return n.Mass
	} else if (n.Units=="MG"){
		return n.Mass/1000
	} else if (n.Units=="UG"){
		return n.Mass/1000000
	}
	return n.Mass
}
func(n *Nutrient) GetDensityGrams() float32{
	return n.ConvertToGrams()/n.Volume

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
