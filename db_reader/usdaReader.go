package db_reader

// In this file we read the USDA food database csv files and transfer them to a food_db object. After we save the structure directly to a file for easy retrival.
import (
	"encoding/csv"
	"io"
	"strconv"
	//	"encoding/gob"
	//"io/ioutil"
	"log"
	"os"
	"github.com/santigut123/food-parts/fpStructs"
)
const(
	nutrientID=iota
	nutrientName=iota
	nutrientUnit=iota
)
type DBNutrient struct{
	Unit float32
	Name string
}
func NewDBNutrient(name string,unit string) DBNutrient{
	var unitInGrams float32
	if(unit=="G"){
		unitInGrams=1
	}else if(unit=="MG"){
		unitInGrams=.001
	}else{
		unitInGrams=.000001
	}
	newDBNutrient:=DBNutrient{
		Unit: unitInGrams,
		Name: name,
	}
	return newDBNutrient

}
func GetFileLoc(relDirectory string) string{
	// gets the directory the program is in and then finds the Database
	pwd,_:=os.Getwd()
	filePath := pwd+relDirectory
	return filePath
}
func skipFirstLine(reader *csv.Reader){
	// Skip first line in csv which has column names
	if _, err := reader.Read(); err != nil {
		panic(err)
	}
}
// Look through USDA food name file which gets us the foods name and its ID
func addFoods(db *fpStructs.FoodDB) {
	// the standard measurement in the usda database is 100 grams
	const usdaWeightGrams=100
	// find the food names and id's file
	foodNameFilePath := GetFileLoc("/food_data/food.csv")
	foodNameFile,err := os.Open(foodNameFilePath)
	if err!=nil{
		log.Panicf("Failed reading data from food name file %s",err)
	}
	reader:= csv.NewReader(foodNameFile)
	skipFirstLine(reader)
	db.SetName("USDA")
	// temp variables to use below
	var id int
	var name string
	// go line by line in csv and then add new food with name and id

	for{
		line,err:=reader.Read()
		if err==io.EOF{
			break
		}
		if err!=nil{
			log.Panicf("Problems reading lines in csv",err)
		}
		id,_ = strconv.Atoi(line[0])
		name = line[2]
		// intialize food struct with the foodID and blank macros
		// ( we wil nutrients from our getNutrient function )
		newFood := fpStructs.NewFood(fpStructs.FoodID{
			Id:   id,
			Name: name,
			Mass: usdaWeightGrams,
		},
		fpStructs.Macros{
			Protein: 0,
			Fat:     0,
			Carbs:   0,
		})
		db.AddFood(*newFood)
		}
	foodNameFile.Close()
}
// Looks through nutrient file, and adds nutrients of interest
func addNutrients(foodDB *fpStructs.FoodDB,nutrientLookup map[int]DBNutrient ){
	nutrientDir := GetFileLoc("/food_data/food_nutrient.csv")
	nutrientFile,err := os.Open(nutrientDir)
	if err!=nil{
		log.Panicf("Failed reading data from nutrient file %s",err)
	}
	reader:=csv.NewReader(nutrientFile)
	skipFirstLine(reader)
	for{
		line,err:=reader.Read()
		if err==io.EOF{
			break
		}
		if err!=nil{
			log.Panicf("Problems reading lines in csv",err)
		}
	}
	nutrientFile.Close()


}
// So the csv file that has the food nutrients only has IDs so we need to make a map so we can easily
// look up the nutrients
func makeNutrientLookupMap() map[int]DBNutrient{
	nutrientDir := GetFileLoc("/food_data/food-parts/nutrient.csv")
	nutrientFile,err := os.Open(nutrientDir)
	nutrientMap := make(map[int]DBNutrient)
	if err!=nil{
		log.Panicf("Failed reading data from nutrient file %s",err)
	}
	reader:=csv.NewReader(nutrientFile)
	skipFirstLine(reader)
	var id int
	for{
		line,err:=reader.Read()
		if err==io.EOF{
			break
		}
		if err!=nil{
			log.Panicf("Problems reading lines in csv",err)
		}
		// convert id from string to int
		id,_ = strconv.Atoi(line[nutrientID])
		// add pair to map
		nutrientMap[id]=NewDBNutrient(line[nutrientName],line[nutrientID])
	}
	return nutrientMap

}
func ReadDatabase() *fpStructs.FoodDB{
		// This gets the food name and IDs in the FoodDB
	foodDatabase:=fpStructs.NewFoodDB("USDA")
	nutrientLookup:= makeNutrientLookupMap()
	addFoods(foodDatabase)
	//addNutrients(&foodDatabase)
	return foodDatabase
}
