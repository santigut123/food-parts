package foodParser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/santigut123/food-parts/fpStructs"
)

type FoodParser struct {
	fileName string
	fileType string
	db *fpStructs.FoodDB
}
func NewFoodParser(fileName string, fileType string,db *fpStructs.FoodDB) *FoodParser{
	nfp:=FoodParser{
		fileName: fileName,
		fileType: fileType,
		db:       db,
	}
	return &nfp
}
func fillDayRDA(fd *fpStructs.FoodDay){
	// for each food in the day
	for _,food := range fd.Foods{
		// go through each nutrient type and update the nutrient in the day rda
		for _,nutri :=range food.Nutrients.Vitamins{
			fd.DayRDA.UpdateNutrient(nutri.Name, nutri.NType, ((nutri.Mass/nutri.Volume)*food.GetMass()))
		}
		for _,nutri :=range food.Nutrients.Minerals{
			fd.DayRDA.UpdateNutrient(nutri.Name, nutri.NType, ((nutri.Mass/nutri.Volume)*food.GetMass()))
		}
		for _,nutri :=range food.Nutrients.AminoAcids{
			fd.DayRDA.UpdateNutrient(nutri.Name, nutri.NType, ((nutri.Mass/nutri.Volume)*food.GetMass()))
		}
		for _,nutri :=range food.Nutrients.Macros{
			fd.DayRDA.UpdateNutrient(nutri.Name, nutri.NType, ((nutri.Mass/nutri.Volume)*food.GetMass()))
		}
	}


}
func(fp *FoodParser) ReadFoodDayFile(dayname string) *fpStructs.FoodDay{
	// these arrays can be optimized so they dont have to realloc everytime they get appended
	fdFile,err := os.Open(fp.fileName)
	if err != nil{
		log.Println("Did not find your Food Day File.")
		panic(err)
	}

	s := bufio.NewScanner(fdFile)
	newFd := fpStructs.FoodDay{
		Name:        dayname,
		Foods:       []fpStructs.Food{},
		DayRDA:      fpStructs.RDA{
			RDANutri: fpStructs.Nutrients{
				Vitamins:   map[string]*fpStructs.Nutrient{},
				Minerals:   map[string]*fpStructs.Nutrient{},
				AminoAcids: map[string]*fpStructs.Nutrient{},
				Macros:     map[string]*fpStructs.Nutrient{},
			},
		},
		RDAStandard: *GetReferenceRDA(),
	}

	var id int
	var mass float32
	var units string
	// current food being processed
	var curFood fpStructs.Food
	for s.Scan(){
		// This Checks for comments
		if strings.Contains(s.Text(),"//"){
			continue
		}
		if s.Text()=="\n"{
			continue
		}
		fmt.Sscanf(s.Text(),"%d %f%s",&id,&mass,&units)
		curFood=*fp.db.GetFood(id)
		curFood.SetMass(mass, units)
		newFd.AddFood(curFood)
	}
	return &newFd
}
