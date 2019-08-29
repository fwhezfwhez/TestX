package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"time"
)

type TestGorm struct {
	Id        int     `gorm:"column:id"`
	Username  string  `gorm:"column:username;default:"`
	State     int     `gorm:"column:state;default:"`
	CreatedAt string  `gorm:"column:created_at;default:"`
	Money     float64 `gorm:"column:money"`
}

func (r TestGorm) TableName() string {
	return "test_gorm"
}

func (r *TestGorm) AfterUpdate() error {
	b, _ := json.MarshalIndent(r, "  ", "  ")
	fmt.Println(string(b))
	return nil
}

type TestGorm2 struct {
	Id          int    `gorm:"column:id"`
	TestGorm1Id int    `gorm:"column:test_gorm1_id"`
	Username    string `gorm:"column:username"`
	//TG1 TestGorm `gorm:"ForeignKey:TestGorm1Id;AssociationForeignKey:Id" `
}

func (r TestGorm2) TableName2() string {
	return "test_gorm2"
}

func main() {

	db, err := gorm.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
			"localhost",
			"5432",
			"postgres",
			"test",
			"disable",
			"123",
		),
	)

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetConnMaxLifetime(10 * time.Second)
	db.DB().SetMaxIdleConns(30)
	if err != nil {
		panic(err)
	}
	//if e:=db.Model(&TestGorm{}).Create(&tg).Error;e!=nil {
	//	panic(e)
	//}
	//if e:=db.Table("test_gorm").Create(&tg).Error;e!=nil {
	//	panic(e)
	//}
	//db.Model(&TestGorm{}).Where("id = 13").Updates(map[string]interface{}{
	//	"username": "cjc33",
	//})
	//var tg2 TestGorm2
	//db.Table("test_gorm2 as t2").Select("t2.*,t1.username as username").Joins("left join test_gorm as t1 on t1.id=t2.test_gorm1_id").Order("t2.id desc").First(&tg2)
	//fmt.Println(tg2)
	//db.Model(&TestGorm{}).Where("id >1").Exec("update test_gorm set state = state + 1")
	//db.Model(&TestGorm{}).Where("id >1").Delete(&TestGorm{})
	//var t1s []TestGorm
	//if e:=db.Model(&TestGorm{}).Raw("select * from test_gorm").Where("1=?",1).Scan(&t1s).Error;e!=nil {
	//	panic(e)
	//}
	//fmt.Println(t1s)
	//var tx = db.Begin()
	//if e:=tx.Model(&TestGorm{}).Save(&TestGorm{Username:"ftf"}).Error;e!=nil {
	//	panic(e)
	//}
	//tx.Commit()
	//if e:=db.Model(&TestGorm{}).Create(&TestGorm{
	//	Id:60,
	//	Username:"ff",
	//}).Error;e!=nil{
	//	panic(e)
	//}
	var users []struct {
		Id       int    `gorm:"id"`
		Username string `gorm:"username"`
	}
	//if e := db.Raw("select * from user_info").Scan(&users).Error;e!=nil{
	//	panic(e)
	//}

	var ts []TestGorm
	if e:=db.Model(&TestGorm{}).Where("id in (?)", []int{1,2,3,4}).Find(&ts).Error;e!=nil{
		panic(e)
	}
	e:=db.DB().Ping()
	if e!=nil {
		panic(e)
	}
	fmt.Println(users)


}
