package fullhous

import (
	"infohouse/infoAppliances"
	"infohouse/infoRooms"
	"infohouse/infofamily"
	"infohouse/infofurniture"
)

func House() {

	furniture := infofurniture.Furnitureitems()
	rooms := infoRooms.RoomsItems()
	family := infofamily.FamilyItims()
	appliances := infoAppliances.AppliancesItems()

	infofurniture.Showfurniture(furniture)
	infoRooms.ShowRoomDetails(rooms)
	infofamily.Showfamily(family)
	infoAppliances.ShowAppliances(appliances)
}
