package dining

import (
	"github.com/ucscstudentapp/cloud/scraper"
	"strings"
)

var (
	menu_table_path = []scraper.Node{
		{"html", scraper.UNIQ},
		{"body", scraper.UNIQ},
		{"table", 0},
		{"tbody", scraper.UNIQ},
		{"tr", scraper.UNIQ},
		{"td", scraper.ALL},
	}
	menu_rows_path = []scraper.Node{
		{"table", scraper.UNIQ},
		{"tbody", scraper.UNIQ},
		{"tr", 1},
		{"td", scraper.UNIQ},
		{"table", scraper.UNIQ},
		{"tbody", scraper.UNIQ},
		{"tr", scraper.UNIQ},
	}
	row_item_path = []scraper.Node{
		{"td", 0},
		{"table", scraper.UNIQ},
		{"tbody", scraper.UNIQ},
		{"tr", scraper.UNIQ},
	}
	item_name_path = []scraper.Node{
		{"td", 0},
		{"div.menusamprecipes", scraper.UNIQ},
		{"span", scraper.UNIQ},
	}
	item_attrib_path = []scraper.Node{
		{"td", scraper.ALL},
		{"img", scraper.ALL},
	}
)

type menuTable struct {
	scraper.Selection
}

type menuDoc struct {
	scraper.Selection
}

type menuItemNode struct {
	scraper.Selection
}

type Menu struct {
	Breakfast MealMenu `json:"breakfast"`
	Lunch     MealMenu `json:"lunch"`
	Dinner    MealMenu `json:"dinner"`
}

type MealMenu []MenuItem

type MenuItem struct {
	Name    string   `json:"name"`
	Attribs []string `json:"attribs"`
}

func (node menuItemNode) parse() MenuItem {
	nameNode := node.Path(item_name_path)
	attribNodes := node.Path(item_attrib_path).Nodes()
	name := nameNode.Inner(0).Data
	attribs := make([]string, len(attribNodes))
	for i, v := range attribNodes {
		var srcVal string
		for _, vi := range v.Attr {
			if vi.Key == "src" {
				srcVal = vi.Val
			}
		}
		srcVal = strings.Replace(srcVal, "LegendImages/", "", -1)
		srcVal = strings.Replace(srcVal, ".gif", "", -1)
		attribs[i] = srcVal
	}
	return MenuItem{name, attribs}
}

func (table menuTable) parseMenuItems(idx int) []MenuItem {
	rows := table.Index(idx).Path(menu_rows_path)
	size := rows.Size()
	items := make([]MenuItem, size)
	for i := 0; i < size; i++ {
		node := menuItemNode{rows.Index(i).Path(row_item_path)}
		items[i] = node.parse()
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
	sel := doc.Path(menu_table_path)
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
