package db_reader

// In this file we read the USDA food database csv files and transfer them to a food_db object. After we save the structure directly to a file for easy retrival.
import (
	"encoding/csv"
	"io/ioutil"

	//"fmt"
	"io"
	"strconv"

	//	"encoding/gob"
	//"io/ioutil"
	"encoding/json"
	"log"
	"os"

	"github.com/santigut123/food-parts/foodParser"
	"github.com/santigut123/food-parts/fpStructs"
)

const (
	nutrientID   = iota
	nutrientName = iota
	nutrientUnit = iota
)

type DBNutrient struct {
	Unit    float32
	UnitStr string
	Name    string
}

func NewDBNutrient(name string, unit string) DBNutrient {
	var unitInGrams float32
	if unit == "G" {
		unitInGrams = 1
	} else if unit == "MG" {
		unitInGrams = .001
	} else {
		unitInGrams = .000001
	}
	newDBNutrient := DBNutrient{
		Unit:    unitInGrams,
		UnitStr: unit,
		Name:    name,
	}
	return newDBNutrient
}
func GetFileLoc(relDirectory string) string {
	// gets the directory the program is in and then finds the Database
	pwd, _ := os.Getwd()
	filePath := pwd + relDirectory
	return filePath
}
func skipLine(reader *csv.Reader) {
	// Skip first line in csv which has column names
	if _, err := reader.Read(); err != nil {
		panic(err)
	}
}

// Look through USDA food name file which gets us the foods name and its ID
func addFoods(db *fpStructs.FoodDB, nutrLookup map[int]DBNutrient, vitClassifier *DBVitClassifier) {

	var nutrErr error
	// the standard measurement in the usda database is 100 grams
	const usdaWeightGrams = 100
	// find the food names and id's file
	foodNameFilePath := GetFileLoc("/food_data/food.csv")
	foodNameFile, err := os.Open(foodNameFilePath)
	if err != nil {
		log.Panicf("Failed reading data from food name file %s", err)
	}
	foodReader := csv.NewReader(foodNameFile)
	skipLine(foodReader)
	// adding the nutrient file to add nutrients later
	nutrientDir := GetFileLoc("/food_data/food_nutrient.csv")
	nutrientFile, err := os.Open(nutrientDir)
	if err != nil {
		log.Panicf("Failed reading data from nutrient file %s", err)
	}
	// this file is the one that contains every nutrient spec for every food
	nutrReader := csv.NewReader(nutrientFile)
	skipLine(nutrReader)
	// the statement below might be useful to do skipline
	nutrLine, _ := nutrReader.Read()
	db.SetName("USDA")
	// temp variables to use below
	var id int
	var name string
	// go line by line in csv and then add new food with name and id

	for {
		line, err := foodReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Panicf("Problems reading lines in csv %s", err)
		}
		id, _ = strconv.Atoi(line[0])
		name = line[2]
		// intialize food struct with the foodID and blank macros
		// ( we wil nutrients from our getNutrient function )
		newFood := fpStructs.NewFood(fpStructs.FoodID{
			Id:   id,
			Name: name,
			Mass: usdaWeightGrams,
		})
		if nutrErr == nil {
			if nutrLine[1] == line[0] {
				nutrErr = addAllNutrients(nutrReader, newFood, nutrLookup, line[0], vitClassifier, &nutrLine)
			}
		}

		db.AddFood(newFood)

	}
	foodNameFile.Close()
}

