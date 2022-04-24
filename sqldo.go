package main

//爬虫信息记录数据库操作
/*table结构 db:getdata
*name 	命名 		varchar(10)
*time 	爬取日期 		date
*dtype  爬取数据类型 	varchar(20)
*num 	爬取数量 		int
 */
import (
	"database/sql"
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var curtime = time.Now().Format("2006-01-02")

//爬取数据类型
/*
* pic	图片
* hot	热点
 */
const Pic = "pictures"
const Hot = "hot_point"

type RowData struct {
	name  string
	time  string
	dtype string
	num   int
}

var dbname = "testdb"
var tbname = "getdata"

func linkdb() *sql.DB {
	//连接getdata数据库
	db, err := sql.Open("mysql", "root:528320@tcp(127.0.0.1:3306)/"+dbname)
	check(err)
	row, err := db.Query("select * from " + tbname + " where dtype='mky'")
	check(err)
	printrow(row)
	return db
}

func insertdata(data RowData, db *sql.DB) {
	//插入新数据
	_, err := db.Exec("insert into "+tbname+
		"(name,time,dtype,num) "+
		"values(?,?,?,?)", data.name, data.time, data.dtype, data.num)
	check(err)
	fmt.Println("数据插入成功")
}

func findbytime(db *sql.DB) {
	//按照起止时间查询数据库
	var startdate string
	var enddate string
	fmt.Println("进入查询界面")
	fmt.Printf("请输入查询起始日期：")
	fmt.Scanf("%s\n", &startdate)
	fmt.Printf("请输入查询截至日期：")
	fmt.Scanf("%s\n", &enddate)
	fmt.Println(startdate, enddate)
	row, err := db.Query("select * from " + tbname + " where time>='" + startdate + "' and time<='" + enddate + "'")
	check(err)
	printrow(row)
	fmt.Println("查询完毕！")
}

func formdata(n string, d string, num int) RowData {
	var row RowData
	row.name = n
	row.time = curtime
	row.dtype = d
	row.num = num
	return row
}

func closedb(db *sql.DB) {
	//关闭数据库
	db.Close()
	fmt.Println("Database is closed!")
}

func check(err error) {
	//查看是否存在错误
	if err != nil {
		fmt.Println("something is wrong")
		panic(err)
	}
}

func printrow(row *sql.Rows) {
	//打印query得到的数据行
	var sayhi RowData
	for row.Next() {
		row.Scan(&sayhi.name, &sayhi.time, &sayhi.dtype, &sayhi.num)
		fmt.Printf("%s %s %s %d\n", sayhi.name, sayhi.time, sayhi.dtype, sayhi.num)
	}
}
