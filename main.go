package main

import (
	//"os"
	//	"encoding/gob"

	//		"github.com/gdamore/tcell/v2"
	//	"github.com/rivo/tview"
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"time"
	"github.com/santigut123/food-parts/db_reader"
	"github.com/santigut123/food-parts/fpStructs"
)
func main(){
	start :=time.Now()
	foodDB :=db_reader.ReadDatabase()
	foodMap := foodDB.GetFoods()
	var foods int
	for k,_ :=range *foodMap{
		foods+=k
	}
	file,_:=os.ReadFile("foodDB.gob")
	buf:=bytes.NewBuffer(file)
	gobDecoder:=gob.NewDecoder(buf)
	foodDBGob:= fpStructs.NewFoodDB("USDA")
	gobMap :=foodDBGob.GetFoods()
	gobDecoder.Decode(foodDBGob)
	for _,v :=range *gobMap{
		foods++
		fmt.Println("THE NAME OF THE FOOD",v.GetName())
	}
	elapsed:=time.Since(start)
	fmt.Println("reading ",foods," foods "," took ",elapsed)

}