// This function adds nutrient to food in foodDB, It looks at the csv until the id doesnt match
func addAllNutrients(reader *csv.Reader, food *fpStructs.Food, nutrLookup map[int]DBNutrient, foodID string, vitClassifier *DBVitClassifier, line *[]string) error {
	var mass float64
	var mass32 float32
	var volume float32 = 100
	var nutrId int
	var nutrient DBNutrient
	var nutrErr error
	// This is the amount of KJ in a KC. The measurement used by the public is KCAL so we convert it to that later
	var kjInKC float32= 4.184
	// if the id matches, add nutrient to the food
	for ; (nutrErr == nil) && ((*line)[1] == foodID); *line, nutrErr = reader.Read() {
		if nutrErr == io.EOF {
			break
		}
		if nutrErr != nil {
			log.Panicf("Problems reading lines in csv %s", nutrErr)
		}
		mass, _ = strconv.ParseFloat((*line)[3], 32)
		mass32 = float32(mass)
		nutrId, _ = strconv.Atoi((*line)[2])
		// Here is where we convert from KJ to KCAL and add to food
		if nutrId==1062{
			food.AddNutrient(&fpStructs.Nutrient{
				Mass:   mass32/kjInKC,
				Volume: volume,
				Name:   "Calories",
				Units:  "KCAL",
				NType:  'M',
			})
			continue
		}
		// Ignore nutrients that have 0 mass
		if mass32 == 0 {
			continue
		}
		nutrient = nutrLookup[nutrId]
		food.AddNutrient(&fpStructs.Nutrient{
			Mass:   float32(mass),
			Volume: volume,
			Name:   nutrient.Name,
			Units:  nutrient.UnitStr,
			NType:  vitClassifier.GetNType(nutrId),
		})
	}
	return nutrErr
}

// So the csv file that has the food nutrients only has IDs so we need to make a map so we can easily
// look up the nutrients
func makeNutrientLookupMap() map[int]DBNutrient {
	nutrientDir := GetFileLoc("/food_data/supporting_data/nutrient.csv")
	nutrientFile, err := os.Open(nutrientDir)
	nutrientMap := make(map[int]DBNutrient)
	if err != nil {
		log.Panicf("Failed reading data from nutrient file %s", err)
	}
	reader := csv.NewReader(nutrientFile)
	skipLine(reader)
	var id int
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Panicf("Problems reading lines in csv %s", err)
		}
		// convert id from string to int
		id, _ = strconv.Atoi(line[nutrientID])
		// add pair to map
		nutrientMap[id] = NewDBNutrient(line[nutrientName], line[nutrientUnit])
	}
	return nutrientMap
}
func EncodeDBToFile(db *fpStructs.FoodDB, name string) {
	file, _ := os.Create(name)
	jsonContent, _ := json.MarshalIndent(db, "", " ")
	file.Write(jsonContent)
}
func ReadDBFromFile() fpStructs.FoodDB {
	readFile, _ := ioutil.ReadFile("foodDB.json")
	newFoodDB := fpStructs.FoodDB{}
	json.Unmarshal([]byte(readFile), &newFoodDB)
	EncodeDBToFile(&newFoodDB, "newFoodDB.json")
	return newFoodDB
}
// gets recipies from saved JSON
func RetrieveRecipies() map[string]fpStructs.Recipe{

	recpFile,err:=ioutil.ReadFile(GetFileLoc("/recipes/recipes.json"))
	retrivedRecipes:= make(map[string]fpStructs.Recipe)
	if err!= nil{
		return retrivedRecipes
	}
	json.Unmarshal([]byte(recpFile), &retrivedRecipes)
	return retrivedRecipes

}
// Saves Recipes when there is a change
func SaveRecipies(rs map[string]fpStructs.Recipe) {
	recpFile,_ :=os.Create(GetFileLoc("/recipes/recipes.json"))
	jsonContent,_ := json.Marshal(rs)
	recpFile.Write(jsonContent)
}
func AddRecipe(foodDB *fpStructs.FoodDB,recipeLoc string) {
	 rp:= foodParser.NewFoodParser(recipeLoc,"Recipe",foodDB)
	 foodDB.AddRecipe(rp.ReadRecipeFile())
	 SaveRecipies(foodDB.Recipes)
}
func ReadDatabase() *fpStructs.FoodDB {
	// This gets the food name and IDs in the FoodDB
	vitClassifier := newDBVitClassifier()
	foodDatabase := fpStructs.NewFoodDB("USDA")
	foodDatabase.Recipes=RetrieveRecipies()
	// Nutrient lookup map to see the name and unit of the nutrient from its id
	nutrientLookup := makeNutrientLookupMap()
	// this adds the foods and nutrients to the foods
	addFoods(foodDatabase, nutrientLookup, vitClassifier)

	return foodDatabase
}
