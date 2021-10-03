package fpStructs

type Recipe struct{
	name string
	ingredients map[string]Food

}
func(r *Recipe) SetName(name string) {
	r.name=name
}
