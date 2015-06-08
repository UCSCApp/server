package dining

import (
	"encoding/json"
	"fmt"
	"github.com/ucscstudentapp/cloud/scraper"
	"os"
	"testing"
)

var (
	closedDiningHall = Menu{
		nil,
		nil,
		nil,
	}
	weekendDiningHall = Menu{
		nil,
		[]MenuItem{
			{"Mushroom & Barley Soup", nil},
			{"Stockpot Vegan Chili", nil},
			{"Apple Crepes Nancy", nil},
			{"Cage Free Scrambled Eggs", nil},
			{"Eggs Benedict", nil},
			{"Hard cooked Cage Free Eggs", nil},
			{"Natural BridgesTofu Scramble", nil},
			{"Oatmeal Gluten-Free", nil},
			{"Tator Tots", nil},
			{"3 Berry Muffin", nil},
			{"Blueberry Muffin", nil},
			{"Cowboy Cookies", nil},
			{"Donut Raised", nil},
			{"French Rolls", nil},
			{"Nutella Cheese Danish", nil},
			{"Orange Cream Cheese Spice Cake", nil},
			{"Paul's Vegan Cookies", nil},
			{"Bar Pasta", nil},
			{"Bread Sticks", nil},
			{"Cheese Manicotti with Marinara", nil},
			{"Condiments", nil},
			{"Marinara Sauce", nil},
			{"Meatballs", nil},
			{"Pasta Bar", nil},
			{"Penne", nil},
			{"Puttanesca Sauce", nil},
		},
		[]MenuItem{
			{"Korean BBQ Pork Spareribs", nil},
			{"Sizzling Thai Chicken Salad", nil},
			{"Sizzling Thai Seitan Salad", nil},
			{"5 Spice BBQ Beef Chow Mein", nil},
			{"5 Spice BBQ Tofu Chow Mein", nil},
			{"Veggie Fried Rice", nil},
			{"Chocolate Cream Pie", nil},
			{"French Rolls", nil},
			{"Orange Cream Cheese Spice Cake", nil},
			{"Bar Pasta", nil},
			{"Bread Sticks", nil},
			{"Cheese Manicotti with Marinara", nil},
			{"Condiments", nil},
			{"Marinara Sauce", nil},
			{"Meatballs", nil},
			{"Penne", nil},
			{"Puttanesca Sauce", nil},
		},
	}
)

func EqualMenuItems(this, that []MenuItem) bool {
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if this[i].Name != that[i].Name {
			return false
		}
	}
	return true
}

func EqualTestStruct(this, that Menu) bool {
	return EqualMenuItems(this.Breakfast, that.Breakfast) &&
		EqualMenuItems(this.Lunch, that.Lunch) &&
		EqualMenuItems(this.Dinner, that.Dinner)
}

func htmlFileTest(t *testing.T, path string, expected Menu) {
	read, err := os.Open(path)
	if err != nil {
		t.Skipf("Unable to open file: %s: %#v", path, err)
	}
	sel, err := scraper.NewFromReader(read)
	doc := menuDoc{sel}
	if err != nil {
		t.Errorf("Invalid HTML in file: %s: %#v", path, err)
	}
	actualResult := doc.Parse()
	if !EqualTestStruct(actualResult, expected) {
		t.Errorf("Values not equal: actual: %+v != expected: %+v", actualResult, closedDiningHall)
	}
}

func TestClosedDiningHall(t *testing.T) {
	htmlFileTest(t, "./generic_closed_menu.html", closedDiningHall)
}

func TestWeekendDiningHall(t *testing.T) {
	htmlFileTest(t, "./generic_weekend_open.html", weekendDiningHall)
}

func TestAll(t *testing.T) {
	byte, _ := json.Marshal(ParseAll())
	fmt.Println(string(byte))
}
