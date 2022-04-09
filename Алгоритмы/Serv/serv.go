package main

import (

"fmt"

"github.com/mmcdole/gofeed/rss"

)

func main() {

rssObject, err := rss.ParseRSS("http://blagnews.ru/rss_vk.xml")

if err != nil {

for v := range rssObject.Channel.Items {

item := rssObject.Channel.Items[v]

fmt.Println()

fmt.Printf("Item Number : %d\n", v)

fmt.Printf("Title : %s\n", item.Title)

fmt.Printf("Link : %s\n", item.Link)

fmt.Printf("Description : %s\n", item.Description)Ш

fmt.Printf("Guid : %s\n", item.Guid.Value)

}

}

}