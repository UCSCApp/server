package dining

import (
	"fmt"
	"github.com/ucscstudentapp/cloud/scraper"
	"log"
)

const (
	website_format = "http://nutrition.sa.ucsc.edu/menuSamp.asp?locationNum=%s"
)

var (
	dhalls = []diningId{
		{"Cowell", "05"},
		{"Crown & Merill", "20"},
		{"Porter", "25"},
		{"College Eight", "30"},
		{"College Nine & Ten", "40"},
	}
)

type diningId struct {
	name  string
	locId string
}

type DiningLocation struct {
	Name string `json:"name"`
	Menu Menu   `json:"items"`
}

func handleUrlError(err error, url string) {
	if err != nil {
		log.Printf("Unable to create scraper for url: %s: %s", url, err.Error())
	}
}

func ParseAll() []DiningLocation {
	menus := make([]DiningLocation, len(dhalls))
	for i, v := range dhalls {
		url := fmt.Sprintf(website_format, v.locId)
		menu, err := scraper.NewFromURL(url)
		handleUrlError(err, url)
		menus[i] = DiningLocation{v.name, menuDoc{menu}.Parse()}
	}
	return menus
}
