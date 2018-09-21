package test

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin/json"
)

type Data struct{
	Id int
	Phone int
	Lon string
	Lat string
	Date int64
}
func findbyphone(w http.ResponseWriter, r *http.Request){
	data:=make([]Data,0)
	r.ParseForm()
	phone := r.PostFormValue("phone")
	if phone == ""{
		fmt.Fprint(w, "请勿非法访问")
		return
	}
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/key")
	if err != nil {
		log.Fatal("db", err)
		return
	}
	result, err :=db.Query(
		"SELECT * FROM record WHERE phone=?",
		phone,
	)
	if err != nil {
		log.Fatal("db", err)
		return
	}

	for result.Next(){
		dataTemp:=Data{}
		err = result.Scan(&dataTemp.Id,&dataTemp.Phone, &dataTemp.Lon, &dataTemp.Lat,&dataTemp.Date)
		data = append(data,dataTemp)
		if err != nil {
			log.Fatal("db", err)
			return
		}
	}
	dataJson,er:= json.Marshal(data)
	if er!=nil{
		fmt.Println(er)
		return
	}
	w.Write(dataJson)
	fmt.Println(string(dataJson))
}
