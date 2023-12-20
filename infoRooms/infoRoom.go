package infoRooms

import "fmt"

type Rooms struct {
	Name string
	h    int
	w    int
	l    int
}

func RoomsItems() []Rooms {

	livingRoom := Rooms{"Гостинная", 3, 4, 5}
	kitchen := Rooms{"Кухня", 3, 3, 4}
	bathroom := Rooms{"Ванная", 3, 2, 3}

	return []Rooms{livingRoom, kitchen, bathroom}
}

func ShowRoomDetails(rooms []Rooms) {
	fmt.Print("плошадь дома:\n")
	fmt.Println("-----------------")
	for _, room := range rooms {
		fmt.Printf("Комната: %s \nПлощадь: %d кв м\n", room.Name, room.w*room.l)
	}

}
