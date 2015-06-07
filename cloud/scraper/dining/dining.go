package dining

import (
	"github.com/ucscstudentapp/cloud/scraper"
	"fmt"
	"log"
)

const (
	website_format= "http://nutrition.sa.ucsc.edu/menuSamp.asp?locationNum=%d"
)

var (
	dhalls = []diningId{
		{"Cowell", 5},
		{"Crown & Merill", 20},
		{"Porter", 25},
		{"College Eight", 30},
		{"College Nine & Ten", 40},
	}
)

type diningId struct {
	name string
	locNum int
}

type DiningLocation struct {
	Name string `json:"name"`
	Menu Menu`json:"items"`
}

func handleUrlError(err error, url string) {
	if err != nil {
		log.Printf("Unable to create scraper for url: %s: %s", url, err.Error())
	}
}

func ParseAll() []DiningLocation {
	menus := make([]DiningLocation, len(dhalls))
	for i, v := range dhalls {
		url := fmt.Sprintf(website_format, v.locNum)
		menu, err := scraper.NewFromURL(url)
		handleUrlError(err, url)
		menus[i] = DiningLocation{v.name, menuDoc{menu}.Parse()}
	}
	return menus
}
