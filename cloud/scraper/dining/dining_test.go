package dining

import (
	"testing"
	"os"
	"github.com/ucscstudentapp/cloud/scraper"
)

type testStruct struct {
	Breakfast []MenuItem
	Lunch []MenuItem
	Dinner []MenuItem
}

var (
	closedDiningHall = testStruct{
		nil,
		nil,
		nil,
	}
	weekendDiningHall = testStruct{
		nil,
		[]MenuItem{{"Mushroom & Barley Soup"},
			{"Stockpot Vegan Chili"},
			{"Apple Crepes Nancy"},
			{"Cage Free Scrambled Eggs"},
			{"Eggs Benedict"},
			{"Hard cooked Cage Free Eggs"},
			{"Natural BridgesTofu Scramble"},
			{"Oatmeal Gluten-Free"},
			{"Tator Tots"},
			{"3 Berry Muffin"},
			{"Blueberry Muffin"},
			{"Cowboy Cookies"},
			{"Donut Raised"},
			{"French Rolls"},
			{"Nutella Cheese Danish"},
			{"Orange Cream Cheese Spice Cake"},
			{"Paul's Vegan Cookies"},
			{"Bar Pasta"},
			{"Bread Sticks"},
			{"Cheese Manicotti with Marinara"},
			{"Condiments"},
			{"Marinara Sauce"},
			{"Meatballs"},
			{"Pasta Bar"},
			{"Penne"},
			{"Puttanesca Sauce"},
		},
		[]MenuItem{
			{"Korean BBQ Pork Spareribs"},
			{"Sizzling Thai Chicken Salad"}, 
			{"Sizzling Thai Seitan Salad"}, 
			{"5 Spice BBQ Beef Chow Mein"}, 
			{"5 Spice BBQ Tofu Chow Mein"}, 
			{"Veggie Fried Rice"}, 
			{"Chocolate Cream Pie"}, 
			{"French Rolls"}, 
			{"Orange Cream Cheese Spice Cake"}, 
			{"Bar Pasta"}, 
			{"Bread Sticks"}, 
			{"Cheese Manicotti with Marinara"}, 
			{"Condiments"}, 
			{"Marinara Sauce"}, 
			{"Meatballs"}, 
			{"Penne"}, 
			{"Puttanesca Sauce"},
		},
	}
)

func EqualMenuItems(this, that []MenuItem) bool {
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if this[i] != that[i] {
			return false
		}
	}
	return true
}

func EqualTestStruct(this, that testStruct) bool {
	return EqualMenuItems(this.Breakfast, that.Breakfast) &&
	       EqualMenuItems(this.Lunch, that.Lunch) &&
	       EqualMenuItems(this.Dinner, that.Dinner)
}

func htmlFileTest(t *testing.T, path string, expected testStruct) {
	read, err := os.Open(path)
	if err != nil {
		t.Skipf("Unable to open file: %s: %#v", path, err)
	}
	sel, err := scraper.NewFromReader(read)
	doc := menuDoc{sel}
	if err != nil {
		t.Errorf("Invalid HTML in file: %s: %#v", path, err)
	}
	menuTable := doc.selectMenuTable()
	actualResult := testStruct{
		menuTable.parseBreakfastMenu(),
		menuTable.parseLunchMenu(),
		menuTable.parseDinnerMenu(),
	}
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
