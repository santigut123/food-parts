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
func (fp *FoodParser) readFoodFile(fh fpStructs.FoodHolder){
	s:=fp.MakeScanner()
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
		// checks for empty lines
		if s.Text()=="\n"{
			continue
		}
		_,err :=fmt.Sscanf(s.Text(),"%d %f%s",&id,&mass,&units)
		if err !=nil{
			fmt.Println("Error reading your file, check your formatting")
			panic(err)
		}
		curFood=*fp.db.GetFood(id)
		curFood.SetMass(mass, units)
		fh.AddFood(curFood)
	}
}
func (fp *FoodParser) readFoodDay(fd *fpStructs.FoodDay){
	s:=fp.MakeScanner()
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
		// checks for empty lines
		if s.Text()=="\n"{
			continue
		}
		if strings.Contains(s.Text(),"recipe"){
			var recipeName string
			_,err:= fmt.Sscanf(s.Text(), "recipe %s",recipeName)
			if err!= nil{
				fmt.Println("Problem reading your file, check your recipe formatting")
				panic(err)
			}
			fd.AddFoods(fp.db.GetRecipeFoods(recipeName))
		}
		_,err :=fmt.Sscanf(s.Text(),"%d %f%s",&id,&mass,&units)
		if err !=nil{
			fmt.Println("Error reading your file, check your formatting")
			panic(err)
		}
		curFood=*fp.db.GetFood(id)
		curFood.SetMass(mass, units)
		fd.AddFood(curFood)
}
}
func (fp *FoodParser) MakeScanner() *bufio.Scanner{
	fdFile,err := os.Open(fp.fileName)
	if err != nil{
		log.Println("Did not find your Food Day File.")
		panic(err)
	}

	s := bufio.NewScanner(fdFile)
	return s

}
func(fp *FoodParser) ReadRecipeFile() fpStructs.Recipe{
	newRecip := fpStructs.Recipe{}
	fp.readFoodFile(&newRecip)
	return newRecip
}
func(fp *FoodParser) ReadFoodDayFile(dayName string) *fpStructs.FoodDay{
	// these arrays can be optimized so they dont have to realloc everytime they get appended
		newFd := fpStructs.FoodDay{
		Name:        dayName,
		Foods:       []fpStructs.Food{},
		DayRDA:      fpStructs.RDA{
			RDANutri: fpStructs.Nutrients{
				Vitamins:   map[string]*fpStructs.Nutrient{},
				Minerals:   map[string]*fpStructs.Nutrient{},
				AminoAcids: map[string]*fpStructs.Nutrient{},
				Macros:     map[string]*fpStructs.Nutrient{},
			},
		},
	}
	fp.readFoodFile(&newFd)
	return &newFd
}
