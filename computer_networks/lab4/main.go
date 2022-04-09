package main

import (
	"fmt"
	"html/template"
	"net/http"

	"golang.org/x/net/html"
)

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

func readItem(item *html.Node, variant string) *One_news {
	if a := item.FirstChild; isElem(a, "a") {
		if cs := getChildren(a); len(cs) == 2 && isElem(cs[0], "time") && isText(cs[1]) && variant == "span4" {
			return &One_news{
				Link:      "https://lenta.ru" + getAttr(a, "href"),
				Published: getAttr(cs[0], "title"),
				Title:     cs[1].Data,
			}
		} else {
			return &One_news{
				Link:      "https://lenta.ru" + getAttr(a, "href"),
				Published: "В желтом блоке нет времени публикации",
				Title:     getChildren(a)[0].Data,
			}
		}
	}
	return nil
}

type One_news struct {
	Link, Published, Title string
}

type ViewData struct {
	Lenta []*One_news
}

func downloadNews() []*One_news {
	fmt.Println("sending request to lenta.ru")
	if response, err := http.Get("http://lenta.ru"); err != nil {
		fmt.Println("request to lenta.ru failed", "error", err)
	} else {
		defer response.Body.Close()
		status := response.StatusCode
		fmt.Println("got response from lenta.ru", "status", status)
		if status == http.StatusOK {
			if doc, err := html.Parse(response.Body); err != nil {
				fmt.Println("invalid HTML from lenta.ru", "error", err)
			} else {
				fmt.Println("HTML from lenta.ru parsed successfully")
				return search(doc)
			}
		}
	}
	return nil
}

var items []*One_news

func search(node *html.Node) []*One_news {
	var res []*One_news
	if isDiv(node, "span4") {
		var items []*One_news
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			if isDiv(c, "item") {
				if item := readItem(c, "span4"); item != nil {
					items = append(items, item)
				}
			}
		}
		res = append(res, items...)
	}

	if isDiv(node, "b-yellow-box__wrap") {
		var items []*One_news
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			if isDiv(c, "item") {
				if item := readItem(c, "yellow"); item != nil {
					items = append(items, item)
				}
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

func mainpage(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Downloader started")
	news := downloadNews()

	data := ViewData{Lenta: news}
	tmpl, _ := template.ParseFiles("index.html")
	tmpl.Execute(w, data)
}

func main() {

	http.HandleFunc("/", mainpage)
	err := http.ListenAndServe(":2000", nil) // задаем слушать порт
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}

}
