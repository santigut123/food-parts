package fpStructs
import (
	"github.com/derekparker/trie"
)
// This struct holds a database of food nutrition
type FoodDB struct{
	Name string
	Foods map[int]Food
	Search trie.Trie
}
func NewFoodDB(name string) *FoodDB{
	newDB:=FoodDB{
		Name:   name,
		Foods:  make(map[int]Food),
		Search: *trie.New(),
	}
	return &newDB
}
func(f *FoodDB) SetName(name string) {
	f.Name=name;
}
func(f *FoodDB) SetFoods(newFoods map[int]Food){
	f.Foods=newFoods
}
func(f *FoodDB) GetFoods() *map[int]Food{
	return &f.Foods
}
func(f *FoodDB) AddFood(newFood Food){
	f.Foods[newFood.GetID()]=newFood
	f.Search.Add(newFood.GetName(),newFood.GetID())
}
func(f *FoodDB) GetFoodSearch() *trie.Trie{
	return &f.Search
}
