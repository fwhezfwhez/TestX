package Dao

import (
	db "Pay/Common/xormTool"
	"time"
)
type Order struct{
	Id int
	OrderId string `xorm:"orderId"`     //如果不做适配，会默认编程order_id导致数据库找不到
	Cid string
	Pversion string
	Ttime time.Time
	Ip string
	OrderStatus int `xorm:"orderStatus"`
	Phone string
	Uname string
	PayStatus int `xorm:"payStatus"`
}
func  Query(dest interface{}) {
	db.Select(dest,"select id,orderId,cid,pversion,ttime,ip,orderStatus,phone,uname,paystatus from payorder")
	db.Select(new([]Order),"select id,orderId,cid,pversion,ttime,ip,orderStatus,phone,uname,paystatus from payorder")

}