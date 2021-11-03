package cliInterface

import (
	"fmt"
	"strconv"

	"github.com/santigut123/food-parts/db_reader"
	"github.com/santigut123/food-parts/fpStructs"
)
type CommandInterface struct{
	userOptions map[string]string
	args []string
}
func(ci *CommandInterface) PrintFoodSearch (foodDB *fpStructs.FoodDB){
	foodName:= ci.args[2]
	foodResults:= foodDB.Search.SearchFood(foodName)
	for _,x := range foodResults {
		fmt.Println("--------------------------------")
		fmt.Println("ID: ",foodDB.Search.GetFoodID(x)," \nName: ",x)
	}
}
func (ci *CommandInterface) PrintNutrients(f *fpStructs.Food){
	fmt.Println("Nutrients:")
	fmt.Println("  Macros")
	for k,v := range f.Nutrients.Macros{
		fmt.Println("    Name: ",k," Amount: ", v.Mass,v.Units)
	}
	fmt.Println("  Vitamins")
	for k,v := range f.Nutrients.Vitamins{
		fmt.Println("    Name: ",k," Amount: ", v.Mass,v.Units)
	}
	fmt.Println("  Minerals")
	for k,v := range f.Nutrients.Minerals{
		fmt.Println("    Name: ",k," Amount: ", v.Mass,v.Units)
	}
	fmt.Println("  Amino Acids")
	for k,v := range f.Nutrients.AminoAcids{
		fmt.Println("    Name: ",k," Amount: ", v.Mass,v.Units)
	}
}
func(ci *CommandInterface) PrintFoodInfo(db *fpStructs.FoodDB){
	id,_ := strconv.Atoi(ci.args[2])
	food := db.GetFood(id)
	fmt.Println("--------------------------------")
	fmt.Println("Name:",food.GetName())
	ci.PrintNutrients(food)
}
func(ci* CommandInterface) ProcessFile (filename string){

}
func(ci *CommandInterface) ExecuteCommand(){
	mainCommand := ci.args[1];
	if mainCommand=="help" {
		ci.printHelp();
	} else if mainCommand=="makeTemplateRda"{

	} else if mainCommand=="makeTemplateFood"{

	} else {
		db := db_reader.ReadDatabase()
		if mainCommand=="search"{
			ci.PrintFoodSearch(db)
		}else if mainCommand=="getInfo"{
			ci.PrintFoodInfo(db)
		}else if mainCommand=="process"{
		}
	}
}
func(ci *CommandInterface) makeDatabase() *fpStructs.FoodDB{
	db:=db_reader.ReadDatabase();
	return db
}
func(ci *CommandInterface) search(searchString string){

}
func MakeCommandInterface(args []string) *CommandInterface{
	userOptions := make(map[string]string)
	userOptions["help"] = "Get List Of available commands and a brief description of how they work."
	userOptions["search"] = "Search food names and recieve closest hits and their IDs."
	userOptions["process"] = "Process food file and determine nutritional qualities with information on RDA."
	userOptions["makeTemplateRda"] = "Makes a template file for editing RDA."
	userOptions["makeTemplateFood"] = "Makes a template file for food information."
	userOptions["getInfo"] ="Gets all info about a food using an ID"

	newCommandInterface:=CommandInterface{
		userOptions: userOptions,
		args:        args,
	}
	return &newCommandInterface
}
func(ci *CommandInterface) printHelp(){
	fmt.Println("Commands:");
	for key, element := range ci.userOptions{
		fmt.Print(key," \n    ",element+"\n\n");
	}
}
