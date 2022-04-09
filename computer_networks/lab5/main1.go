package main

import (
	"context"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"golang.org/x/net/html"

	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
)

type One_news struct {
	Link, Title string
}

type ViewData struct {
	Lenta []*One_news
}

func getChildren(node *html.Node) []*html.Node {
	var children []*html.Node
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		children = append(children, c)
	}
	return children
}

func getAttr(node *html.Node, key string) string {
	for _, attr := range node.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

func isText(node *html.Node) bool {
	return node != nil && node.Type == html.TextNode
}

func isElem(node *html.Node, tag string) bool {
	return node != nil && node.Type == html.ElementNode && node.Data == tag
}

func isDiv(node *html.Node, class string) bool {
	return isElem(node, "div") && getAttr(node, "class") == class
}

func isSpan(node *html.Node) bool {
	return isElem(node, "span") && getAttr(node, "class") == "main__feed__title"
}

var items []*One_news

func search1(node *html.Node) string {

	if isElem(node, "span") && getAttr(node, "class") == "main__feed__title" {
		return node.FirstChild.Data
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if items := search1(c); items != "" {
			return items
		}
	}
	return ""
}

func search(node *html.Node) []*One_news {
	var res []*One_news

	if isDiv(node, "main__feed js-main-reload-item") {
		fmt.Println("Нашли Новость!!!")
		var items []*One_news
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			if isElem(c, "a") {
				item := &One_news{
					Link:  getAttr(c, "href"),
					Title: search1(c),
				}
				items = append(res, item)
			}
		}
		res = append(res, items...)
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if items := search(c); items != nil {
			res = append(res, items...)
		}
	}
	return res
}

func getCompleteHTML(link string) string {
	var res string
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	err := chromedp.Run(ctx,
		chromedp.Navigate("http://"+link),
		chromedp.ActionFunc(func(ctx context.Context) error {
			node, err := dom.GetDocument().Do(ctx)
			res, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			return err
		}),
	)
	if err != nil {
		panic(err)
	}
	return res
}

func downloadNews() []*One_news {
	fmt.Println("sending request to rbc.ru")
	getIn := getCompleteHTML("www.rbc.ru/")
	re := regexp.MustCompile(`(<!--)|(-->)`)
	p := re.ReplaceAllString(getIn, ``)
	ioutil.WriteFile("123.html", []byte(p), 0644)
	if doc, err := html.Parse(strings.NewReader(string(p))); err != nil {
		fmt.Println("invalid HTML from rbc.ru", "error", err)
	} else {
		fmt.Println("HTML from rbc.ru parsed successfully")

		return search(doc)
	}

	return nil
}

func mainpage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Downloader started")
	news := downloadNews()
	data := ViewData{Lenta: news}
	fmt.Println(data.Lenta)
	tmpl, _ := template.ParseFiles("index.html")
	tmpl.Execute(w, data)
}

func main() {

	http.HandleFunc("/", mainpage)
	err := http.ListenAndServe(":9785", nil) // задаем слушать порт
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
