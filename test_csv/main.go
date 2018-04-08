package main

import (
	"fmt"
	"github.com/xormplus/xorm"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"github.com/gin-gonic/gin"
	"bytes"
	"encoding/csv"
	"strings"
	_ "github.com/lib/pq"
	"strconv"
)

/**
 * 导出处理
 */

const ()

var DataSource string
var DB *xorm.Engine

type User struct {
	Name     string
	Province string
	City     string
	Address  string
	Phone    string
	Status   int
	Tdate    string
}

func init() {
	var err error
	//DataSource = "postgres://postgres:123@localhost:5432/travel?sslmode=disable"
	DataSource = "postgres://medium:mediuml4eLxglxL8@111.231.137.127:5432/medium?sslmode=disable"
	DB, err = xorm.NewPostgreSQL(DataSource)
	if err != nil {
		fmt.Println(err)
	}
	//2.显示sql语句
	DB.ShowSQL(true)

	//3.设置连接数
	DB.SetMaxIdleConns(2000)
	DB.SetMaxOpenConns(1000)
}

func main() {

	router := gin.Default()
	router.GET("/data",Data)
	router.Run(":8000")
}
func GetUserByDate(startTime string, endTime string) ([]User, error) {
	users := make([]User, 0)
	err := DB.SQL("select distinct name,province,city,address,phone,status,tdate from travelUser where tdate between ? and ?", startTime, endTime).Find(&users)
	if err != nil {
		return nil, err
	}
	for i := range users {
		users[i].Tdate = users[i].Tdate[:10]
	}
	return users, nil
}

func UTF82GBK(src []byte) (string, error) {
	reader:= transform.NewReader(strings.NewReader(string(src)), simplifiedchinese.GBK.NewEncoder())
	if buf, err := ioutil.ReadAll(reader); err != nil {
		return "", err
	} else {
		return string(buf), nil
	}
}
func Data(c *gin.Context) {
	fileName := "test.csv"
	users, err := GetUserByDate("2018/4/3", "2018/4/3")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(users)
	usersStrArray :=structs2StringArr(users)
	categoryHeader := []string{"姓名","省份","城市","详细地址","填写日期","联系方式","是否联系"}
	b := &bytes.Buffer{}
	wr := csv.NewWriter(b)
	wr.Write(categoryHeader)
	for i := 0; i < len(usersStrArray); i++ {
		wr.Write(usersStrArray[i])
	}
	wr.Flush()
	c.Writer.Header().Set("Content-Type", "text/csv")
	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%s", fileName))
	//tet, _ := UTF82GBK(b.Bytes())
	tet:=b.String()
	c.String(200, tet)
}
func structs2StringArr(users []User) [][]string{
	var userArr =make([][]string,0)
	var user = User{}
	for i:=0;i<len(users);i++{
		user = users[i]
		userArr = append(userArr,[]string{user.Name,user.Province,user.City,user.Address,user.Tdate,user.Phone,strconv.Itoa(user.Status)})
	}
	return userArr
}