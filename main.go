package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

var urlname = "https://top.baidu.com/board?tab=realtime"

func main() {
	var order string
	var nums int = 0
	var types string = Hot
	var name string = "百度头条"
	var db myDB
	StoreS := make([]string, 100)
	db.linkdb()
	defer db.closedb()
	c := colly.NewCollector()

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println(r.StatusCode, " ", r.Request.URL, err)
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.Headers)
	})
	c.OnHTML("div[class='category-wrap_iQLoo horizontal_1eKyQ']", func(e *colly.HTMLElement) {
		title := e.DOM.Find(".c-single-text-ellipsis").Text()
		hot := e.DOM.Find(".hot-index_1Bl1a").Text()
		StoreS = append(StoreS, title+"	"+hot+"\n")
		nums++
	})
	c.Visit(urlname)
	arow := formdata(name, types, nums)
	fmt.Println(arow)
	//db.findbytime()
	fmt.Println("确定将数据载入数据库输入yes")
	fmt.Scanf("%s", &order)
	if order != "yes" {
		return
	}
	//文件生成与数据库填充
	storefile("baidu"+curtime, TXT, StoreS)
	printfile("baidu"+curtime, TXT)
	db.insertdata(arow)
}
