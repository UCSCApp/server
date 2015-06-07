package dining

import (
	"github.com/ucscstudentapp/cloud/scraper"
)

var (
	MENU_TABLE_PATH = []scraper.Node{
		{"html", scraper.UNIQ},
		{"body", scraper.UNIQ},
		{"table", 0},
		{"tbody", scraper.UNIQ},
		{"tr", scraper.UNIQ},
		{"td", scraper.ALL},
	}
	MENU_ROWS_PATH = []scraper.Node{
		{"table", scraper.UNIQ},
		{"tbody", scraper.UNIQ},
		{"tr", 1},
		{"td", scraper.UNIQ},
		{"table", scraper.UNIQ},
		{"tbody", scraper.UNIQ},
		{"tr", scraper.UNIQ},
	}
	ROW_NAME_PATH = []scraper.Node{
		{"td", 0},
		{"table", scraper.UNIQ},
		{"tbody", scraper.UNIQ},
		{"tr", scraper.UNIQ},
		{"td", scraper.UNIQ},
		{"div.menusamprecipes", scraper.UNIQ},
		{"span", scraper.UNIQ},
	}
)

type menuTable struct {
	scraper.Selection
}

type menuDoc struct {
	scraper.Selection
}

type Menu struct {
	Breakfast MealMenu `json:"breakfast"`
	Lunch MealMenu `json:"lunch"`
	Dinner MealMenu `json:"dinner"`
}

type MealMenu []MenuItem

type MenuItem struct {
	Name string `json:"name"`
}

func (table menuTable) parseMenuItems(idx int) []MenuItem {
	rows := table.Index(idx).Path(MENU_ROWS_PATH)
	size := rows.Size()
	items := make([]MenuItem, size)
	for i := 0; i < size ; i++ {
		menuNameNode := rows.Index(i).Path(ROW_NAME_PATH)
		items[i] = MenuItem{menuNameNode.Inner(0).Data}
	}
	return items
}

func (table menuTable) parseBreakfastMenu() []MenuItem {
	if table.Size() != 3 {
		return nil
	} else {
		return table.parseMenuItems(0)
	}
}

func (table menuTable) parseLunchMenu() []MenuItem {
	if size := table.Size(); size == 1 {
		return nil
	} else if size == 2 {
		return table.parseMenuItems(0)

	} else {
		return table.parseMenuItems(1)
	}
}

func (table menuTable) parseDinnerMenu() []MenuItem {
	if size := table.Size(); size == 1 {
		return nil
	} else if size == 2 {
		return table.parseMenuItems(1)
	} else {
		return table.parseMenuItems(2)
	}
}


func (doc menuDoc) selectMenuTable() menuTable {
	sel := doc.Path(MENU_TABLE_PATH)
	return menuTable{sel}
}

func (doc menuDoc) Parse() Menu {
	menuTable := doc.selectMenuTable()
	return Menu{
		menuTable.parseBreakfastMenu(),
		menuTable.parseLunchMenu(),
		menuTable.parseDinnerMenu(),
	}
}
