package main

import (
	"fmt"
	"os"
	//	"github.com/rivo/tview"
	//	"github.com/santigut123/food-parts/db_reader"
	//	"github.com/santigut123/food-parts/tui"
)

//"os"
//	"encoding/gob"

//		"github.com/gdamore/tcell/v2"
//	"github.com/rivo/tview"
type CommandInterface struct{
	userOptions map[string]string
}
func(ci *CommandInterface) ExecuteCommand(command []string){
	mainCommand := command[1];
	if mainCommand=="help" {
		ci.printHelp();
	} else if mainCommand=="search"{
		ci.search(command[2]);
	} else if mainCommand=="process"{

	} else if mainCommand=="makeTemplateRda"{

	} else if mainCommand=="makeTemplateFood"{

	}
}
func(ci *CommandInterface) search(searchString string){

}
func MakeCommandInterface() *CommandInterface{
	userOptions := make(map[string]string)
	userOptions["help"] = "Get List Of available commands and a brief description of how they work."
	userOptions["search"] = "Search food names and recieve closest hits and their IDs."
	userOptions["process"] = "Process food file and determine nutritional qualities with information on RDA."
	userOptions["makeTemplateRda"] = "Makes a template file for editing RDA."
	userOptions["makeTemplateFood"] = "Makes a template file for food information."

	newCommandInterface:=CommandInterface{
		userOptions: userOptions,
	}
	return &newCommandInterface
}
func(ci *CommandInterface) printHelp(){
	fmt.Println("Commands:");
	for key, element := range ci.userOptions{
		fmt.Print(key," \n    ",element+"\n\n");
	}
}
func main(){
	args := os.Args
	ui:= MakeCommandInterface();
	ui.ExecuteCommand(args);
}
