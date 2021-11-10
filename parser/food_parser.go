package parser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/santigut123/food-parts/fpStructs"
)

type FileParser struct {
	filename string
	fileType string
	db *fpStructs.FoodDB
}
func NewFileParser(fileName string, fileType string,db *fpStructs.FoodDB) *FileParser{
	nfp:=FileParser{
		filename: fileName,
		fileType: fileType,
		db:       db,
	}
	return &nfp
}
func(fp *FileParser) ReadFoodDayFile(dayname string) *fpStructs.FoodDay{
	// these arrays can be optimized so they dont have to realloc everytime they get appended
	fdFile,err := os.Open(fp.filename)
	s := bufio.NewScanner(fdFile)
	if err != nil{
		log.Println("Did not find your Food Day File.")
		panic(err)
	}
	newFd := fpStructs.FoodDay{
		Name:        dayname,
		Foods:       []fpStructs.Food{},
		DayRDA:      fpStructs.RDA{},
		RDAStandard: fpStructs.RDA{},
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
		fmt.Sscanf(s.Text(),"%d %d%s",&id,&mass,&units)
		curFood=*fp.db.GetFood(id)
		newFood:=fpStructs.Food{
			FoodID:    curFood.FoodID,
			Macros:    curFood.Macros,
			Nutrients: curFood.Nutrients,
		}
		newFood.SetMass(mass, units)
	}
	return &newFd
}
