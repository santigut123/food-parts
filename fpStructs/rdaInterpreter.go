package fpStructs

import (
)

type RDAInterpreter struct{
	standard RDA
	actual RDA
}
func MakeRDAInterpreter(actual RDA, standard RDA) *RDAInterpreter {
	newRDAI := RDAInterpreter{
		standard: standard,
		actual:   actual,
	}
	return &newRDAI
}
func(RDAI *RDAInterpreter) GetNutriPercentages() map[string]int{



}
