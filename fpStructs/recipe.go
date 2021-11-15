package fpStructs

type Recipe struct{
	name string
	ingredients []Food
	mass int
}
func MakeRecipe() *Recipe{
	r := Recipe{
		name:        "",
		ingredients: make([]Food,10),
	}
	return &r

}
func(r *Recipe) AddFood(food Food){
	r.ingredients=append(r.ingredients,food)
}
func(r *Recipe) SetName(name string) {
	r.name=name
}
func(r *Recipe) GetFood() []Food{
	return r.ingredients

}
