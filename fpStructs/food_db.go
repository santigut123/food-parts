package fpStructs

import (
	"strings"

	"github.com/derekparker/trie"
)
type FoodSearch struct{
	searchTrie trie.Trie
}
func makeFoodSearch() *FoodSearch{
	fs := FoodSearch{
		searchTrie: *trie.New(),
	}
	return &fs

}
func (fs *FoodSearch) addFood(name string, id int){
	fs.searchTrie.Add(strings.ToUpper(name),id)
}
func(fs *FoodSearch) SearchFood(foodName string) []string{
	return fs.searchTrie.PrefixSearch(foodName);
}
func(fs *FoodSearch) GetFoodID(foodName string) int{
	food,_:=fs.searchTrie.Find(foodName)
	id := food.Meta()
	return id.(int)
}
// This struct holds a database of food nutrition
type FoodDB struct{
	Name string
	Foods map[int]*Food
	Search FoodSearch
	Recipes map[string]Recipe
}
func NewFoodDB(name string) *FoodDB{
	newDB:=FoodDB{
		Name:    name,
		Foods:   make(map[int]*Food),
		Search:  *makeFoodSearch(),
		Recipes: map[string]Recipe{},
	}
	return &newDB
}
func(f *FoodDB) SetName(name string) {
	f.Name=name;
}
func(f *FoodDB) SetFoods(newFoods map[int]*Food){
	f.Foods=newFoods
}
func(f *FoodDB) GetFood(id int) *Food{
	food := f.Foods[id]
	return food
}
func(f *FoodDB) GetFoods() *map[int]*Food{
	return &f.Foods
}
func(f *FoodDB) AddFood(newFood *Food){
	f.Foods[newFood.GetID()]=newFood
	f.Search.addFood(newFood.GetName(),newFood.GetID())
}
func(f *FoodDB) AddRecipe (r Recipe){
	f.Recipes[r.name] = r
}
func(f *FoodDB) GetRecipeFoods(recipeName string) []Food{
	return f.Recipes[recipeName].ingredients

}
