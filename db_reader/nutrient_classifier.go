package db_reader

type DBVitClassifier struct{
	vitaminNumbers []int
	mineralNumbers []int
	aminoNumbers []int
	macroNumbers []int
	vitSet map[int]bool
	mineralSet map[int]bool
	aminoSet map[int]bool
	macroSet map[int]bool

}
func fillSet(array []int,set map[int]bool){
	for _,v :=range array{
		set[v]=true
	}
}
func(dbvc *DBVitClassifier) fillSets(){
	fillSet(dbvc.vitaminNumbers, dbvc.vitSet)
	fillSet(dbvc.mineralNumbers, dbvc.mineralSet)
	fillSet(dbvc.aminoNumbers, dbvc.aminoSet)
	fillSet(dbvc.macroNumbers, dbvc.macroSet)
}
func newDBVitClassifier() *DBVitClassifier{
	dbVitClass:=DBVitClassifier{
		vitaminNumbers: []int{1106, 1109, 1114, 1124, 1162, 1165, 1166, 1167, 1174, 1175, 1176, 1177, 1178, 1180, 1183, 1184, 1185, 1187},
		mineralNumbers: []int{1087, 1088, 1089, 1090, 1091, 1092, 1093, 1094, 1095, 1096, 1097, 1098, 1099, 1100, 1101, 1102, 1103, 1141, 1142, 1149},
		aminoNumbers:   []int{1213, 1214, 1215, 1216, 1217, 1218, 1219, 1220, 1221, 1222, 1223, 1224, 1225, 1226, 1227, 1232, 1233, 1234},
		macroNumbers:   []int{1003, 1004, 1005, 1009, 1010, 1011, 1012, 1013, 1050, 1063, 1257, 1258, 1291, 1292, 1293},
		vitSet:         map[int]bool{},
		mineralSet:     map[int]bool{},
		aminoSet:       map[int]bool{},
		macroSet:       map[int]bool{},
	}
	dbVitClass.fillSets()
	return &dbVitClass
}
func(dbvc *DBVitClassifier) GetNType(id int) rune{
	_,vitCheck := dbvc.vitSet[id]
	_,mineralCheck := dbvc.mineralSet[id]
	_,aminoCheck := dbvc.aminoSet[id]
	_,macroCheck := dbvc.macroSet[id]
	cat:='x'
	if vitCheck{
		cat='v'
	}else if mineralCheck{
		cat= 'm'
	}else if aminoCheck{
		cat= 'a'
	}else if macroCheck{
		cat= 'M'
	}
	return cat

}
func(dbvc *DBVitClassifier) IsProtein(id int ) bool{
	if(id==1003){
		return true
	}
	return false;
	}
func(dbvc *DBVitClassifier) IsFat(id int ) bool{
	if(id==1004){
		return true
	}
	return false;
	}
func(dbvc *DBVitClassifier) IsCarbs(id int ) bool{
	if(id==1005||id==1050){
		return true
	}
	return false;
	}
