package main

import (
	"infohouse/infoAppliances"
	"infohouse/infoRooms"
	"infohouse/infofamily"
	"infohouse/infofurniture"
)

func main() {

	furniture := infofurniture.Furnitureitems()
	rooms := infoRooms.RoomsItems()
	family := infofamily.FamilyItims()
	appliances := infoAppliances.AppliancesItems()

	infofurniture.Showfurniture(furniture)
	infoRooms.ShowRoomDetails(rooms)
	infofamily.Showfamily(family)
	infoAppliances.ShowAppliances(appliances)
}
