package db_reader

// In this file we read the USDA food database csv files and transfer them to a food_db object. After we save the structure directly to a file for easy retrival.
import (
	"encoding/csv"
	"io"
	"strconv"

	//	"encoding/gob"
	"fmt"
	//"io/ioutil"
	"log"
	"os"

	"github.com/santigut123/food-parts/fpStructs"
)
func GetFoodNames(reader *csv.Reader) fpStructs.FoodDB{
	foodDB:=fpStructs.FoodDB{}
	foodDB.SetName("USDA")
	itemsRead := 0
	// go line by line in csv and then add new food with name and id
	for{
		line,err:=reader.Read()
		if err==io.EOF{
			break
		}
		if err!=nil{
			log.Panicf("Problems reading lines in csv",err)
		}

		fmt.Println(line[2])
		id,_ := strconv.Atoi(line[0])
		name := line[2]
				newFood := fpStructs.Food{}
		newFood.SetID(id)
		newFood.SetName(name);
		itemsRead++
	}
	fmt.Println("Items Read %d",itemsRead)
	return foodDB
}

func ReadDatabase() fpStructs.FoodDB{
	// gets the directory the program is in and then finds the Database
	pwd,_:=os.Getwd()
	foodNameFilePath := pwd+"/food_data/food.csv"
	foodNameFile,err:=os.Open(foodNameFilePath)
	if err!=nil{
		log.Panicf("Failed reading data from file %s",err)
	}
	csvReader:= csv.NewReader(foodNameFile)
	// Skip first line in csv which has column names
	if _, err := csvReader.Read(); err != nil {
		panic(err)
	}
	// This gets the food name and IDs in the FoodDB
	foodDatabase:=GetFoodNames(csvReader)
	foodMap:=foodDatabase.GetFoods()
	for k,v :=range *foodMap{
		fmt.Printf("key is %d name is %s \n",k,v.GetName())
	}
	return foodDatabase
}
