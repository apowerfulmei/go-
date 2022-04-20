package main

import (
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
func storefile(name string) {
	//创建文件存储爬取的信息
	file, err := os.Create(path + name)
	check(err)
	file.Close()
}
