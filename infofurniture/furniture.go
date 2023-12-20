package infofurniture

import "fmt"

type furnitureStruct struct {
	name string
}

func Furnitureitems() []furnitureStruct {

	computer := furnitureStruct{"компьютер"}
	closet := furnitureStruct{"шкаф"}
	couch := furnitureStruct{"диван"}
	table := furnitureStruct{"стол"}
	armchair := furnitureStruct{"кресло"}
	return []furnitureStruct{computer, closet, couch, table, armchair}
}

func Showfurniture(furniture []furnitureStruct) {

	fmt.Println("\nМебель в комнате:")
	fmt.Println("-----------------")
	for _, item := range furniture {
		fmt.Println(item.name, " ")
	}
	fmt.Print("\n")
}
