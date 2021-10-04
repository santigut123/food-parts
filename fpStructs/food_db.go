package fpStructs
import (
	"github.com/derekparker/trie"
)
// This struct holds a database of food nutrition
type FoodDB struct{
	name string
	foods map[int]Food
	search trie.Trie
}
func NewFoodDB(name string) *FoodDB{
	newDB:=FoodDB{
		name:   name,
		foods:  make(map[int]Food),
		search: *trie.New(),
	}
	return &newDB
}
func(f *FoodDB) SetName(name string) {
	f.name=name;
}
func(f *FoodDB) SetFoods(newFoods map[int]Food){
	f.foods=newFoods
}
func(f *FoodDB) GetFoods() *map[int]Food{
	return &f.foods
}
func(f *FoodDB) AddFood(newFood Food){
	f.foods[newFood.GetID()]=newFood
	f.search.Add(newFood.GetName(),newFood.GetID())
}
func(f *FoodDB) GetFoodSearch() *trie.Trie{
	return &f.search
}
