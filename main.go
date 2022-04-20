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
	db := linkdb()
	defer closedb(db)
	c := colly.NewCollector()
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println(r.StatusCode, " ", r.Request.URL, err)
	})
	c.OnHTML("div[class='category-wrap_iQLoo horizontal_1eKyQ']", func(e *colly.HTMLElement) {
		//fmt.Println(e)
		title := e.DOM.Find(".c-single-text-ellipsis").Text()
		hot := e.DOM.Find(".hot-index_1Bl1a").Text()
		nums++
		fmt.Printf("%d %8s %s\n", nums, hot, title)
	})
	c.Visit(urlname)
	arow := formdata(name, types, nums)
	fmt.Println(arow)
	fmt.Println("确定将数据载入数据库输入yes")
	fmt.Scanf("%s", &order)
	filename := FnamePro(BMP)("hello")
	fmt.Println(filename)
	if order != "yes" {
		return
	}
	insertdata(formdata(name, types, nums), db)
}
