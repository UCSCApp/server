package dining

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
)

type PathNode struct {
	Name string
	Idx  int
}

const (
	UNIQ    = -1
	ALL     = -2
	WEBSITE = "http://nutrition.sa.ucsc.edu/menuSamp.asp?locationNum=30&locationName=&sName=&naFlag="
)

var (
	MENU_TABLE_PATH = []PathNode{
		{"html", UNIQ},
		{"body", UNIQ},
		{"table", 0},
		{"tbody", UNIQ},
		{"tr", UNIQ},
	}
	MENU_ROWS_PATH = []PathNode{
		{"table", UNIQ},
		{"tbody", UNIQ},
		{"tr", 1},
		{"td", UNIQ},
		{"table", UNIQ},
		{"tbody", UNIQ},
	}
	ROW_NAME_PATH = []PathNode{
		{"td", 0},
		{"table", UNIQ},
		{"tbody", UNIQ},
		{"tr", UNIQ},
		{"td", UNIQ},
		{"div.menusamprecipes", UNIQ},
		{"span", UNIQ},
	}
)

type selection struct {
	*goquery.Selection
}

type menuTable selection
type menuDoc selection

type MenuItem struct {
	name string
}

func (sel selection) Print() {
	for i, v := range sel.Nodes {
		fmt.Printf("%d: %s\n", i, v.Data)
	}
}

func (sel selection) PrintChildren() {
	selection{sel.Children()}.Print()
}

func (sel selection) PrintNodes() {
	for i, v := range sel.Contents().Nodes {
		fmt.Printf("%d: %s\n", i, v.Data)
	}
}

func failOnError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open C8 Dining Hall Website: %v", err)
		os.Exit(1)
	}
}

func (table menuTable) parseMenuItems(idx int) []MenuItem {
	selection{table.Children().Filter("td").Eq(idx).Children().Filter("table").Children().Filter("tbody").Children().Filter("tr").Eq(1).Children().Filter("td").Children().Filter("table").Children().Filter("tbody")}.PrintChildren()
	rows := selection{selection{table.Children().Filter("td").Eq(idx)}.transPath(MENU_ROWS_PATH).Children().Filter("tr")}
	size := rows.Size()
	items := make([]MenuItem, size)
	fmt.Println(size)
	for i := 1; i < size; i++ {
		selection{rows.Eq(i)}.transPath(ROW_NAME_PATH).PrintNodes()
		items[i-1] = MenuItem{selection{rows.Eq(i)}.transPath(ROW_NAME_PATH).Contents().Get(0).Data}
	}
	return items
}

func (table menuTable) parseBreakfastMenu() []MenuItem {
	if table.Children().Size() != 3 {
		return nil
	} else {
		return table.parseMenuItems(0)
	}
}

func (table menuTable) parseLunchMenu() []MenuItem {
	if table.Children().Size() != 3 {
		return table.parseMenuItems(0)
	} else {
		return table.parseMenuItems(1)
	}
}

func (table menuTable) parseDinnerMenu() []MenuItem {
	if table.Children().Size() != 3 {
		return table.parseMenuItems(1)
	} else {
		return table.parseMenuItems(2)
	}
}

func (sel selection) transPath(nodes []PathNode) (final selection) {
	final = sel
	for _, v := range nodes {
		children := selection{final.Children()}
		final = selection{children.Filter(v.Name)}
		if idx := v.Idx; idx != UNIQ {
			final = selection{final.Eq(idx)}
		}
	}
	return
}

func (doc menuDoc) selectMenuTable() menuTable {
	selection := selection(doc).transPath(MENU_TABLE_PATH)
	return menuTable(selection)
}

func c8Doc() menuDoc {
	doc, err := goquery.NewDocument(WEBSITE)
	failOnError(err)
	return menuDoc{(*doc).Selection}
}
