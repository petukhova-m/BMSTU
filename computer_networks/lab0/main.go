// package main

// import (
// 	"html/template"
// 	"log"
// 	"net/http"
// 	"sort"

// 	"github.com/mmcdole/gofeed"
// )

// type ViewData struct {
// 	Lenta []News
// }

// type News struct {
// 	Link        string
// 	Description string
// 	Published   string
// 	Title       string
// }

// func Rss_1(w http.ResponseWriter, r *http.Request) {
// 	Rss1 := gofeed.NewParser()
// 	feed1, _ := Rss1.ParseURL("https://lenta.ru/rss")

// 	Rss2 := gofeed.NewParser()
// 	feed2, _ := Rss2.ParseURL("https://news.mail.ru/rss/90/")

// 	var lenta []gofeed.Item

// 	for i := range feed1.Items {
// 		lenta = append(lenta, *feed1.Items[i])
// 	}

// 	for i := range feed2.Items {
// 		lenta = append(lenta, *feed2.Items[i])
// 	}

// 	sort.Slice(lenta, func(i, j int) bool {
// 		return lenta[i].Published < lenta[j].Published
// 	})

// 	data := ViewData{}

// 	for i := range lenta {
// 		var New = lenta[i]

// 		var oneNews = News{
// 			Link:        New.Link,
// 			Description: New.Description,
// 			Published:   New.Published,
// 			Title:       New.Title,
// 		}

// 		data.Lenta = append(data.Lenta, oneNews)

// 		tmpl, _ := template.ParseFiles("index.html")
// 		tmpl.Execute(w, data)
// 	}
// }

// func main() {
// 	http.HandleFunc("/", Rss_1)
// 	err := http.ListenAndServe(":9015", nil) // задаем слушать порт
// 	if err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// }
