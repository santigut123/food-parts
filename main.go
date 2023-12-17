package main

import (
	"fmt"
	"log"
	"os"
	"github.com/santigut123/food-parts/cliInterface"
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
	if len(args)==1{
		log.Fatal("No specified command, type foodparts help for list of commands.")
	}
	for index,x :=range args{
		fmt.Println("Arg ",index)
		fmt.Println("Content ",x)
	}
	clInteface:=cliInterface.MakeCommandInterface(args)
	clInteface.ExecuteCommand()

}
