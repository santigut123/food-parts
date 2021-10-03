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
func(f *FoodDB) SetName(name string) {
	f.name=name;
}

func(f *FoodDB) SetFoods(newFoods map[int]Food){
	f.foods=newFoods
}
func(f *FoodDB) GetFoods() *map[int]Food{
	return &f.foods
}
func(f *FoodDB) addFood(id int ,newFood Food){
	f.foods[id]=newFood
	f.search.Add(newFood.name,id)
}
func(f *FoodDB) GetFoodSearch() *trie.Trie{
	return &f.search
}
