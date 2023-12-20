package infoAppliances

import "fmt"

type Appliances struct {
	name string
}

func AppliancesItems() []Appliances {
	oven := Appliances{"духовка"}
	washingMachine := Appliances{"стиральная машина"}
	fridge := Appliances{"холодильник"}
	hairdryer := Appliances{"фен"}
	blender := Appliances{"блендер"}
	iron := Appliances{"утюг"}
	return []Appliances{oven, washingMachine, fridge, hairdryer, blender, iron}
}

func ShowAppliances(appliances []Appliances) {
	fmt.Println("Бытовая техника:")
	for _, everyItem := range appliances {
		fmt.Print(everyItem.name, "\n")
	}

}
