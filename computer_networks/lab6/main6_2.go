package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

const (
	PORT = "8080"
)

func main() {
	http.HandleFunc("/", handleRequestAndRedirect)

	http.HandleFunc("/link", google)

	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}

func google(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("google.html")
	if err != nil {
		log.Fatal("Unable to open html: ", err.Error())
	}

	if err := tmpl.Execute(w, nil); err != nil {
		log.Fatal("Unable to execute: ", err.Error())
	}

}

func handleRequestAndRedirect(w http.ResponseWriter, r *http.Request) {
	link := r.URL.Query().Get("link")

	req, err := http.NewRequest(r.Method, link, nil)
	if err != nil {
		http.Error(w, fmt.Errorf("new request %w", err).Error(), http.StatusServiceUnavailable)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, fmt.Errorf("do %w", err).Error(), http.StatusServiceUnavailable)
		return
	}

	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)

	contentType := resp.Header.Get("Content-type")

	if strings.Contains(contentType, "text/html") {
		node, err := html.Parse(resp.Body)
		if err != nil {
			http.Error(w, fmt.Errorf("html parse %w", err).Error(), http.StatusServiceUnavailable)
			return
		}
		defer resp.Body.Close()

		changeURL(node, req.URL)

		var buffer bytes.Buffer
		html.Render(&buffer, node)
		ans := buffer.String()

		w.Write([]byte(ans))
		return
	}

	io.Copy(w, resp.Body)
	return
}

func copyHeader(dst, src http.Header) {
	for key, values := range src {
		if key != "Content-Security-Policy" {
			for _, val := range values {
				dst.Add(key, val)
			}
		}
	}
}

func changeURL(node *html.Node, link *url.URL) {
	for i, attr := range node.Attr {
		if attr.Key == "href" || attr.Key == "src" {
			attr.Val = strings.TrimSpace(attr.Val)
			_, err := url.Parse(attr.Val)
			if err == nil {
				if strings.HasPrefix(attr.Val, "http") {
					attr.Val = firstSite() + attr.Val
				} else if strings.HasPrefix(attr.Val, "//") {
					attr.Val = firstSite() + link.Scheme + "://" + attr.Val[2:]
				} else if strings.HasPrefix(attr.Val, "/") {
					attr.Val = firstSite() + link.Scheme + "://" + link.Host + attr.Val
				}
				debagPrintLn("changeURL:",
					"was:", node.Attr[i].Val, "became:", attr.Val)

				node.Attr[i].Val = attr.Val
			}
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		changeURL(child, link)
	}
}

func debagPrintLn(v ...interface{}) {
	log.Println(v...)
}

func firstSite() string {
	return "http://localhost:" + PORT + "/?link="
}

func getAttrValue(node *html.Node, attrName string) string {
	for _, attr := range node.Attr {
		if attr.Key == attrName {
			return attr.Val
		}
	}
	return ""
}
