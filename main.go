package main

import (
	//	"fmt"
	//"os"
	//"./food_components"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"encoding/gob"
)
func main(){
	//
	fp:=tview.NewApplication()
	modal := tview.NewModal().
		SetText("Do you want to quit the application?").
		AddButtons([]string{"Quit", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Quit" {
				fp.Stop()
			}
		})
	fpView := tview.NewFlex().
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Left (1/2 x width of Top)"), 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Food-Parts"), 0, 1, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Middle (3 x height of Top)"), 0, 3, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Bottom (5 rows)"), 5, 1, false), 0, 2, false)
		fp.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Rune() == 'q' {
				 fp.SetRoot(modal,false).SetFocus(modal)
			}
	return event
})	/*exitModal:=tview.NewModal().
		SetText("Do you want to quit the application").
		SetDoneFunc(func(buttonIndex int,buttonLabel string){
			if buttonLabel=="Quit"{
				fp.Stop()
			}
		})
	*/
	if err := fp.SetRoot(fpView, true).SetFocus(fpView).Run(); err != nil {
		panic(err)
	}




}
