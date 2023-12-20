package infofamily

import "fmt"

type personStruct struct {
	name    string
	surname string
	age     int
}

func FamilyItims() []personStruct {
	man1 := personStruct{"олимжон", "акбаралиев", 23}
	man2 := personStruct{"ахрор", "акбаралиев", 19}
	man3 := personStruct{"баха", "хусанов", 23}

	return []personStruct{man1, man2, man3}
}

func Showfamily(family []personStruct) {
	fmt.Println("\n Состав семьи:")
	fmt.Println("-------------")
	for _, person := range family {
		fmt.Printf("Имя: %s \nФамилия: %s\nВозраст: %d\n", person.name, person.surname, person.age)
	}
}
