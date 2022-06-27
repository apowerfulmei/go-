package main

import (
	"fmt"
	"log"
	"os"
)

var path = "D:\\Code\\go\\gethotpot\\store\\"

const (
	BMP = 0
	TXT = 1
)

func FnamePro(sel int) func(string) string {
	//文件名生成器加工厂
	var bac string
	switch sel {
	case BMP:
		bac = "bmp"
	case TXT:
		bac = "txt"
	default:
		bac = "txt"
	}
	return func(pre string) string {
		return pre + "." + bac
	}
}
func storefile(name string, filetype int, mesg []string) {
	//创建文件存储爬取的信息
	Cname := FnamePro(filetype)
	file, err := os.Create(Cname(path + name))
	for _, s := range mesg {
		file.WriteString(s)
	}
	check(err)
	file.Close()
	log.Println("文件生成成功")
}
func printfile(name string, filetype int) {
	//展示文件内容，以时间为查询条件
	var mesg []byte
	Cname := FnamePro(filetype)
	file, err := os.Open(Cname(path + name))
	check(err)
	file.Read(mesg)
	fmt.Println(string(mesg))
}
