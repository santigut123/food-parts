package cliInterface

import (
	"fmt"

	"github.com/santigut123/food-parts/db_reader"
	"github.com/santigut123/food-parts/fpStructs"
)
type CommandInterface struct{
	userOptions map[string]string
	args []string
}
func(ci *CommandInterface) PrintFoodSearch (foodDB *fpStructs.FoodDB,foodName string){
	foodResults:= foodDB.Search.SearchFood(foodName)
	for _,x := range foodResults {
		fmt.Println("Name: ",x,"ID: ",foodDB.Search.GetFoodID(x))
	}
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
			ci.PrintFoodSearch(db, ci.args[2])
		} else if mainCommand=="process"{

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
