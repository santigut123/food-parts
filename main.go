package main

import (
	"fmt"
	//"os"
	//	"encoding/gob"

	//		"github.com/gdamore/tcell/v2"
	//	"github.com/rivo/tview"
	"github.com/santigut123/food-parts/db_reader"
)
func main(){
	foodDB :=db_reader.ReadDatabase()
	for k,v :=range *foodDB.GetFoods(){
		fmt.Printf("key is %d name is %s \n",k,v.GetName())
	}


}
