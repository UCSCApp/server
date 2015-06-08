package scraper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"io"
)

type Node struct {
	Name string
	Idx  int
}

const (
	UNIQ = -1
	ALL  = -2
)

type Selection struct {
	impl *goquery.Selection
}

func (sel Selection) Print() {
	for i, v := range sel.impl.Nodes {
		fmt.Printf("%d: %s\n", i, v.Data)
	}
}

func (sel Selection) PrintChildren() {
	children := Selection{sel.impl.Children()}
	children.Print()
}

func NewFromURL(url string) (Selection, error) {
	doc, err := goquery.NewDocument(url)
	return Selection{(*doc).Selection}, err
}

func NewFromReader(reader io.Reader) (Selection, error) {
	doc, err := goquery.NewDocumentFromReader(reader)
	return Selection{(*doc).Selection}, err
}

func (sel Selection) ChildrenFilter(selector string) Selection {
	return Selection{sel.impl.Children().Filter(selector)}
}

func (sel Selection) Index(idx int) Selection {
	return Selection{sel.impl.Eq(idx)}
}

func (sel Selection) Children() Selection {
	return Selection{sel.impl.Children()}
}

func (sel Selection) Size() int {
	return sel.impl.Size()
}

func (sel Selection) Nodes() []*html.Node {
	return sel.impl.Nodes
}

func (sel Selection) Inner(idx int) *html.Node {
	return sel.impl.Contents().Get(idx)
}

func (sel Selection) Path(nodes []Node) Selection {
	final := sel
	for _, v := range nodes {
		final = final.ChildrenFilter(v.Name)
		if idx := v.Idx; idx != UNIQ && idx != ALL {
			final = final.Index(idx)
		}
	}
	return final
}
