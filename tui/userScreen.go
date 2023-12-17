package tui

import (
	"fmt"

	"github.com/derekparker/trie"
	//	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)
type SearchWidget struct{
	search trie.Trie
}
func printMenu(){
	menuoptions := {"test"}
	fmt.Println("--------Food Parts--------")
	fmt("")
}
func InitializeTUI(){

}
func NewSearch(SearchTrie *trie.Trie) *SearchWidget{
	newSearchWidget:=SearchWidget{
		search: *SearchTrie,
	}
	return &newSearchWidget
}
