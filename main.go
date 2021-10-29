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
func main(){
	args := os.Args


	user_options := make(map[string]string)
	user_options["help"] = "Get List Of Availiable commands and a brief description of how they work"
	user_options["search"] = "Search food names and recieve closest hits and their IDs"
	user_options["process"] = "Process food file and determine nutritional qualities with information on RDA"
	user_options["updateRda"] = " "

	if(args=="help"){
		fmt.Printf("")
	}





}
